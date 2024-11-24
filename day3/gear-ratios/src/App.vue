<script setup lang="ts">
import HelloWorld from './components/Blurb.vue';
import GearRatiosSolver from './components/GearRatiosSolver.vue';
</script>

<template>
  <header>
    <v-icon icon="mdi-cog" class="logo"/>
    <div class="wrapper">
      <HelloWorld msg="Gear Ratios" />
    </div>
  </header>
  <main>
    <GearRatiosSolver />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
    font-size: 10rem;
    color: aquamarine
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
