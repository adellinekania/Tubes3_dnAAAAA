<script setup>
import { onMounted, reactive, computed } from 'vue';
import {
  NCard, NResult, NSpin,
} from 'naive-ui';

import { useStore } from '@/stores/insert';

const store = useStore();
const result = reactive({
  isLoading: true,
});

onMounted(() => {
  // API call
  setTimeout(() => {
    result.isLoading = false;
    result.penyakit = store.data.penyakit;
    result.date = new Date();
    result.isError = false;

    store.reset();
  }, 2000);
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
      title="Tambahkan DNA Penyakit Baru"
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="result-box"
    >
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
@use "@/assets/styles/layout";

.result-box {
  width: layout.$form-box-width;
  margin: layout.$content-margin;
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
