<script setup>
import { useRoute } from 'vue-router';
import { computed } from 'vue';
import { NConfigProvider, NGlobalStyle } from 'naive-ui';
import NavigationBar from './components/layout/NavigationBar.vue';

/**
 * @type import('naive-ui').GlobalThemeOverrides
 */
const themeOverrides = {
  common: {
    fontFamily: "'Plus Jakarta Sans', sans-serif",
    primaryColor: '#ff8a65',
    primaryColorHover: '#ff5722',
    primaryColorPressed: '#c63f17',
    primaryColorSuppl: '#ff8a65',
  },
};

const isHome = computed(() => useRoute().name === 'home');
</script>

<template>
  <NConfigProvider
    abstract
    :theme-overrides="themeOverrides"
  >
    <NGlobalStyle />
    <NavigationBar :minimal="isHome" />
    <main>
      <RouterView #="{ Component, route }">
        <Transition
          name="fade"
          mode="out-in"
        >
          <component
            :is="Component"
            :key="route.path"
          />
        </Transition>
      </RouterView>
    </main>
  </NConfigProvider>
</template>

<style lang="scss">
@use "normalize.css/normalize.css";
@use "@/assets/styles/fonts";
@use "@/assets/styles/layout";
@use "@/assets/styles/animations";

html {
  font-size: 14px;
}

body {
  height: 100vh;
  min-height: 100vh;
  position: relative;
}

#app {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

main {
  flex-grow: 1;
  width: 100%;
  max-width: layout.$content-max-width;
  position: relative;
}
</style>
