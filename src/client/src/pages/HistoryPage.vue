<script setup>
import {
  h, onMounted, reactive, ref,
} from 'vue';
import {
  NDataTable, NInputGroup, NInput, NButton, NIcon, NTag,
} from 'naive-ui';

import SearchIcon from '@/assets/icons/Search.svg';

const columns = [
  {
    title: 'Waktu',
    key: 'waktu',
    render(row) {
      return row.waktu.toISOString();
    },
  },
  {
    title: 'Nama',
    key: 'nama',
  },
  {
    title: 'Prediksi',
    key: 'prediksi',
  },
  {
    title: 'Hasil',
    key: 'result',
    render(row) {
      return h(
        NTag,
        {
          type: row.result ? 'success' : 'error',
          size: 'large',
        },
        { default: () => (row.result ? 'Positif' : 'Negatif') },
      );
    },
  },
];

const filter = ref(null);
const data = reactive([]);
const isLoading = ref(false);

const fetchData = () => {
  // API call based on filter value...
  isLoading.value = true;
  setTimeout(() => {
    for (let i = 1; i <= 100; i++) {
      data.push({
        nama: `Nama ${i}`,
        waktu: new Date(),
        prediksi: `Prediksi ${i}`,
        result: i % 2 === 0,
      });
    }
    isLoading.value = false;
  }, 2000);
};

onMounted(() => {
  fetchData();
});

const handleClick = () => fetchData();
</script>

<template>
  <div class="page-container history">
    <h1>Riwayat Cek DNA</h1>
    <div class="history-search">
      <NInputGroup>
        <NInput
          v-model:value="filter"
          placeholder="Masukkan pencarian..."
          @click="handleClick"
        />
        <NButton
          type="primary"
          size="large"
        >
          <template #icon>
            <NIcon>
              <SearchIcon />
            </NIcon>
          </template>
        </NButton>
      </NInputGroup>
    </div>
    <div class="table">
      <NDataTable
        :loading="isLoading"
        :columns="columns"
        :data="data"
        :pagination="{ pageSize: 10 }"
      />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.history {
  justify-content: start;
  // margin-top: 20px;
}

.history-search {
  margin-top: 10px;
  width: 550px;
  display: flex;
}

.table {
  width: 100%;
  max-width: 960px;
  margin-top: 25px;
}
</style>
