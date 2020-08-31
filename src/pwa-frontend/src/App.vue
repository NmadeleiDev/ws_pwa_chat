<template>
  <v-app>
    <router-view></router-view>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'App',

  components: {
  },

  data: () => ({
  }),
});
</script>

<style>

body {
  margin: 0;
  padding: 0;
  max-height: 100vh;
  height: 100vh;
}

html {
  width: 100%;
  max-height: 100vh;
  height: 100vh;
  margin: 0;
  padding: 0;
}
</style>
