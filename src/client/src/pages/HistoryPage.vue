<script setup>
import {
  h, onMounted, reactive, ref,
} from 'vue';
import {
  NDataTable, NInputGroup, NInput, NButton, NIcon, NTag,
} from 'naive-ui';

import { api } from '@/services/api';
import EmptyStatus from '@/components/display/EmptyStatus.vue';
import SearchIcon from '@/assets/icons/Search.svg';

const columns = [
  {
    title: 'Tanggal',
    key: 'waktu',
    // sortOrder: 'descend',
    // sorter: 'default',
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
          type: row.result ? 'warning' : '',
          size: 'large',
        },
        { default: () => (row.result ? 'Positif' : 'Negatif') },
      );
    },
  },
];

const tableRef = ref(null);
const filter = ref('');
const data = reactive([]);

const fetchData = () => {
  data.isLoading = true;
  let url = '/tesDNA';
  if (filter.value !== null && filter.value !== '') {
    url += `/${filter.value}`;
  }
  data.length = 0;
  api.get(url).then((res) => {
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
    tableRef.value.sort('waktu');
    data.error = null;
  }).catch((err) => {
    data.error = err;
  }).finally(() => {
    data.isLoading = false;
  });
};

onMounted(() => {
  fetchData();
});

const handleChange = () => {
  if (filter.value !== '') {
    fetchData();
  }
};
</script>

<template>
  <div class="page-container history">
    <h1>Riwayat Cek DNA</h1>
    <div class="history-search">
      <NInputGroup>
        <NInput
          v-model:value="filter"
          placeholder="Masukkan pencarian..."
          @change="handleChange"
        />
        <NButton
          type="primary"
          size="large"
          @click="handleChange"
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
        ref="tableRef"
        :loading="data.isLoading"
        :columns="columns"
        :data="data"
        :pagination="{ pageSize: 10 }"
      >
        <template #empty>
          <EmptyStatus
            :error="data.error"
            @retry="fetchData"
          />
        </template>
      </NDataTable>
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
