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
}

html {
  width: 100%;
  margin: 0;
  padding: 0;
}
</style>
