<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { createReusableTemplate } from '@vueuse/core';
import { $t } from '@/locales';
import { fetchStat } from '@/service/api/home';

defineOptions({
  name: 'CardData'
});

const router = useRouter();

interface CardData {
  key: string;
  title: string;
  value: number;
  unit: string;
  color: {
    start: string;
    end: string;
  };
  icon: string;
  route: string;
  subtitle: string;
}

const stat = ref<Api.Home.Stat>({
  userCount: 0,
  deviceCount: 0,
  onlineCount: 0,
  visitsCount: 0
});

const cardData = computed<CardData[]>(() => [
  {
    key: 'userCount',
    title: $t('page.home.userCount'),
    value: stat.value.userCount,
    unit: '',
    color: {
      start: '#ec4786',
      end: '#b955a4'
    },
    icon: 'gravity-ui:person',
    route: '/user/list',
    subtitle: 'Kullanıcıları Yönet'
  },
  {
    key: 'deviceCount',
    title: $t('page.home.deviceCount'),
    value: stat.value.deviceCount,
    unit: '',
    color: {
      start: '#865ec0',
      end: '#5144b4'
    },
    icon: 'fluent:desktop-32-regular',
    route: '/devices',
    subtitle: 'Tüm Cihazlar'
  },
  {
    key: 'onlineCount',
    title: $t('page.home.onlineCount'),
    value: stat.value.onlineCount,
    unit: '',
    color: {
      start: '#0284c7',
      end: '#0ea5e9'
    },
    icon: 'fluent:desktop-checkmark-20-regular',
    route: '/devices',
    subtitle: 'Çevrimiçi Cihazlar'
  },
  {
    key: 'visitCount',
    title: $t('page.home.visitsCount'),
    value: stat.value.visitsCount,
    unit: '',
    color: {
      start: '#f59e0b',
      end: '#d97706'
    },
    icon: 'ant-design:bar-chart-outlined',
    route: '/audit/baselogs',
    subtitle: 'Bağlantı Kayıtları'
  }
]);

interface GradientBgProps {
  gradientColor: string;
}

const [DefineGradientBg, GradientBg] = createReusableTemplate<GradientBgProps>();

function getGradientColor(color: CardData['color']) {
  return `linear-gradient(to bottom right, ${color.start}, ${color.end})`;
}

function handleCardClick(routePath: string) {
  if (routePath) {
    router.push(routePath);
  }
}

onMounted(async () => {
  const s = (await fetchStat()).data;
  if (s !== null) {
    stat.value = s;
  }
});
</script>

<template>
  <NCard :bordered="false" size="small" class="card-wrapper">
    <!-- define component start: GradientBg -->
    <DefineGradientBg v-slot="{ $slots, gradientColor }">
      <div
        class="rd-12px px-18px py-12px text-white shadow-md hover:shadow-xl hover:-translate-y-1 active:scale-98 transition-all duration-200 cursor-pointer select-none group"
        :style="{ backgroundImage: gradientColor }"
      >
        <component :is="$slots.default" />
      </div>
    </DefineGradientBg>
    <!-- define component end: GradientBg -->

    <NGrid cols="s:1 m:2 l:4" responsive="screen" :x-gap="16" :y-gap="16">
      <NGi v-for="item in cardData" :key="item.key">
        <GradientBg
          :gradient-color="getGradientColor(item.color)"
          class="flex-1"
          @click="handleCardClick(item.route)"
        >
          <div class="flex items-center justify-between">
            <h3 class="text-15px font-medium opacity-90">{{ item.title }}</h3>
            <span class="text-xs opacity-75 group-hover:opacity-100 group-hover:translate-x-1 transition-all">
              {{ item.subtitle }} →
            </span>
          </div>
          <div class="flex justify-between items-center pt-10px">
            <SvgIcon :icon="item.icon" class="text-36px opacity-90 group-hover:scale-110 transition-transform" />
            <CountTo
              :prefix="item.unit"
              :start-value="0"
              :end-value="item.value"
              class="text-32px font-bold text-white dark:text-dark tracking-tight"
            />
          </div>
        </GradientBg>
      </NGi>
    </NGrid>
  </NCard>
</template>

<style scoped></style>
