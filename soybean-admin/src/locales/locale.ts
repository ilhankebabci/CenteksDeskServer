import zhCN from './langs/zh-cn';
import enUS from './langs/en-us';
import ruRU from './langs/ru-ru';
import itIT from './langs/it-it';
import trTR from './langs/tr-tr';

const locales: Record<App.I18n.LangType, App.I18n.Schema> = {
  'tr-TR': trTR,
  'en-US': enUS,
  'zh-CN': zhCN,
  'ru-RU': ruRU,
  'it-IT': itIT
};

export default locales;
