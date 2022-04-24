<script setup>
import { h, computed } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { NMenu } from 'naive-ui';
import { routes } from '@/router';

const props = defineProps({
  minimal: {
    type: Boolean,
    default: false,
  },
});

const route = useRoute();

const options = computed(() => routes
  .filter((r) => r.meta?.showNavlink(route))
  .map((r) => ({
    key: r.name,
    label: () => h(
      RouterLink,
      { to: { name: r.name } },
      { default: () => r.meta.label },
    ),
  })));

</script>

<template>
  <header :class="props.minimal ? 'header-min' : 'header'">
    <div class="header-wrapper">
      <router-link
        custom
        :to="{ name: 'home' }"
        #="{ href, navigate }"
      >
        <div
          :href="href"
          class="brand"
          @click="navigate"
        >
          <h1>🧬dnAAAAA</h1>
        </div>
      </router-link>
      <div class="navlinks">
        <NMenu
          mode="horizontal"
          :options="options"
          :value="route.name"
        />
      </div>
    </div>
  </header>
</template>

<style lang="scss" scoped>
@use "@/assets/styles/layout";
@use "@/assets/styles/fonts";

* {
  transition-duration: .3s;
  transition-timing-function: ease-in-out;
}

header {
  width: 100%;
  display: flex;
  justify-content: center;

  .header-wrapper {
    width: 100%;
    max-width: layout.$content-max-width;
    padding: .5rem 1.5rem;
    display: flex;
    align-items: center;
  }
}

h1 {
  margin: 0;
  font-family: fonts.$brand;
}

.brand {
  cursor: pointer;
  .header-min & {
    opacity: 0;
  }
}

.navlinks {
  flex-grow: 1;
  display: flex;
  justify-content: end;
}
</style>
