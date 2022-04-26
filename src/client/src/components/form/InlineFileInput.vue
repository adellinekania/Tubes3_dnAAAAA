<script setup>
import {
  NUpload, NUploadTrigger, NInputGroup, NButton, NInput,
} from 'naive-ui';
import { ref, computed } from 'vue';

const emit = defineEmits(['change', 'update:value']);

const props = defineProps({
  value: {
    type: Object,
    default: null,
  },
});

const fileList = ref([]);

const handleListChange = (list) => {
  if (list.length > 1) {
    list.shift();
  }

  fileList.value = [list.pop()];
};

const handleChange = (file) => {
  emit('update:value', file.file);
  emit('change', file.file);
};

if (props.value != null) {
  fileList.value.push(props.value);
}

const filename = computed(() => {
  if (fileList.value.length === 0) {
    return null;
  }

  return fileList.value[0].name;
});

</script>

<template>
  <NUpload
    :file-list="fileList"
    abstract
    :show-file-list="false"
    @change="handleChange"
    @update:file-list="handleListChange"
  >
    <NUploadTrigger
      #="{ handleClick }"
      abstract
    >
      <NInputGroup style="display: flex; max-width: 100%;">
        <NInput
          ref="textField"
          placeholder="Pilih file..."
          size="large"
          :value="filename"
        />
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
