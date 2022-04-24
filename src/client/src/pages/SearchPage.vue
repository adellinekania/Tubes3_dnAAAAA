<script setup>
import {
  NCard, NForm, NFormItem, NInput, NButton, NIcon,
} from 'naive-ui';
import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Search.svg';
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
      <NForm>
        <NFormItem label="Nama Pasien">
          <NInput
            placeholder="Masukkan nama pasien"
            size="large"
          />
        </NFormItem>
        <NFormItem label="Prediksi Penyakit">
          <NInput
            placeholder="Masukkan nama penyakit yang diprediksi"
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
