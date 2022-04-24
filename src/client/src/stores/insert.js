import { defineStore } from 'pinia';

export const useStore = defineStore('insert', () => {
  const data = {
    penyakit: '',
    file: null,
  };

  const reset = () => {
    data.penyakit = '';
    data.file = null;
  };

  return { data, reset };
});
