<script setup>
import { onMounted, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import {
  NCard, NButton, NIcon, NResult, NSpin,
} from 'naive-ui';

import { useStore } from '@/stores/search';
import SearchIcon from '@/assets/icons/Search.svg';

const store = useStore();

const router = useRouter();
const handleClick = () => {
  router.push({ name: 'home' });
};

const result = reactive({
  isLoading: true,
});

onMounted(() => {
  // API call
  setTimeout(() => {
    result.isLoading = false;
    result.nama = store.data.nama;
    result.penyakit = store.data.penyakit;
    result.date = new Date();
    result.isMatch = true;
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
    description: 'Tidak itemukan kecocokan DNA dengan penyakit',
  };
});

</script>

<template>
  <div class="page-container result">
    <NCard
      title="Hasil Tes DNA"
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

.action {
  display: flex;
  justify-content: end;
}

</style>
