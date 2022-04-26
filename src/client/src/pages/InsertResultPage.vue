<script setup>
import { onMounted, reactive, computed } from 'vue';
import {
  NCard, NResult, NSpin,
} from 'naive-ui';

import { useStore } from '@/stores/insert';
import axios from 'axios';

const store = useStore();
const result = reactive({
  isLoading: true,
});

onMounted(async () => {
  result.isLoading = true;

  const namaPenyakit = await store.data.penyakit;
  const sequenceDNA = await store.data.file.file.arrayBuffer();
  const dnaFile = new Blob([sequenceDNA]);

  const formData = new FormData();
  formData.append('namaPenyakit', namaPenyakit);
  formData.append('sequenceDNA', dnaFile, 'sequence.txt');

  try {
    const testResult = await axios.post('/api/upload', formData);

    result.isLoading = false;
    result.penyakit = testResult.data.Data.Nama_penyakit;
    result.date = new Date();
    result.isError = false;
  } catch (e) {
    result.isLoading = false;
    result.penyakit = store.data.penyakit;
    result.date = new Date();
    result.isError = true;
  }

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
      description: 'Data penyakit gagal dimasukkan',
    };
  }

  return {
    status: 'success',
    title: 'Berhasil',
    description: 'Data penyakit berhasil dimasukkan',
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
          Tambahkan DNA Penyakit Baru
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
                  <td>: {{ result.date.toLocaleDateString() }}</td>
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
