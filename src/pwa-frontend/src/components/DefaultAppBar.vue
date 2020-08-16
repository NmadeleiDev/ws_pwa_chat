<template>
     <div>
    <v-app-bar
      color="dark accent-4"
      dark
      width="100%"
    >
      <v-toolbar-title class="ml-1">{{name}}</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-chip @click="requestNotificationsPerm()" v-if="!notifications" color="amber">Notifications are disabled</v-chip>
      <ManagePoolDialog class="mr-1"></ManagePoolDialog>

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
          <v-list-item v-if="routeParts[0] === 'chat'">
            <v-list-item-action @click.stop="dialog = true">Chat settings</v-list-item-action>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item>
            <v-list-item-action @click="$router.push('/home')">My chats</v-list-item-action>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item>
            <v-list-item-action @click="$router.push('/users')">Find users</v-list-item-action>
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item v-if="routeParts[0] !== 'chat'">
            <v-list-item-action @click="reset">Logout</v-list-item-action>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
       <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
         <v-card v-if="currentChat !== undefined">
           <v-card-title>
             <v-text-field color="green" v-model="chatName" outlined placeholder="Chat name">
               <template v-slot:append-outer>
                 <v-btn @click="saveChatName()">Save</v-btn>
               </template>
             </v-text-field>
           </v-card-title>
          <v-card-text class="d-flex flex-column justify-space-around align-center">
            <div class="d-flex flex-row justify-space-around align-center mb-2">
              <v-text-field type="number"
              placeholder="Hours to store messages" v-model="hoursToStore"></v-text-field>
              <v-btn texr @click="saveStorePeriod()">save</v-btn>
            </div>
            <div class="d-flex flex-row justify-space-around">
              <div class="d-flex flex-column justify-start align-center">
                <v-list dense>
                  <v-list-item-title>Chat members:</v-list-item-title>
                  <v-list-item  v-for="user in currentChat.usernames.concat(addedUsers)" :key="user"
                                class="d-flex fill-width flex-row justify-start align-center">
                    <h3>{{user}}
                      <v-badge color="green" v-if="user === currentChat.admin" content="Admin">
                      </v-badge>
                      <!--                    <v-btn v-else icon><v-icon>remove_circle_outline</v-icon></v-btn>-->
                    </h3>
                  </v-list-item>
                </v-list>
              </div>
              <v-divider vertical></v-divider>
              <div class="d-flex flex-column justify-start align-center">
                <v-list dense class="overflow-x-hidden overflow-y-auto">
                  <v-list-item-title>Tap on user to add:</v-list-item-title>
                  <v-list-item  v-for="(user, index) in allUsers.filter(item => !currentChat.usernames.some((el) => el === item.username) && !addedUsers.some((el) => el === item.username))" :key="index"
                                class="fill-width d-flex flex-row justify-start align-center">
                    <h4>{{user.username}}</h4>
                    <v-btn @click="addUserToChat(user.username)" class="mr-auto" icon><v-icon>add</v-icon></v-btn>
                  </v-list-item>
                </v-list>
              </div>
            </div>

          </v-card-text>
           <v-card-actions>
             <v-btn dark class="d-block ml-auto mr-auto" @click="dialog = false" >Close</v-btn>
           </v-card-actions>
         </v-card>
       </v-dialog>
  </div>
</template>

<script lang="ts">
    import Vue from 'vue';
    import {Chat, User} from "@/interfaces/main";
    import ManagePoolDialog from "@/components/ManagePoolDialog.vue";

    export default Vue.extend( {
      name: "DefaultAppBar",
      components: {ManagePoolDialog},
      props: {
          name: String
      },
      data: () => {
          return {
            notifications: false,
            dialog: false,

            chatName: '' as string,
            hoursToStore: 30,

            addedUsers: [] as Array<string>
          }
      },
      created() {
        if (this.$store.getters.username().length === 0 || !this.$store.getters.isConnected()) {
          this.$store.dispatch('initUserState')
        }
        this.requestNotificationsPerm()
      },
      mounted() {
          if (this.$store.getters.isConnected() === false) {
            this.$store.dispatch('initWebSocket')
          }

          this.hoursToStore = this.currentChat.storePeriod
      },
      methods: {
          requestNotificationsPerm() {
            Notification.requestPermission((status) =>  {
              console.log('Notification permission status:', status);
              if (status === 'granted') {
                this.notifications = true;
              }
            });
          },
          saveStorePeriod() {
            this.$store.dispatch('setCurrentChatStorePeriod', this.hoursToStore)
          },
          reset() {
              localStorage.setItem('sessionKey', '');
              localStorage.setItem('userSecret', '');
              this.$router.push("/")
          },
          saveChatName() {
            this.$store.dispatch('saveCurrentChatName', this.chatName)
            this.dialog = false
          },
          addUserToChat(username: string) {
            this.$store.dispatch('addUserToCurrentChat', username).then(ok => {
              console.log("Result is ", ok)
              if (ok)
                this.addedUsers.push(username)
            }).catch(console.warn)
          }
        },
      watch: {
          dialog: function (val: boolean) {
            if (val) {
              this.chatName = this.$store.getters.getCurrentChat().name
            }
          }
      },
      computed: {
          isOnline(): boolean {
            return this.$store.getters.isConnected()
          },
        routeParts(): Array<string> {
            return this.$route.fullPath.split("/").slice(1)
        },
        currentChat(): Chat {
            return this.$store.getters.getCurrentChat()
        },
        allUsers(): Array<User> {
            return this.$store.getters.allUsers()
        }
      }
    })
</script>

<style scoped>

</style>