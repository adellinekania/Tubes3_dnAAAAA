<script setup>
import { onMounted, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import {
  NCard, NButton, NIcon, NResult, NSpin,
} from 'naive-ui';

import { useStore } from '@/stores/search';
import SearchIcon from '@/assets/icons/Search.svg';
import axios from 'axios';

const store = useStore();

const router = useRouter();
const handleClick = () => {
  router.push({ name: 'home' });
};

const result = reactive({
  isLoading: true,
});

onMounted(async () => {
  result.isLoading = true;

  const namaPengguna = await store.data.nama;
  const prediksiPenyakit = await store.data.penyakit;
  const sequenceDNA = await store.data.file.file.arrayBuffer();
  const metode = 'KMP';
  const dnaFile = new Blob([sequenceDNA]);

  const formData = new FormData();
  formData.append('namaPengguna', namaPengguna);
  formData.append('prediksiPenyakit', prediksiPenyakit);
  formData.append('metodeStringMatching', metode);
  formData.append('sequenceDNA', dnaFile, 'sequence.txt');

  const testResult = await axios.post('/api/tesDNA', formData);

  result.isLoading = false;
  result.nama = testResult.data.Data.Nama_pengguna;
  result.penyakit = testResult.data.Data.Prediksi_penyakit;
  result.date = new Date();
  result.isMatch = testResult.data.Data.Hasil_tes;
  result.isError = false;

  store.reset();
});

const props = computed(() => {
  if (result.isLoading) {
    return {};
  }

  if (result.isError) {
    return {
      status: 'error',
      title: 'Gagal',
      description: 'Gagal melakukan pencocokan DNA penyakit',
    };
  }

  if (result.isMatch) {
    return {
      status: 'warning',
      title: 'Positif',
      description: 'Ditemukan kecocokan DNA dengan penyakit',
    };
  }

  return {
    status: 'info',
    title: 'Negatif',
    description: 'Tidak ditemukan kecocokan DNA dengan penyakit',
  };
});

</script>

<template>
  <div class="page-container result">
    <NCard
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="center-box result-box"
    >
      <template #header>
        <div class="box-header">
          Hasil Tes DNA
        </div>
      </template>
      <div v-if="result.isLoading">
        <NSpin size="large" />
      </div>
      <NResult
        v-else
        v-bind="props"
      >
        <template #footer>
          <div class="result-content">
            <table>
              <tbody>
                <tr>
                  <td>Waktu</td>
                  <td>: 2022/04/25 13:29</td>
                </tr>
                <tr>
                  <td>Nama</td>
                  <td>: {{ result.nama }}</td>
                  <td />
                </tr>
                <tr>
                  <td>Penyakit</td>
                  <td>: {{ result.penyakit }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </template>
      </NResult>
      <template #action>
        <div class="action">
          <NButton
            type="primary"
            size="large"
            :disabled="result.isLoading"
            @click="handleClick"
          >
            <template #icon>
              <NIcon>
                <SearchIcon />
              </NIcon>
            </template>
            Cari Lagi
          </NButton>
        </div>
      </template>
    </NCard>
  </div>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/box";

.result-box {
  text-align: center;
}

.result-content {
  display: flex;
  justify-content: center;

  table {
    width: 250px;
    text-align: left;
  }
}

</style>
