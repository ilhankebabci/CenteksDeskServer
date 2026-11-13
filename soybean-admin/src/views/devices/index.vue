<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { deleteDevice, fetchCompanyList, fetchDevicesList } from '@/service/api/devices';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import AuditBaseLogsSearch from './components/search.vue';
import DeviceEditModal from './components/device-edit-modal.vue';
import CompanyModal from './components/company-modal.vue';

const appStore = useAppStore();

const editModalVisible = ref(false);
const companyModalVisible = ref(false);
const currentEditingDevice = ref<any>(null);

const selectedDepartment = ref('Tümü');
const selectedCompany = ref('Tümü');

const departmentList = ['Tümü', 'Bilgi İşlem', 'Fabrika', 'Ofis / Muhasebe', 'Sunucular'];
const companyList = ref<string[]>(['Centeks Group']);

async function loadCompanies() {
  const res = await fetchCompanyList();
  if (res.data && res.data.length > 0) {
    companyList.value = res.data;
  }
}

function handleOpenEdit(row: any) {
  currentEditingDevice.value = row;
  editModalVisible.value = true;
}

function handleConnect(row: any) {
  const cleanId = String(row.rustdesk_id || '').replace(/\s+/g, '');
  if (!cleanId) return;
  window.open(`rustdesk://${cleanId}?password=Centeks2027!`, '_self');
  window.$message?.info(`rustdesk://${cleanId} tetiklendi`);
}

async function handleDelete(id: number) {
  await deleteDevice(id);
  window.$message?.success('Cihaz başarıyla silindi');
  getData();
}

function handleDepartmentChange(dept: string) {
  selectedDepartment.value = dept;
  searchParams.department = dept === 'Tümü' ? null : dept;
  getDataByPage();
}

function handleCompanyChange(comp: string) {
  selectedCompany.value = comp;
  searchParams.company = comp === 'Tümü' ? null : comp;
  getDataByPage();
}

function onCompanyAdded() {
  loadCompanies();
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
            <span class="font-bold text-slate-800 text-sm">
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
        const typeMap: Record<string, 'info' | 'warning' | 'success' | 'error' | 'default'> = {
          'Bilgi İşlem': 'info',
          'Fabrika': 'warning',
          'Ofis / Muhasebe': 'success',
          'Sunucular': 'error'
        };
        return <NTag type={typeMap[dept] || 'default'} size="small">{dept}</NTag>;
      }
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.device.rustdesk_id'),
      align: 'center',
      width: 140,
      render: (row: any) => {
        return <span class="font-mono font-bold text-amber-600 text-sm">{row.rustdesk_id}</span>;
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
  loadCompanies();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <!-- Firma ve Departman Hızlı Filtre Barı -->
    <NCard :bordered="false" size="small" class="card-wrapper">
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
              <NRadioButton v-for="comp in companyList" :key="comp" :value="comp">
                {{ comp }}
              </NRadioButton>
            </NRadioGroup>
          </div>
          <NButton
            size="small"
            type="info"
            dashed
            @click="companyModalVisible = true"
          >
            + Yeni Firma / Şirket Ekle
          </NButton>
        </div>

        <!-- 2. Departman Filtresi -->
        <div class="flex items-center gap-2 flex-wrap pt-2 border-t border-slate-100 dark:border-slate-800">
          <span class="text-xs font-semibold text-slate-500 flex items-center gap-1">
            📁 Departman:
          </span>
          <NRadioGroup
            :value="selectedDepartment"
            size="small"
            @update:value="handleDepartmentChange"
          >
            <NRadioButton v-for="dept in departmentList" :key="dept" :value="dept">
              {{ dept }}
            </NRadioButton>
          </NRadioGroup>
        </div>
      </div>
    </NCard>

    <AuditBaseLogsSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <NCard :title="$t('route.devices')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
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
          <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="getData" />
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
      :company-list="companyList"
      @submitted="getData"
    />

    <!-- Yeni Firma Ekleme Modalı -->
    <CompanyModal
      v-model:visible="companyModalVisible"
      @submitted="onCompanyAdded"
    />
  </div>
</template>
