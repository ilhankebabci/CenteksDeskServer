<script setup lang="ts">
defineOptions({
  name: 'GlobalFooter'
});
</script>

<template>
  <DarkModeContainer class="h-full flex items-center justify-between px-6 text-xs text-slate-500 border-t border-slate-100 dark:border-slate-800">
    <div class="flex items-center gap-2">
      <span class="font-bold text-amber-600">CENTEKS Destek v1.5.3</span>
      <span class="text-slate-300 dark:text-slate-700">•</span>
      <span>Centeks Group Bilgi İşlem Departmanı Tarafından Geliştirilmiştir</span>
    </div>
    <div class="flex items-center gap-4">
      <span class="text-emerald-600 font-medium flex items-center gap-1.5">
        <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
        Sistem Aktif & Güvenli
      </span>
      <span class="text-slate-300 dark:text-slate-700">•</span>
      <span>© 2026 CENTEKS Group • Tüm Hakları Saklıdır</span>
    </div>
  </DarkModeContainer>
</template>

<style scoped></style>
