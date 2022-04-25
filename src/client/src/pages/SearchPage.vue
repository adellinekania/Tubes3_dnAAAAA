<script setup>
import {onMounted, reactive, ref} from 'vue';
import {useRouter} from 'vue-router';
import {
  NCard, NForm, NFormItem, NInput, NSelect, NButton, NIcon,
} from 'naive-ui';

import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Search.svg';
import {useStore as useSearchStore} from '@/stores/search';
import axios from "axios";

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
      .then(() => router.push({name: 'result'}))
      .catch(() => {
      });
};

const listPenyakit = reactive([]);
const listMetode = reactive([]);
const metodeMap = {
  "KMP": "Knuth–Morris–Pratt",
  "BM": "Boyer Moore"
}

onMounted(async () => {
  for (let key in metodeMap) {
    listMetode.push({
      label: metodeMap[key],
      value: key
    })
  }

  const penyakit = await axios.get('/api/penyakit')
  if (penyakit.data.Data) {
    penyakit.data.Data.forEach(d => {
      listPenyakit.push({
        label: d.nama_penyakit,
        value: d.nama_penyakit
      })
    })
  }
});

</script>

<template>
  <div class="page-container search">
    <NCard
        title="Tes DNA"
        :segmented="{
        content: true,
        footer: 'soft'
      }"
        class="search-form-box"
    >
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
              filterable=""
              :options="listPenyakit"
              :input-props="{
              autocomplete: 'disabled'
            }"
          />
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
          <InlineFileInput v-model:value="input.file"/>
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
                <SearchIcon/>
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
@use "@/assets/styles/layout";

.search-form-box {
  width: layout.$form-box-width;
  margin: layout.$content-margin;
}

.action {
  display: flex;
  justify-content: end;
}

</style>
