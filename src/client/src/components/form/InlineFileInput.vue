<script setup>
import {
  NUpload, NUploadTrigger, NInputGroup, NInputGroupLabel, NButton, NText,
} from 'naive-ui';

const emit = defineEmits(['change']);
</script>

<template>
  <NUpload
    abstract
    :show-file-list="false"
    :max="1"
    @change="emit('change', $event)"
  >
    <NUploadTrigger
      #="{ handleClick }"
      abstract
    >
      <NInputGroup>
        <NInputGroupLabel
          size="large"
          style="width: 100%;"
        >
          <NText
            :depth="3"
            italic
          >
            Pilih file...
          </NText>
        </NInputGroupLabel>
        <NButton
          type="primary"
          size="large"
          @click="handleClick"
        >
          Pilih File
        </NButton>
      </NInputGroup>
    </NUploadTrigger>
  </NUpload>
</template>
