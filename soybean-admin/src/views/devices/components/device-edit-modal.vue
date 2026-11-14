<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { updateDevice } from '@/service/api/devices';

interface Props {
  visible: boolean;
  deviceData?: any;
  treeData?: Api.Devices.CompanyTreeItem[];
}

const props = withDefaults(defineProps<Props>(), {
  treeData: () => []
});

const emit = defineEmits(['update:visible', 'submitted']);

const formModel = reactive({
  id: 0,
  alias: '',
  company: 'Centeks Group',
  department: 'Bilgi İşlem'
});

const companyOptions = computed(() => {
  const names = (props.treeData || []).map(c => c.name);
  const list = Array.from(new Set(['Centeks Group', ...names]));
  return list.map(c => ({ label: `🏢 ${c}`, value: c }));
});

const departmentOptions = computed(() => {
  const currentComp = (props.treeData || []).find(c => c.name === formModel.company);
  const compDepts = currentComp ? currentComp.departments : [];
  const defaults = ['Bilgi İşlem', 'Muhasebe', 'Kalite', 'Fabrika', 'Üretim', 'Sevkiyat', 'Sunucular', 'Yönetim', 'Genel'];
  const merged = Array.from(new Set([...compDepts, ...defaults]));
  return merged.map(d => ({ label: `📁 ${d}`, value: d }));
});

watch(
  () => props.deviceData,
  newVal => {
    if (newVal) {
      formModel.id = newVal.id;
      formModel.alias = newVal.alias || '';
      formModel.company = newVal.company || 'Centeks Group';
      formModel.department = newVal.department || 'Bilgi İşlem';
    }
  },
  { immediate: true }
);

async function handleSubmit() {
  if (!formModel.id) return;
  await updateDevice({
    id: formModel.id,
    alias: formModel.alias.trim(),
    company: formModel.company.trim(),
    department: formModel.department.trim()
  });
  window.$message?.success('Cihaz bilgileri başarıyla güncellendi');
  emit('update:visible', false);
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
    title="Cihazı Düzenle & İsimlendir"
    class="w-520px max-w-90vw"
    :mask-closable="true"
    @update:show="handleClose"
  >
    <NForm :model="formModel" label-placement="top">
      <NFormItem label="Cihaz Adı / Takma İsim">
        <NInput
          v-model:value="formModel.alias"
          placeholder="Örn: İlhan Kebabcı Laptop, Muhasebe Fatura PC"
        />
      </NFormItem>

      <NFormItem label="Bağlı Olduğu Şirket / Alt Firma">
        <NSelect
          v-model:value="formModel.company"
          :options="companyOptions"
          filterable
          tag
          placeholder="Şirket seçin veya yeni firma adı yazın..."
        />
      </NFormItem>

      <NFormItem label="Departman / Birim">
        <NSelect
          v-model:value="formModel.department"
          :options="departmentOptions"
          filterable
          tag
          placeholder="Departman seçin veya yeni yazın..."
        />
      </NFormItem>
    </NForm>

    <template #footer>
      <NSpace justify="end">
        <NButton @click="handleClose">İptal</NButton>
        <NButton type="primary" @click="handleSubmit">Kaydet</NButton>
      </NSpace>
    </template>
  </NModal>
</template>
