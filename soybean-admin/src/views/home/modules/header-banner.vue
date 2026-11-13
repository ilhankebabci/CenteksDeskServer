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
      <NGi span="24 s:24 m:18">
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
    </NGrid>
  </NCard>
</template>

<style scoped></style>
