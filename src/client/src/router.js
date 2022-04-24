import { createRouter, createWebHistory } from 'vue-router';
import LandingPage from '@/pages/LandingPage.vue';
import SearchPage from '@/pages/SearchPage.vue';
import InsertPage from '@/pages/InsertPage.vue';
import HistoryPage from '@/pages/HistoryPage.vue';

export const routes = [
  {
    path: '/',
    name: 'home',
    component: LandingPage,
  },
  {
    path: '/search',
    name: 'search',
    component: SearchPage,
    meta: {
      label: 'Tes DNA',
      showNavlink: (route) => route.name !== 'home',
    },
  },
  {
    path: '/new',
    name: 'new',
    component: InsertPage,
    meta: {
      label: 'Masukkan Baru',
      showNavlink: () => true,
    },
  },
  {
    path: '/history',
    name: 'history',
    component: HistoryPage,
    meta: {
      label: 'Riwayat',
      showNavlink: () => true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
