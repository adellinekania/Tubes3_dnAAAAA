import { createRouter, createWebHistory } from 'vue-router';

export const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/LandingPage.vue'),
  },
  {
    path: '/search',
    name: 'search',
    component: () => import('@/pages/SearchPage.vue'),
    meta: {
      label: 'Tes DNA',
      showNavlink: (route) => route.name !== 'home',
    },
  },
  {
    path: '/search/result',
    name: 'result',
    component: () => import('@/pages/SearchResultPage.vue'),
  },
  {
    path: '/new',
    name: 'new',
    component: () => import('@/pages/InsertPage.vue'),
    meta: {
      label: 'Tambah Data',
      showNavlink: () => true,
    },
  },
  {
    path: '/new/post',
    name: 'post',
    component: () => import('@/pages/InsertResultPage.vue'),
  },
  {
    path: '/history',
    name: 'history',
    component: () => import('@/pages/HistoryPage.vue'),
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
