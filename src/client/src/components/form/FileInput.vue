<script setup>
import {
  NUpload, NUploadTrigger, NText, NIcon, NButton, NCard,
} from 'naive-ui';
import { reactive } from 'vue';
import AttachIcon from '@/assets/icons/Upload.svg';

const trigger = reactive({
  isDragging: false,
});

const handleDragEnter = (e, uploadHandler) => {
  trigger.isDragging = true;
  uploadHandler(e);
};

const handleDragLeave = (e, uploadHandler) => {
  trigger.isDragging = false;
  uploadHandler(e);
};

const emit = defineEmits(['change']);

</script>

<template>
  <n-upload
    action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
    abstract
    :max="1"
    :show-file-list="false"
    @change="emit('change', $event)"
  >
    <n-upload-trigger
      #="nHandler"
      abstract
    >
      <n-card
        :class="{ 'trigger': true, 'dragging': trigger.isDragging }"
        size="medium"
        :theme-overrides="{ paddingMedium: '0'}"
        @click="nHandler.handleClick"
        @dragover="handleDragEnter($event, nHandler.handleDragOver)"
        @dragenter="handleDragEnter($event, nHandler.handleDragEnter)"
        @dragleave="handleDragLeave($event, nHandler.handleDragLeave)"
        @drop="handleDragLeave($event, nHandler.handleDrop)"
      >
        <div class="trigger-content">
          <div>
            <n-icon
              size="48"
              :depth="3"
            >
              <attach-icon />
            </n-icon>
          </div>
          <div
            class="button-wrapper"
            style="margin-top: 10px;"
          >
            <n-button
              strong
              type="primary"
              size="large"
            >
              Unggah File DNA
            </n-button>
          </div>
          <n-text style="font-size: 1.25rem; margin-top: 10px;">
            atau geser file kamu ke area ini.
          </n-text>
          <n-text
            :depth="3"
            style="margin-top: 20px;"
          >
            <div>Pastikan file kamu adalah file teks yang berisi rangkaian DNA</div>
            <div>yang terdiri dari huruf A, G, T, dan C dalam huruf kapital.</div>
          </n-text>
        </div>
      </n-card>
    </n-upload-trigger>
  </n-upload>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/colors";

.trigger {
  background-color: colors.$grey-light;
  border-radius: 10px;
  border: 2px dashed colors.$grey-light-outline;
  padding: 50px 25px;
  cursor: pointer;

  &.dragging {
    background-color: colors.$green-light;
    border-color: colors.$green-light-outline;
  }
}

.trigger-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;
}

.button-wrapper {
  display: flex;
  justify-content: center;
}
</style>
