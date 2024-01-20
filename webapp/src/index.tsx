import * as ReactDOM from 'react-dom';
import * as React from 'react';
import * as i18n from 'i18next';
import * as localeRU from './i18n/ru.json';
import * as localeEN from './i18n/en.json';
import { App } from './App';
import { initReactI18next } from 'react-i18next';

i18n.default.use(initReactI18next).init({
  supportedLngs: ['en', 'ru'],
  resources: {
    en: {
      translation: localeEN,
    },
    ru: {
      translation: localeRU,
    },
  },
  lng: 'en',
  fallbackLng: ['en', 'ru'],
  interpolation: {
    escapeValue: false,
  },
});

ReactDOM.render(<App />, document.getElementById('root'));
