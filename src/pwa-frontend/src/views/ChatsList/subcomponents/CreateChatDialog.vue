<template>
    <v-dialog
            v-model="dialog"
            width="500"
    >
        <template v-slot:activator="{ on, attrs }">
            <v-btn large
                   v-bind="attrs"
                   v-on="on"
                   text>Create chat</v-btn>
        </template>

        <v-card>
            <v-card-title class="headline grey lighten-2">
                Create new chat room
            </v-card-title>

            <v-card-text>
                <v-text-field
                        class="mt-4"
                        label="Set chat name"
                        placeholder="Chat name"
                        v-model="name"></v-text-field>

                <v-subheader>Add users</v-subheader>
                <v-chip-group column>
                    <v-chip v-for="(user, index) in users"
                            close
                            v-on:click:close="users.splice(index, 1)">{{user}}</v-chip>
                </v-chip-group>
                <v-virtual-scroll
                :item-height="50"
                height="300"
                :items="allUsers.filter(item => !users.includes(item.username) && item.username !== user)">
                    <template v-slot="{ item }">
                        <v-list-item>
                            <v-list-item-content >
                                <div class="d-flex flex-row justify-start">
                                    <v-icon :color="item.isOnline ? 'green' : 'grey'">fiber_manual_record</v-icon>
                                    <v-list-item-title class="flex-reduce ml-3 mr-0">{{item.username}}</v-list-item-title>
                                </div>
                            </v-list-item-content>

                            <v-list-item-action>
                                <v-btn icon @click="users.push(item.username)">add</v-btn>
                            </v-list-item-action>
                        </v-list-item>
                    </template>
                </v-virtual-scroll>
            </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                        color="primary"
                        text
                        @click="createChat(); dialog = false"
                >
                    Create
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
    import Vue from 'vue'
    import {Chat, User} from "@/interfaces/main";


    export default Vue.extend({
        name: "CreateChatDialog",
        data: () => {
            return {
                dialog: false,

                name: '',
                users: [] as Array<string>
            }
        },
        methods: {
            createChat() {
                const chat: Chat = {
                    id: '',
                    messages: [],
                    admin: this.$store.getters.username(),
                    name: this.name,
                    usernames: this.users,
                    storePeriod: 24,
                }
                console.log('Creating: ', chat)
                this.$store.dispatch('createChat', chat)
            }
        },
        computed: {
            allUsers(): User[] {
                return this.$store.getters.allUsers()
            },
            user(): string {
                return this.$store.getters.username()
            }
        }
    })
</script>

<style scoped>
.flex-reduce {
    flex: 0 0 30% !important;
}
</style>