<script setup lang="tsx">
import { computed, onMounted, ref } from 'vue';
import { NButton, NCard, NDataTable, NPopconfirm, NRadioButton, NRadioGroup, NSpace, NTag } from 'naive-ui';
import { deleteDevice, fetchCompanyTree, fetchDevicesList } from '@/service/api/devices';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import AuditBaseLogsSearch from './components/search.vue';
import DeviceEditModal from './components/device-edit-modal.vue';
import OrgManageModal from './components/org-manage-modal.vue';

const appStore = useAppStore();

const editModalVisible = ref(false);
const orgModalVisible = ref(false);
const currentEditingDevice = ref<any>(null);

const selectedCompany = ref('Tümü');
const selectedDepartment = ref('Tümü');
const companyTree = ref<Api.Devices.CompanyTreeItem[]>([]);

async function loadOrgTree() {
  const res = await fetchCompanyTree();
  if (res.data) {
    companyTree.value = res.data;
  }
}

const activeDepartments = computed(() => {
  if (selectedCompany.value === 'Tümü') {
    const allDepts = new Set<string>();
    companyTree.value.forEach(comp => {
      comp.departments.forEach(d => allDepts.add(d));
    });
    const defaultDepts = ['Bilgi İşlem', 'Muhasebe', 'Kalite', 'Fabrika', 'Sunucular'];
    defaultDepts.forEach(d => allDepts.add(d));
    return ['Tümü', ...Array.from(allDepts)];
  }

  const currentComp = companyTree.value.find(c => c.name === selectedCompany.value);
  if (!currentComp || currentComp.departments.length === 0) {
    return ['Tümü', 'Bilgi İşlem', 'Muhasebe', 'Kalite', 'Fabrika'];
  }
  return ['Tümü', ...currentComp.departments];
});

function handleOpenEdit(row: any) {
  currentEditingDevice.value = row;
  editModalVisible.value = true;
}

function copyToClipboard(text: string) {
  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard.writeText(text);
  } else {
    const textArea = document.createElement('textarea');
    textArea.value = text;
    textArea.style.position = 'fixed';
    textArea.style.left = '-999999px';
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();
    try {
      document.execCommand('copy');
    } catch (e) {}
    textArea.remove();
  }
}

function handleCopyId(id: string) {
  const cleanId = String(id || '').replace(/\s+/g, '');
  if (!cleanId) return;
  copyToClipboard(cleanId);
  window.$message?.success(`📋 ID Kopyalandı: ${cleanId}`);
}

function handleConnect(row: any) {
  const cleanId = String(row.rustdesk_id || '').replace(/\s+/g, '');
  if (!cleanId) return;
  
  copyToClipboard(cleanId);
  
  try {
    const link = document.createElement('a');
    link.href = `rustdesk://${cleanId}?password=Centeks2027!`;
    link.click();
  } catch (e) {}
  
  window.$message?.success(`⚡ ID Kopyalandı: ${cleanId} (RustDesk açılıyor...)`);
}

async function handleDelete(id: number) {
  await deleteDevice(id);
  window.$message?.success('Cihaz başarıyla silindi');
  loadOrgTree();
  getData();
}

function handleCompanyChange(comp: string) {
  selectedCompany.value = comp;
  selectedDepartment.value = 'Tümü';
  searchParams.company = comp === 'Tümü' ? null : comp;
  searchParams.department = null;
  getDataByPage();
}

function handleDepartmentChange(dept: string) {
  selectedDepartment.value = dept;
  searchParams.department = dept === 'Tümü' ? null : dept;
  getDataByPage();
}

function onOrgSubmitted() {
  loadOrgTree();
  getData();
}

const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: fetchDevicesList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    hostname: null,
    username: null,
    rustdesk_id: null,
    department: null,
    company: null
  },
  columns: () => [
    {
      key: 'is_online',
      title: 'Durum',
      align: 'center',
      width: 120,
      render: (row: any) => {
        return (
          <NTag type={row.is_online ? 'success' : 'default'} round size="small">
            {row.is_online ? '🟢 Çevrimiçi' : '⚫ Çevrimdışı'}
          </NTag>
        );
      }
    },
    {
      key: 'alias',
      title: 'Cihaz Adı / İsim',
      align: 'left',
      minWidth: 180,
      render: (row: any) => {
        return (
          <div class="flex flex-col py-1">
            <span class="font-bold text-slate-800 dark:text-slate-100 text-sm">
              {row.alias || row.hostname || 'İsimsiz Cihaz'}
            </span>
            {row.alias && (
              <span class="text-xs text-slate-400 font-mono">
                {row.hostname || row.username}
              </span>
            )}
          </div>
        );
      }
    },
    {
      key: 'company',
      title: 'Firma / Şirket',
      align: 'center',
      width: 160,
      render: (row: any) => {
        const comp = row.company || 'Centeks Group';
        return (
          <NTag type="info" size="small" round>
            🏢 {comp}
          </NTag>
        );
      }
    },
    {
      key: 'department',
      title: 'Departman / Grup',
      align: 'center',
      width: 150,
      render: (row: any) => {
        const dept = row.department || 'Genel';
        return (
          <NTag type="warning" size="small" round>
            📁 {dept}
          </NTag>
        );
      }
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.device.rustdesk_id'),
      align: 'center',
      width: 140,
      render: (row: any) => {
        return (
          <span
            class="font-mono font-bold text-amber-600 text-sm cursor-pointer hover:bg-amber-50 dark:hover:bg-amber-950/40 px-2 py-0.5 rounded transition-all select-all inline-flex items-center gap-1"
            title="Kopyalamak için tıklayın"
            onClick={() => handleCopyId(row.rustdesk_id)}
          >
            {row.rustdesk_id}
            <span class="text-xs opacity-60">📋</span>
          </span>
        );
      }
    },
    {
      key: 'username',
      title: $t('dataMap.device.username'),
      align: 'center',
      width: 130
    },
    {
      key: 'os',
      title: $t('dataMap.device.os'),
      align: 'center',
      minWidth: 150
    },
    {
      key: 'version',
      title: $t('dataMap.device.version'),
      align: 'center',
      width: 110
    },
    {
      key: 'updated_at',
      title: 'Son Sinyal',
      align: 'center',
      width: 160,
      render: (row: any) => row.updated_at || row.created_at
    },
    {
      key: 'actions',
      title: 'İşlemler',
      align: 'center',
      width: 220,
      fixed: 'right',
      render: (row: any) => {
        return (
          <NSpace justify="center" size="small">
            <NButton
              type="primary"
              size="small"
              onClick={() => handleConnect(row)}
            >
              ⚡ Bağlan
            </NButton>
            <NButton
              size="small"
              onClick={() => handleOpenEdit(row)}
            >
              ✏️ Düzenle
            </NButton>
            <NPopconfirm onPositiveClick={() => handleDelete(row.id)}>
              {{
                trigger: () => <NButton size="small" type="error" quaternary>🗑️ Sil</NButton>,
                default: () => 'Bu cihazı listeden silmek istediğinize emin misiniz?'
              }}
            </NPopconfirm>
          </NSpace>
        );
      }
    }
  ]
});

onMounted(() => {
  loadOrgTree();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <!-- Firma ve Departman Hızlı Filtre Barı -->
    <NCard :bordered="false" size="small" class="card-wrapper shadow-sm">
      <div class="space-y-3">
        <!-- 1. Firma / Şirket Filtresi -->
        <div class="flex items-center justify-between flex-wrap gap-2">
          <div class="flex items-center gap-2 flex-wrap">
            <span class="text-xs font-bold text-slate-700 dark:text-slate-300 flex items-center gap-1">
              🏢 Şirket Filtresi:
            </span>
            <NRadioGroup
              :value="selectedCompany"
              size="small"
              @update:value="handleCompanyChange"
            >
              <NRadioButton value="Tümü">Tüm Şirketler</NRadioButton>
              <NRadioButton v-for="comp in companyTree" :key="comp.name" :value="comp.name">
                🏢 {{ comp.name }}
                <span v-if="comp.device_count > 0" class="ml-1 text-[11px] opacity-75">
                  ({{ comp.device_count }})
                </span>
              </NRadioButton>
            </NRadioGroup>
          </div>
          <NButton
            size="small"
            type="primary"
            dashed
            @click="orgModalVisible = true"
          >
            🏢 Şirket & Birim Yönetimi
          </NButton>
        </div>

        <!-- 2. Departman Filtresi (Seçili Firmaya Göre Değişir) -->
        <div class="flex items-center gap-2 flex-wrap pt-2 border-t border-slate-100 dark:border-slate-800">
          <span class="text-xs font-semibold text-slate-500 flex items-center gap-1">
            📁 {{ selectedCompany === 'Tümü' ? 'Tüm Departmanlar' : `${selectedCompany} Departmanları` }}:
          </span>
          <NRadioGroup
            :value="selectedDepartment"
            size="small"
            @update:value="handleDepartmentChange"
          >
            <NRadioButton v-for="dept in activeDepartments" :key="dept" :value="dept">
              📁 {{ dept }}
            </NRadioButton>
          </NRadioGroup>
        </div>
      </div>
    </NCard>

    <AuditBaseLogsSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <NCard :title="$t('route.devices')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper shadow-sm">
      <template #header-extra>
        <div class="flex items-center gap-3">
          <a
            href="/downloads/CenteksDestek-x86_64.exe"
            download
            class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-amber-500 hover:bg-amber-600 text-white font-bold text-xs shadow transition-all active:scale-95 no-underline"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path></svg>
            <span>+ Yeni İstemci İndir</span>
          </a>
          <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="() => { loadOrgTree(); getData(); }" />
        </div>
      </template>
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1350"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
    </NCard>

    <!-- Cihaz Düzenleme Modalı -->
    <DeviceEditModal
      v-model:visible="editModalVisible"
      :device-data="currentEditingDevice"
      :tree-data="companyTree"
      @submitted="() => { loadOrgTree(); getData(); }"
    />

    <!-- Şirket & Alt Departman Yönetim Modalı -->
    <OrgManageModal
      v-model:visible="orgModalVisible"
      :tree-data="companyTree"
      @submitted="onOrgSubmitted"
    />
  </div>
</template>
