<script setup>
import { useRouter } from 'vue-router';

import { useStore as useSearchStore } from '@/stores/search';
import FileInput from '@/components/form/FileInput.vue';

const router = useRouter();
const searchStore = useSearchStore();
const handleFileInput = async (file) => {
  searchStore.data.file = file;
  router.push({ name: 'search' });
};
</script>

<template>
  <div class="page-container landing">
    <div class="brand">
      <h1>🧬</h1>
      <h1>dnAAAAA</h1>
      <div>Cari nama penyakit dari sequence DNA yang kamu punya!</div>
    </div>
    <div class="input">
      <FileInput @change="handleFileInput" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/layout";
@use '@/assets/styles/fonts';

.brand {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-bottom: 2rem;
  width: layout.$content-width;
}

h1 {
  margin: 0;
  font-family: fonts.$brand;
  font-size: 3rem;
}

.input {
  padding-bottom: 5rem;
  width: min(layout.$content-width, 650px);
}
</style>
