<script setup>
import {
  NCard, NForm, NFormItem, NInput, NButton, NIcon,
} from 'naive-ui';
import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Add.svg';
</script>

<template>
  <div class="page-container insert">
    <NCard
      title="Tambahkan DNA Penyakit Baru"
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="insert-form-box"
    >
      <NForm>
        <NFormItem label="Nama Penyakit">
          <NInput
            placeholder="Masukkan nama penyakit"
            size="large"
          />
        </NFormItem>
        <NFormItem label="File Sequence DNA">
          <InlineFileInput />
        </NFormItem>
      </NForm>
      <template #action>
        <div class="action">
          <NButton
            type="info"
            size="large"
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
@use "@/assets/styles/layout";

.insert-form-box {
  width: layout.$form-box-width;
}

.action {
  display: flex;
  justify-content: end;
}

</style>
