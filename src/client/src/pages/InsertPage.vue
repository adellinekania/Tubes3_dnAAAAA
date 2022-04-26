<script setup>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import {
  NCard, NForm, NFormItem, NInput, NButton, NIcon,
} from 'naive-ui';

import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Add.svg';
import { useStore } from '@/stores/insert';

const store = useStore();
const input = reactive(store.data);

const rules = {
  penyakit: {
    required: true,
    message: 'Nama penyakit harus diisi',
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
    .then(() => router.push({ name: 'post' }))
    .catch(() => {});
};
</script>

<template>
  <div class="page-container insert">
    <NCard
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="center-box insert-form-box"
    >
      <template #header>
        <div class="box-header">
          Tambahkan DNA Penyakit Baru
        </div>
      </template>
      <NForm
        ref="formRef"
        :model="input"
        :rules="rules"
      >
        <NFormItem
          label="Nama Penyakit"
          path="penyakit"
        >
          <NInput
            v-model:value="input.penyakit"
            placeholder="Masukkan nama penyakit"
            size="large"
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
            Tambahkan
          </NButton>
        </div>
      </template>
    </NCard>
  </div>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/box";
</style>
