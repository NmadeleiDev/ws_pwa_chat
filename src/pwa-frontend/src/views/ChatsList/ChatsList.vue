<template>
    <v-app>
        <DefaultAppBar :name="'Hi, ' + username"></DefaultAppBar>
        <v-main>
            <v-row>
                <v-col>
                    <v-list shaped v-if="chats.length > 0" class="mt-0">
                    <v-subheader class="mt-0 text-lg-h4">Chats</v-subheader>
                    <v-list-item-group color="primary">
                        <v-list-item
                        v-for="(chat, i) in chats"
                        :key="i"
                        @click="setChat(chat)"
                        >
                        <v-list-item-icon>
                            <v-icon large v-if="chat.usernames.length <= 2"
                            :color="chat.messages.find(item => item.state < 3) !== undefined ? 'green' : 'dark'">person</v-icon>
                            <v-icon large v-else
                            :color="chat.messages.find(item => item.state < 3) !== undefined ? 'green' : 'dark'">people</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content>
                            <v-list-item-title>
                                <h4>{{ chat.name }}</h4>
                            </v-list-item-title>
                            <v-list-item-subtitle v-text="getLastMessage(chat)"></v-list-item-subtitle>
                        </v-list-item-content>
                        </v-list-item>
                    </v-list-item-group>
                    </v-list>

                    <v-sheet v-else class="mt-8 d-flex flex-column justify-space-around align-center">
                        <h3>No chats yet!</h3>
                        <v-btn large @click="$router.push('/users')" text>Find user</v-btn>
                    </v-sheet>
                </v-col>
            </v-row>
        </v-main>
    </v-app>
</template>

<script lang="ts">
    import DefaultAppBar from "@/components/DefaultAppBar.vue";
    import Vue from 'vue';
    import {Chat, Message, User} from '@/interfaces/main';

    export default Vue.extend( {
        name: "ChatsList",
        components: {DefaultAppBar},
        created() {
            this.$store.dispatch('initUserState')
        },
        methods: {
            setChat(chat: Chat) {
                this.$store.dispatch('setCurrentChat', {data: chat, isNew: false})
                this.$router.push('/chat')
            },
            getLastMessage(chat: Chat): string {
                let len = chat.messages.length
                if (len > 0) {
                    return chat.messages[len - 1].sender + ': ' + chat.messages[len - 1].text
                } else {
                    return ''
                }
            }
        },
        computed: {
            username(): string {
                return this.$store.getters.username()
            },
            chats(): Array<Chat> {
                return this.$store.getters.getAllChats()
            }
        }
    })
</script>

<style scoped>

</style>