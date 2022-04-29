<script setup>
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import {
  NCard, NForm, NFormItem, NInput, NSelect, NButton, NIcon,
} from 'naive-ui';

import { useStore as useSearchStore } from '@/stores/search';
import { api } from '@/services/api';
import InlineFileInput from '@/components/form/InlineFileInput.vue';
import EmptyStatus from '@/components/display/EmptyStatus.vue';
import SearchIcon from '@/assets/icons/Search.svg';

const searchStore = useSearchStore();

const input = reactive(searchStore.data);

const rules = {
  nama: {
    required: true,
    message: 'Nama harus diisi',
  },
  penyakit: {
    required: true,
    message: 'Prediksi penyakit harus diisi',
  },
  metode: {
    required: true,
    message: 'Metode harus diisi',
  },
  file: {
    required: true,
    message: 'File DNA harus dipilih',
  },
};

const formRef = ref(null);
const router = useRouter();

const handleClick = () => {
  formRef.value?.validate()
    .then(() => router.push({ name: 'result' }))
    .catch(() => {
    });
};

const listMetode = reactive([]);
const metodeMap = {
  KMP: 'Knuth–Morris–Pratt',
  BM: 'Boyer Moore',
};

onMounted(() => {
  // eslint-disable-next-line no-restricted-syntax, guard-for-in
  for (const key in metodeMap) {
    listMetode.push({
      label: metodeMap[key],
      value: key,
    });
  }
});

const listPenyakit = reactive([]);

const fetchData = async () => {
  listPenyakit.isLoading = true;
  try {
    const penyakit = await api.get('/penyakit');
    if (penyakit.data.Data) {
      penyakit.data.Data.forEach((d) => {
        listPenyakit.push({
          label: d.nama_penyakit,
          value: d.nama_penyakit,
        });
      });
    }
  } catch (err) {
    listPenyakit.error = err;
  } finally {
    listPenyakit.isLoading = false;
  }
};

onMounted(fetchData);

</script>

<template>
  <div class="page-container search">
    <NCard
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="center-box search-form-box"
    >
      <template #header>
        <div class="box-header">
          Tes DNA
        </div>
      </template>
      <NForm
        ref="formRef"
        :model="input"
        :rules="rules"
      >
        <NFormItem
          label="Nama"
          path="nama"
        >
          <NInput
            v-model:value="input.nama"
            placeholder="Masukkan nama kamu"
            size="large"
          />
        </NFormItem>
        <NFormItem
          label="Prediksi Penyakit"
          path="penyakit"
        >
          <NSelect
            v-model:value="input.penyakit"
            placeholder="Masukkan nama penyakit yang diprediksi"
            size="large"
            filterable
            :options="listPenyakit"
            :loading="listPenyakit.isLoading"
            :disabled="listPenyakit.isLoading"
            :input-props="{
              autocomplete: 'disabled'
            }"
          >
            <template #empty>
              <EmptyStatus
                :error="listPenyakit.error"
                @retry="fetchData"
              />
            </template>
          </NSelect>
        </NFormItem>
        <NFormItem
          label="Metode Pencocokan String"
          path="metode"
        >
          <NSelect
            v-model:value="input.metode"
            placeholder="Pilih metode pencocokan string"
            size="large"
            filterable=""
            :options="listMetode"
            :input-props="{
              autocomplete: 'disabled'
            }"
          />
        </NFormItem>
        <NFormItem
          label="File Sequence DNA"
          path="file"
        >
          <InlineFileInput v-model:value="input.file" />
        </NFormItem>
      </NForm>
      <template #action>
        <div class="action">
          <NButton
            type="info"
            size="large"
            @click="handleClick"
          >
            <template #icon>
              <NIcon>
                <SearchIcon />
              </NIcon>
            </template>
            Cocokkan DNA
          </NButton>
        </div>
      </template>
    </NCard>
  </div>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/box";
</style>
