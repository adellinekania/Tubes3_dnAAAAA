<script setup>
import { NEmpty, NButton, NIcon } from 'naive-ui';
import { computed } from 'vue';
import RefreshIcon from '@/assets/icons/Refresh.svg';

const props = defineProps({
  message: {
    type: Object,
    default: null,
  },
  error: {
    type: Object,
    default: null,
  },
});

const emit = defineEmits(['retry']);

const message = computed(() => {
  if (typeof props.message === 'string') {
    return props.message;
  }

  if (typeof props.message === 'function') {
    return props.message(props.error);
  }

  if (props.error) {
    if (props.error.response?.data?.Message) {
      return props.error.response.data.Message;
    }

    return 'Tidak dapat menghubungi server';
  }

  return 'Tidak ada data';
});
</script>

<template>
  <NEmpty
    :description="message"
    size="large"
  >
    <template #extra>
      <NButton
        type="tertiary"
        @click="emit('retry', $event)"
      >
        <template #icon>
          <NIcon>
            <RefreshIcon />
          </NIcon>
        </template>
        Coba Lagi
      </NButton>
    </template>
  </NEmpty>
</template>
