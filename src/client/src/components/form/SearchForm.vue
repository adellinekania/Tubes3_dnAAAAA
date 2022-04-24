<script setup>
import { reactive } from 'vue';
import {
  NForm, NFormItem, NInput, NSelect,
} from 'naive-ui';

import InlineFileInput from '@/components/form/InlineFileInput.vue';

const props = defineProps({
  value: {
    type: Object,
    default: () => ({
      nama: '',
      penyakit: null,
      file: null,
    }),
  },
  penyakits: {
    type: Array,
    default: () => [],
  },
});

const model = reactive(props.value);

const rules = {
  nama: {
    required: true,
    message: 'Nama harus diisi',
  },
  penyakit: {
    required: true,
    message: 'Prediksi penyakit harus diisi',
  },
  file: {
    required: true,
    message: 'File DNA harus dipilih',
  },
};

</script>

<template>
  <NForm
    :model="model"
    :rules="rules"
  >
    <NFormItem
      label="Nama Pasien"
      path="nama"
    >
      <NInput
        v-model:value="model.nama"
        placeholder="Masukkan nama pasien"
        size="large"
      />
    </NFormItem>
    <NFormItem
      label="Prediksi Penyakit"
      path="penyakit"
    >
      <NSelect
        v-model:value="model.penyakit"
        placeholder="Masukkan nama penyakit yang diprediksi"
        size="large"
        filterable=""
        :options="props.penyakits"
        :input-props="{
          autocomplete: 'disabled'
        }"
      />
    </NFormItem>
    <NFormItem
      label="File Sequence DNA"
      path="file"
    >
      <InlineFileInput v-model:value="model.file" />
    </NFormItem>
  </NForm>
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
