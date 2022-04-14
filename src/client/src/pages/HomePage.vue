<template>
  <div class="wrapper">
    <h3 v-if="message !== null">{{ message }}</h3>
    <button class="btn-test" @click="testApi">TEST</button>
  </div>
</template>
<script>
export default {
  name: "HomePage",
  data() {
    return {
      message: null
    }
  },
  methods: {
    testApi() {
      this.$http._get("/tesDNA").then(res => {
        this.message = res.Message;
      })
    }
  }
};
</script>
<style scoped>
.btn-test {
  background: #00adb5;
  height: 4rem;
  padding: 8px 16px;
  color: #222831;
}

.wrapper {
  background-color: #fdf9f7;
  display: flex;
  flex-direction: column;
  width: 40vw;
  height: 80vh;
  border-radius: 4px;
  box-shadow: 5px 5px 8px 5px rgba(255, 226, 226, 0.48);
  justify-content: center;
  align-items: center;
}

h3 {
  color: #222831;
}

@media only screen and (max-width: 886px) {
  .wrapper {
    width: 50vw;
  }
}

@media only screen and (max-width: 570px) {
  .wrapper {
    width: 60vw;
  }
}

@media only screen and (max-width: 375px) {
  .wrapper {
    width: 80vw;
  }
}

</style>
