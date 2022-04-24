<script setup>
import {
  NDataTable, NInputGroup, NInput, NButton, NIcon, NTag,
} from 'naive-ui';
import SearchIcon from '@/assets/icons/Search.svg';
import { h } from 'vue';

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
        { default: () => (row.result ? 'True' : 'False') },
      );
    },
  },
];

const data = [];

for (let i = 1; i <= 100; i++) {
  data.push({
    nama: `Nama ${i}`,
    waktu: new Date(),
    prediksi: `Prediksi ${i}`,
    result: i % 2 === 0,
  });
}
</script>

<template>
  <div class="page-container history">
    <h1>Riwayat Cek DNA</h1>
    <div class="history-search">
      <NInputGroup>
        <NInput placeholder="Masukkan pencarian..." />
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
