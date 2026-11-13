<script setup lang="ts">
defineOptions({
  name: 'CreativityBanner'
});
</script>

<template>
  <NCard title="CENTEKS Destek Bilgi Merkezi" :bordered="false" size="small" class="h-full card-wrapper">
    <template #header-extra>
      <NTag type="success" size="small">🟢 Ağ Aktif</NTag>
    </template>
    <div class="h-full flex flex-col justify-center space-y-3 p-2">
      <div class="flex items-center justify-between text-xs py-1 border-b border-slate-100">
        <span class="text-slate-500">Rendezvous Sunucusu:</span>
        <span class="font-mono font-bold text-slate-800">destek.centeks.com</span>
      </div>
      <div class="flex items-center justify-between text-xs py-1 border-b border-slate-100">
        <span class="text-slate-500">Servis Portları:</span>
        <span class="font-mono font-bold text-slate-800">21115 - 21119 / 8080</span>
      </div>
      <div class="flex items-center justify-between text-xs py-1 border-b border-slate-100">
        <span class="text-slate-500">Şifreleme:</span>
        <span class="font-mono text-emerald-600 font-bold">256-bit TLS / WebRTC</span>
      </div>
      <div class="flex items-center justify-between text-xs py-1">
        <span class="text-slate-500">Yönetim:</span>
        <span class="font-semibold text-amber-600">CENTEKS Bilgi İşlem Operasyon</span>
      </div>
    </div>
  </NCard>
</template>

<style scoped></style>
