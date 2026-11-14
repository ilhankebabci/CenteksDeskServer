<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import {
  NButton,
  NCard,
  NDivider,
  NEmpty,
  NGrid,
  NGi,
  NInput,
  NModal,
  NPopconfirm,
  NSpace,
  NTag
} from 'naive-ui';
import {
  addCompany,
  addDepartment,
  deleteCompany,
  deleteDepartment
} from '@/service/api/devices';

interface Props {
  visible: boolean;
  treeData: Api.Devices.CompanyTreeItem[];
}

const props = defineProps<Props>();
const emit = defineEmits(['update:visible', 'submitted']);

const selectedCompanyName = ref<string>('Centeks Group');
const newCompanyName = ref('');
const newCompanyNote = ref('');
const newDeptName = ref('');

watch(
  () => props.treeData,
  val => {
    if (val && val.length > 0 && !selectedCompanyName.value) {
      selectedCompanyName.value = val[0].name;
    }
  },
  { immediate: true }
);

const activeCompany = computed(() => {
  return props.treeData.find(c => c.name === selectedCompanyName.value) || props.treeData[0];
});

async function handleAddCompany() {
  const name = newCompanyName.value.trim();
  if (!name) {
    window.$message?.warning('Lütfen bir firma adı girin');
    return;
  }
  await addCompany({ name, note: newCompanyNote.value.trim() });
  window.$message?.success(`"${name}" firması eklendi`);
  newCompanyName.value = '';
  newCompanyNote.value = '';
  selectedCompanyName.value = name;
  emit('submitted');
}

async function handleDeleteCompany(name: string) {
  if (name === 'Centeks Group' || name === 'Centeks') {
    window.$message?.error('Ana firma silinemez');
    return;
  }
  await deleteCompany(name);
  window.$message?.success(`"${name}" firması silindi`);
  selectedCompanyName.value = 'Centeks Group';
  emit('submitted');
}

async function handleAddDept() {
  const dept = newDeptName.value.trim();
  if (!dept) {
    window.$message?.warning('Lütfen departman adı girin');
    return;
  }
  if (!activeCompany.value) return;

  await addDepartment({
    company_name: activeCompany.value.name,
    name: dept
  });
  window.$message?.success(`"${dept}" departmanı "${activeCompany.value.name}" altına eklendi`);
  newDeptName.value = '';
  emit('submitted');
}

async function handleDeleteDept(companyName: string, deptName: string) {
  await deleteDepartment({
    company_name: companyName,
    name: deptName
  });
  window.$message?.success(`"${deptName}" departmanı kaldırıldı`);
  emit('submitted');
}

function handleClose() {
  emit('update:visible', false);
}
</script>

<template>
  <NModal
    :show="visible"
    preset="card"
    title="🏢 Şirket & Alt Departman Yönetimi"
    class="w-800px max-w-95vw"
    :mask-closable="true"
    @update:show="handleClose"
  >
    <div class="text-xs text-slate-500 mb-4">
      Holding ve alt firmalarınızı buradan oluşturabilir, her bir firmanın altına özel departmanlar (Muhasebe, Kalite, Üretim vb.) tanımlayabilirsiniz.
    </div>

    <NGrid :x-gap="16" :y-gap="16" cols="1 m:2">
      <!-- SOL TARAF: FİRMALAR / ŞİRKETLER LİSTESİ -->
      <NGi>
        <NCard title="1. Firmalar / Şirketler" size="small" class="h-full">
          <div class="space-y-2 max-h-280px overflow-y-auto pr-1">
            <div
              v-for="comp in treeData"
              :key="comp.name"
              class="flex items-center justify-between p-2.5 rounded-xl border transition-all cursor-pointer"
              :class="
                selectedCompanyName === comp.name
                  ? 'bg-amber-500/10 border-amber-500/50 shadow-sm'
                  : 'border-slate-100 hover:border-slate-300 dark:border-slate-800'
              "
              @click="selectedCompanyName = comp.name"
            >
              <div class="flex items-center gap-2">
                <span class="text-base">🏢</span>
                <div>
                  <div class="text-xs font-bold text-slate-800 dark:text-slate-100">
                    {{ comp.name }}
                  </div>
                  <div class="text-[11px] text-slate-400">
                    {{ comp.departments.length }} Departman • {{ comp.device_count }} Cihaz
                  </div>
                </div>
              </div>

              <div class="flex items-center gap-1" @click.stop>
                <NPopconfirm
                  v-if="comp.name !== 'Centeks Group' && comp.name !== 'Centeks'"
                  on-positive-click="() => handleDeleteCompany(comp.name)"
                >
                  <template #trigger>
                    <NButton size="tiny" type="error" quaternary>Sil</NButton>
                  </template>
                  "{{ comp.name }}" firmasını ve bağlı departmanlarını silmek istediğinize emin misiniz?
                </NPopconfirm>
                <NTag v-else size="tiny" type="warning" round>Ana Firma</NTag>
              </div>
            </div>
          </div>

          <NDivider style="margin: 12px 0;" />

          <!-- Yeni Firma Ekleme Kutusu -->
          <div class="space-y-2">
            <div class="text-xs font-semibold text-slate-700 dark:text-slate-200">
              + Yeni Firma Ekle:
            </div>
            <div class="flex gap-2">
              <NInput
                v-model:value="newCompanyName"
                size="small"
                placeholder="Örn: Akeray, Centeks Kimya..."
                @keyup.enter="handleAddCompany"
              />
              <NButton size="small" type="primary" @click="handleAddCompany">
                Ekle
              </NButton>
            </div>
          </div>
        </NCard>
      </NGi>

      <!-- SAĞ TARAF: SEÇİLİ FİRMANIN DEPARTMANLARI -->
      <NGi>
        <NCard
          :title="`2. ${activeCompany ? activeCompany.name : ''} Departmanları`"
          size="small"
          class="h-full"
        >
          <div v-if="activeCompany" class="space-y-3">
            <div class="text-xs text-slate-400">
              Bu firmanın altında tanımlı departmanlar. Cihaz eklerken veya filtrelerken bu birimler listelenecektir.
            </div>

            <!-- Departman Etiketleri (Silinebilir) -->
            <div class="min-h-120px p-3 rounded-xl bg-slate-50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 flex flex-wrap gap-2 items-start">
              <NTag
                v-for="dept in activeCompany.departments"
                :key="dept"
                type="info"
                closable
                round
                size="medium"
                @close="() => handleDeleteDept(activeCompany.name, dept)"
              >
                📁 {{ dept }}
              </NTag>
              <NEmpty
                v-if="activeCompany.departments.length === 0"
                description="Henüz departman eklenmemiş"
                class="w-full py-4"
              />
            </div>

            <NDivider style="margin: 12px 0;" />

            <!-- Yeni Departman Ekleme Kutusu -->
            <div class="space-y-2">
              <div class="text-xs font-semibold text-slate-700 dark:text-slate-200">
                + "{{ activeCompany.name }}" Altına Departman Ekle:
              </div>
              <div class="flex gap-2">
                <NInput
                  v-model:value="newDeptName"
                  size="small"
                  placeholder="Örn: Muhasebe, Kalite, Üretim, Sevkiyat..."
                  @keyup.enter="handleAddDept"
                />
                <NButton size="small" type="primary" @click="handleAddDept">
                  Departman Ekle
                </NButton>
              </div>
            </div>
          </div>
        </NCard>
      </NGi>
    </NGrid>

    <template #footer>
      <NSpace justify="end">
        <NButton type="primary" @click="handleClose">Tamamla ve Kapat</NButton>
      </NSpace>
    </template>
  </NModal>
</template>
