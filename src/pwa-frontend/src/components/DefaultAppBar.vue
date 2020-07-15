<template>
     <div>
    <v-app-bar
      color="dark accent-4"
      dark
      width="100%"
    >
      <v-toolbar-title class="ml-1">{{name}}</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-icon :color="isOnline ? 'green' : 'red'"
                  v-on="on"
                  v-bind="attrs"
                  class="mr-2"
          >offline_bolt</v-icon>
        </template>
        <span>You are {{isOnline ? 'online' : 'offline'}}</span>
      </v-tooltip>

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

        <v-list nav outlined>
          <v-list-item>
            <v-list-item-action @click="$router.push('/home')">My chats</v-list-item-action>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item>
            <v-list-item-action @click="$router.push('/users')">Find users</v-list-item-action>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item>
            <v-list-item-action @click="reset">Logout</v-list-item-action>
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
                this.$router.push("/")
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