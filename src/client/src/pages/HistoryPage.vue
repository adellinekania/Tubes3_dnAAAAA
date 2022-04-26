<script setup>
import {
  h, onMounted, reactive, ref,
} from 'vue';
import {
  NDataTable, NInputGroup, NInput, NButton, NIcon, NTag,
} from 'naive-ui';

import SearchIcon from '@/assets/icons/Search.svg';
import axios from 'axios';

const columns = [
  {
    title: 'Waktu',
    key: 'waktu',
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
  isLoading.value = true;
  let url = '/api/tesDNA';
  if (filter.value !== null && filter.value !== '') {
    url += `/${filter.value}`;
  }
  data.length = 0;
  axios.get(url).then((res) => {
    if (res.data.Data) {
      res.data.Data.forEach((d) => {
        data.push({
          nama: d.nama_pengguna,
          waktu: d.tanggal,
          prediksi: d.prediksi_penyakit,
          result: d.hasil_tes,
        });
      });
    }
    isLoading.value = false;
  }).catch((err) => {
    alert(err.response.data.Message);
    isLoading.value = false;
  });
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
@use "@/assets/styles/layout";

.history {
  justify-content: start;
}

h1 {
  text-align: center;
  margin-top: 20px;
}

.history-search {
  align-items: center;
  margin-top: 10px;
  width: min(550px, layout.$content-width);
  display: flex;
}

.table {
  width: layout.$content-width;
  max-width: 960px;
  margin-top: 25px;
}
</style>
