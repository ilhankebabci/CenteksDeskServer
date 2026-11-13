<script setup lang="ts">
import { computed } from 'vue';

defineOptions({
  name: 'ChangeLogs'
});

interface LogItem {
  id: number;
  content: string;
  version: string;
  time: string;
}

const logs = computed<LogItem[]>(() => [
  { id: 1, content: 'CENTEKS Destek v1.5.3 İstemci Entegrasyonu ve Türkçe Panel', version: 'v1.5.3', time: '2026-11-13' },
  { id: 2, content: 'CENTEKS Yönetim Paneli Güvenlik ve Yetkilendirme Güncellemesi', version: 'v1.5.0', time: '2026-11-12' }
]);
</script>

<template>
  <NCard :title="$t('page.home.changeLogs')" :bordered="false" size="small" segmented class="card-wrapper">
    <NList>
      <NListItem v-for="item in logs" :key="item.id">
        <template #prefix>
          <SoybeanAvatar class="size-48px!" />
        </template>
        <NThing :title="item.content" :title-extra="item.version" :description="item.time" />
      </NListItem>
    </NList>
  </NCard>
</template>

<style scoped></style>
