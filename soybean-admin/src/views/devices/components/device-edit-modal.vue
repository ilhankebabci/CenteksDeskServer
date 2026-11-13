<script setup lang="ts">
import { reactive, watch } from 'vue';
import { updateDevice } from '@/service/api/devices';

interface Props {
  visible: boolean;
  deviceData?: any;
}

const props = defineProps<Props>();
const emit = defineEmits(['update:visible', 'submitted']);

const formModel = reactive({
  id: 0,
  alias: '',
  department: 'Bilgi İşlem'
});

const departmentOptions = [
  { label: 'Bilgi İşlem', value: 'Bilgi İşlem' },
  { label: 'Fabrika', value: 'Fabrika' },
  { label: 'Ofis / Muhasebe', value: 'Ofis / Muhasebe' },
  { label: 'Sunucular', value: 'Sunucular' },
  { label: 'Yönetim', value: 'Yönetim' },
  { label: 'Pazarlama', value: 'Pazarlama' },
  { label: 'Depo / Lojistik', value: 'Depo / Lojistik' },
  { label: 'Genel', value: 'Genel' }
];

watch(
  () => props.deviceData,
  newVal => {
    if (newVal) {
      formModel.id = newVal.id;
      formModel.alias = newVal.alias || '';
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
    department: formModel.department.trim()
  });
  window.$message?.success('Cihaz başarıyla güncellendi');
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
    class="w-500px max-w-90vw"
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

      <NFormItem label="Departman / Grup">
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

