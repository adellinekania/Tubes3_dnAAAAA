<template>
  <div class="not-found">
    <h1>404 Not Found</h1>
    <div>
      <router-link to="/">
        <a href="">Home</a>
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotFound"
};
</script>

<style scoped>
.not-found {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #222831;
  font-family: "Montserrat", sans-serif;
}
h1 {
  color: #ffffff;
  text-transform: uppercase;
  font-size: 2rem;
}
a {
  color: #59e4f6;
  font-size: 1.2rem;
}
</style>
