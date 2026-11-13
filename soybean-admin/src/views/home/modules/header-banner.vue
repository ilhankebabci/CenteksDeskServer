<script setup lang="ts">
import { computed } from 'vue';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'HeaderBanner'
});

const appStore = useAppStore();
const authStore = useAuthStore();

const gap = computed(() => (appStore.isMobile ? 0 : 16));
</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <NGrid :x-gap="gap" :y-gap="16" responsive="screen" item-responsive>
      <NGi span="24 s:24 m:16">
        <div class="flex-y-center">
          <div class="size-64px shrink-0 overflow-hidden rd-12px bg-slate-50 flex items-center justify-center p-1 border border-slate-100 shadow-sm">
            <img src="@/assets/imgs/logo.png" class="max-h-full max-w-full object-contain" />
          </div>
          <div class="pl-16px">
            <h3 class="text-18px font-semibold text-slate-800">
              {{ $t('page.home.greeting', { userName: authStore.userInfo.userName || 'Yönetici' }) }}
            </h3>
            <p class="text-12px text-slate-400 mt-1">
              CENTEKS Group Uzak Masaüstü ve Cihaz Operasyon Yönetim Paneli
            </p>
          </div>
        </div>
      </NGi>
      <NGi span="24 s:24 m:8" class="flex items-center justify-end">
        <a
          href="/downloads/CenteksDestek-x86_64.exe"
          download
          class="flex items-center gap-2 px-4 py-2.5 rounded-xl bg-gradient-to-r from-amber-500 to-amber-600 hover:from-amber-600 hover:to-amber-700 text-white font-bold text-sm shadow-md hover:shadow-lg transition-all active:scale-95 no-underline"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path></svg>
          <span>İstemci İndir (.exe)</span>
        </a>
      </NGi>
    </NGrid>
  </NCard>
</template>

<style scoped></style>
