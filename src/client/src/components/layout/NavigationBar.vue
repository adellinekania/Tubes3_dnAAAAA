<script setup>
import {
  h, ref, watch, reactive, onMounted,
} from 'vue';
import { RouterLink, useRoute, useRouter } from 'vue-router';
import {
  NMenu, NButton, NIcon, NDrawer,
} from 'naive-ui';

import MenuIcon from '@/assets/icons/Menu.svg';

const props = defineProps({
  minimal: {
    type: Boolean,
    default: false,
  },
});

const router = useRouter();
const route = useRoute();

const options = reactive([]);
const computeOptions = () => {
  options.splice(0, options.length);
  router.getRoutes()
    .filter((r) => r.meta.showNavlink && r.meta.showNavlink(route))
    .forEach((r) => {
      options.push({
        key: r.name,
        label: () => h(
          RouterLink,
          { to: { name: r.name } },
          { default: () => r.meta.label },
        ),
      });
    });
};

onMounted(computeOptions);

const isDrawerOpen = ref(false);

watch(() => route.name, () => {
  if (isDrawerOpen.value) {
    isDrawerOpen.value = false;
    setTimeout(computeOptions, 300);
  } else {
    computeOptions();
  }
});

</script>

<template>
  <header :class="props.minimal ? 'header-min' : 'header'">
    <div class="header-wrapper">
      <RouterLink
        custom
        :to="{ name: 'home' }"
        #="{ href, navigate }"
      >
        <div
          class="brand"
          @click="navigate"
        >
          <h1>🧬dnAAAAA</h1>
        </div>
      </RouterLink>
      <div class="navlinks horizontal">
        <NMenu
          mode="horizontal"
          :options="options"
          :value="route.name"
        />
      </div>
      <div class="navlinks control">
        <NButton
          quaternary
          type="primary"
          size="medium"
          :theme-overrides="{
            heightMedium: '34px',
            paddingMedium: '0 10px',
            iconSizeMedium: '24px',
          }"
          @click="isDrawerOpen = true"
        >
          <template #icon>
            <NIcon>
              <MenuIcon />
            </NIcon>
          </template>
        </NButton>
      </div>
      <NDrawer
        v-model:show="isDrawerOpen"
        placement="top"
        height="max-content"
      >
        <NMenu
          :options="options"
          :value="route.name"
        />
      </NDrawer>
    </div>
  </header>
</template>

<style lang="scss" scoped>
@use "sass-mq";
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
}

h1 {
  margin: 0;
  font-family: fonts.$brand;
}

.header-wrapper {
  width: 100%;
  max-width: layout.$content-max-width;
  margin: .5rem 1.5rem;
  display: flex;
  align-items: center;
  position: relative;
}

.brand {
  cursor: pointer;
  .header-min & {
    opacity: 0;
  }
}

.navlinks {
  display: flex;
  justify-content: end;
  position: absolute;
  right: 0;
}

.horizontal {
  opacity: 0;
  visibility: hidden;
}

@include sass-mq.mq($from: tablet) {
  .horizontal {
    opacity: 1;
    visibility: visible;
  }

  .control {
    opacity: 0;
    visibility: hidden;
  }
}
</style>
