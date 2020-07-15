<template>
     <div>
    <v-app-bar
      color="dark accent-4"
      dense
      dark
      width="100%"
    >
      <v-toolbar-title>{{name}}</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-icon :color="isOnline ? 'green' : 'red'">offline_bolt</v-icon>

      <v-btn @click="reset()" text>
        Reset
      </v-btn>

      <v-menu
        left
        bottom
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            v-bind="attrs"
            v-on="on"
          >
        <v-icon>more_vert</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item>
            <v-list-item-action @click="$router.push('/home')">Chats</v-list-item-action>
            <v-list-item-action @click="$router.push('/users')">Find user</v-list-item-action>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend( {
        name: "DefaultAppBar",
        props: {
          name: String,
        },
      mounted() {
          if (this.isOnline === false) {
            this.$store.dispatch('initWebSocket')
          }
      },
      methods: {
            reset() {
                localStorage.setItem('sessionKey', '');
                localStorage.setItem('userSecret', '');
            }
        },
      computed: {
          isOnline(): boolean {
            return this.$store.getters.isConnected()
          }
      }
    })
</script>

<style scoped>

</style>