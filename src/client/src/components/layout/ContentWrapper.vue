<script setup>
const props = defineProps({
  vertical: {
    type: String,
    default: 'center',
  },
  horizontal: {
    type: String,
    default: 'center',
  },
});

const style = {};
const transform = { x: 0, y: 0 };

// if (props.vertical === 'center') {
//   style.top = '50%';
//   transform.y = '-50%';
// } else if (props.vertical === 'top') {
//   style.top = 0;
// } else if (props.vertical === 'bottom') {
//   style.bottom = 0;
// }

// if (props.horizontal === 'center') {
//   style.left = '50%';
//   transform.x = '-50%';
// } else if (props.horizontal === 'left') {
//   style.left = 0;
// } else if (props.horizontal === 'right') {
//   style.right = 0;
// }

// if (transform.x !== 0 || transform.y !== 0) {
//   style.transform = `translate(${transform.x}, ${transform.y})`;
// }
</script>

<template>
  <div class="content-wrapper">
    <div
      class="content"
      :style="style"
    >
      <slot />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.content-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content {
  margin: 0 20px;
}
</style>
