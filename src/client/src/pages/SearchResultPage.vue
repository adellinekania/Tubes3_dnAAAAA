<script setup>
import { onMounted, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import {
  NCard, NButton, NIcon, NResult, NSpin,
} from 'naive-ui';

import { useStore } from '@/stores/search';
import { api } from '@/services/api';
import SearchIcon from '@/assets/icons/Search.svg';

const store = useStore();

const router = useRouter();
const handleClick = () => {
  router.push({ name: 'home' });
};

const result = reactive({
  isLoading: true,
});

onMounted(async () => {
  if (store.data.file == null) {
    router.replace({ name: 'search' });
    return;
  }

  result.isLoading = true;

  const namaPengguna = await store.data.nama;
  const prediksiPenyakit = await store.data.penyakit;
  const sequenceDNA = await store.data.file.file.arrayBuffer();
  const metode = await store.data.metode;
  const dnaFile = new Blob([sequenceDNA]);

  const formData = new FormData();
  formData.append('namaPengguna', namaPengguna);
  formData.append('prediksiPenyakit', prediksiPenyakit);
  formData.append('metodeStringMatching', metode);
  formData.append('sequenceDNA', dnaFile, 'sequence.txt');

  try {
    const testResult = await api.post('/tesDNA', formData);

    result.isLoading = false;
    result.nama = testResult.data.Data.Nama_pengguna;
    result.penyakit = testResult.data.Data.Prediksi_penyakit;
    result.kemiripan = (testResult.data.Data.Persentase_kemiripan * 100).toFixed(2);
    result.date = new Date();
    result.isMatch = testResult.data.Data.Hasil_tes;
    result.error = null;
  } catch (e) {
    result.isLoading = false;
    result.nama = store.data.nama;
    result.penyakit = store.data.penyakit;
    result.kemiripan = 0;
    result.date = new Date();
    result.error = e;
  }

  store.reset();
});

const props = computed(() => {
  if (result.isLoading) {
    return {};
  }

  if (result.error) {
    let message = 'Gagal melakukan pencocokan DNA penyakit';
    if (result.error.response?.data?.Message) {
      message += `: ${result.error.response.data.Message}`;
    }

    return {
      status: 'error',
      title: 'Gagal',
      description: message,
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
                  <td>Tanggal</td>
                  <td>:</td>
                  <td>{{ result.date.toLocaleDateString() }}</td>
                </tr>
                <tr>
                  <td>Nama</td>
                  <td>:</td>
                  <td>{{ result.nama }}</td>
                  <td />
                </tr>
                <tr>
                  <td>Penyakit</td>
                  <td>:</td>
                  <td>{{ result.penyakit }}</td>
                </tr>
                <tr>
                  <td>Kemiripan</td>
                  <td>:</td>
                  <td>{{ result.kemiripan }}%</td>
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
    text-align: left;
  }
}

</style>
