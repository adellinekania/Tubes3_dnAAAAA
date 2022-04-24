<script setup>
import {
  NUpload, NUploadTrigger, NInputGroup, NInputGroupLabel, NButton, NText,
} from 'naive-ui';

const emit = defineEmits(['change']);
</script>

<template>
  <n-upload
    abstract
    :show-file-list="false"
    :max="1"
    @change="emit('change', $event)"
  >
    <n-upload-trigger
      #="{ handleClick }"
      abstract
    >
      <n-input-group>
        <n-input-group-label
          size="large"
          style="width: 100%;"
        >
          <n-text
            :depth="3"
            italic
          >
            Pilih file...
          </n-text>
        </n-input-group-label>
        <n-button
          type="primary"
          size="large"
          @click="handleClick"
        >
          Pilih File
        </n-button>
      </n-input-group>
    </n-upload-trigger>
  </n-upload>
</template>
