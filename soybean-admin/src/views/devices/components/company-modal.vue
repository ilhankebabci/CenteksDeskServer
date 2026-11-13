<script setup lang="ts">
import { ref } from 'vue';
import { addCompany } from '@/service/api/devices';

defineProps<{ visible: boolean }>();
const emit = defineEmits(['update:visible', 'submitted']);

const companyName = ref('');
const companyNote = ref('');

async function handleAdd() {
  const name = companyName.value.trim();
  if (!name) {
    window.$message?.warning('Lütfen bir firma adı girin');
    return;
  }
  await addCompany({ name, note: companyNote.value.trim() });
  window.$message?.success(`"${name}" firması başarıyla eklendi`);
  companyName.value = '';
  companyNote.value = '';
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
    title="🏢 Yeni Firma / Şirket Ekle"
    class="w-480px max-w-90vw"
    :mask-closable="true"
    @update:show="handleClose"
  >
    <div class="space-y-4">
      <div>
        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 mb-1.5">Firma / Şirket Adı</label>
        <NInput
          v-model:value="companyName"
          placeholder="Örn: Centeks Tekstil A.Ş., Centeks Kimya, Alt Firma..."
          autofocus
        />
      </div>
      <div>
        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 mb-1.5">Açıklama / Not (Opsiyonel)</label>
        <NInput
          v-model:value="companyNote"
          type="textarea"
          placeholder="Firma hakkında kısa not veya konum..."
          :rows="2"
        />
      </div>
    </div>

    <template #footer>
      <NSpace justify="end">
        <NButton @click="handleClose">İptal</NButton>
        <NButton type="primary" @click="handleAdd">Firmayı Kaydet</NButton>
      </NSpace>
    </template>
  </NModal>
</template>

