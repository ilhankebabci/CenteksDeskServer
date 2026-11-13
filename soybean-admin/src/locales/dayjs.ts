import { locale } from 'dayjs';
import 'dayjs/locale/tr';
import 'dayjs/locale/zh-cn';
import 'dayjs/locale/en';
import 'dayjs/locale/ru';
import 'dayjs/locale/it';
import { localStg } from '@/utils/storage';

/**
 * Set dayjs locale
 *
 * @param lang
 */
export function setDayjsLocale(lang: App.I18n.LangType = 'tr-TR') {
  const localMap = {
    'tr-TR': 'tr',
    'zh-CN': 'zh-cn',
    'en-US': 'en',
    'ru-RU': 'ru',
    'it-IT': 'it'
  } satisfies Record<App.I18n.LangType, string>;

  const l = lang || localStg.get('lang') || 'tr-TR';

  locale(localMap[l]);
}
