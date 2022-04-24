<script setup>
import {
  NCard, NForm, NFormItem, NInput, NButton, NIcon,
} from 'naive-ui';
import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Search.svg';
</script>

<template>
  <div class="page-container search">
    <n-card
      title="Tes DNA"
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="search-form-box"
    >
      <n-form>
        <n-form-item label="Nama Pasien">
          <n-input
            placeholder="Masukkan nama pasien"
            size="large"
          />
        </n-form-item>
        <n-form-item label="Prediksi Penyakit">
          <n-input
            placeholder="Masukkan nama penyakit yang diprediksi"
            size="large"
          />
        </n-form-item>
        <n-form-item label="File Sequence DNA">
          <inline-file-input />
        </n-form-item>
      </n-form>
      <template #action>
        <div class="action">
          <n-button
            type="info"
            size="large"
          >
            <template #icon>
              <n-icon>
                <search-icon />
              </n-icon>
            </template>
            Cocokkan DNA
          </n-button>
        </div>
      </template>
    </n-card>
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
