import { defineStore } from 'pinia';

export const useStore = defineStore('search', () => {
  const data = {
    nama: '',
    penyakit: null,
    metode: null,
    file: null,
  };

  const reset = () => {
    data.nama = '';
    data.penyakit = null;
    data.metode = '';
    data.file = null;
  };

  return { data, reset };
});
