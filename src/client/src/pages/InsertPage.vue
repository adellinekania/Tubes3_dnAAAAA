<script setup>
import {
  NCard, NForm, NFormItem, NInput, NButton, NIcon,
} from 'naive-ui';
import InlineFileInput from '@/components/form/InlineFileInput.vue';
import SearchIcon from '@/assets/icons/Add.svg';
</script>

<template>
  <div class="page-container insert">
    <n-card
      title="Tambahkan DNA Penyakit Baru"
      :segmented="{
        content: true,
        footer: 'soft'
      }"
      class="insert-form-box"
    >
      <n-form>
        <n-form-item label="Nama Penyakit">
          <n-input
            placeholder="Masukkan nama penyakit"
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
            Tambahkan
          </n-button>
        </div>
      </template>
    </n-card>
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
