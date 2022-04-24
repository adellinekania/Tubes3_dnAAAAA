<script setup>
import {
  NUpload, NUploadDragger, NText, NIcon, NButton,
} from 'naive-ui';
import { ref, computed } from 'vue';

import AttachIcon from '@/assets/icons/Upload.svg';

const emit = defineEmits(['change']);

const fileList = ref([]);

const el = ref(null);
const isDragging = computed(() => el.value?.dragOver);

</script>

<template>
  <NUpload
    ref="el"
    :file-list="fileList"
    :show-file-list="false"
    class="dragger-file-input"
    @change="emit('change', $event.file)"
    @update:file-list="fileList = []"
  >
    <NUploadDragger
      :class="{ 'dragger': true, 'dragging': isDragging }"
    >
      <div class="dragger-content">
        <div>
          <NIcon
            size="48"
            :depth="3"
          >
            <AttachIcon />
          </NIcon>
        </div>
        <div
          class="button-wrapper"
          style="margin-top: 10px;"
        >
          <NButton
            strong
            type="primary"
            size="large"
          >
            Unggah File DNA
          </NButton>
        </div>
        <NText style="font-size: 1.25rem; margin-top: 10px;">
          atau geser file kamu ke area ini.
        </NText>
        <NText
          :depth="3"
          style="margin-top: 20px;"
        >
          <div>Pastikan file kamu adalah file teks yang berisi rangkaian DNA</div>
          <div>yang terdiri dari huruf A, G, T, dan C dalam huruf kapital.</div>
        </NText>
      </div>
    </NUploadDragger>
  </NUpload>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/colors";
.dragger-file-input {
  display: flex;
  flex-direction: column;
}

.dragger {
  width: 100%;
  background-color: colors.$grey-light;
  border-radius: 10px;
  border: 2px dashed colors.$grey-light-outline;
  padding: 50px 25px;
  cursor: pointer;

  &.dragging {
    background-color: colors.$green-light;
    border: 2px dashed colors.$green-light-outline;
  }
}

.dragger-content {
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
