<script setup>
import FileInput from '../components/form/FileInput.vue';
</script>

<template>
  <div class="page-container landing">
    <div class="brand">
      <h1>🧬</h1>
      <h1>dnAAAAA</h1>
      <div>Cari nama penyakit dari sequence DNA yang kamu punya!</div>
    </div>
    <div class="input">
      <FileInput />
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use '@/assets/styles/fonts';

.brand {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-bottom: 2rem;
}

h1 {
  margin: 0;
  font-family: fonts.$brand;
  font-size: 3rem;
}

.input {
  padding-bottom: 5rem;
  width: min(100%, 650px);
}
</style>
