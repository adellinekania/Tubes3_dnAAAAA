<template>
  <div class="base-layout">
    <slot/>
  </div>
</template>
<script>
export default {
  name: "BaseLayout"
};
</script>
<style scoped>
.base-layout {
  background: url("../images/background.jpg");
  height: 100vh;
  width: 100vw;
  color: #eeeeee;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: "Montserrat", sans-serif;
}
</style>
