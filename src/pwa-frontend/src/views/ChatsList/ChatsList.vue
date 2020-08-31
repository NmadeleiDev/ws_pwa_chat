<template>
    <div>
        <DefaultAppBar :name="'Hi, ' + username"></DefaultAppBar>
        <v-main>
            <v-container fluid>
                    <CommonNotification></CommonNotification>
                    <v-list outlined shaped v-if="chats.length > 0" class="mt-0">
                    <v-subheader class="mt-0 text-lg-h4">Chats</v-subheader>
                    <v-list-item-group color="primary">
                        <v-list-item
                        v-for="(chat, i) in chats"
                        :key="i"
                        @click="setChat(chat)">
                        <v-list-item-icon>
                            <v-icon large v-if="chat.usernames.length <= 2"
                            :color="chat.messages.find(item => item.state < 3) !== undefined ? 'green' : 'dark'">person</v-icon>
                            <v-icon large v-else
                            :color="chat.messages.find(item => item.state < 3) !== undefined ? 'green' : 'dark'">people</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content>
                            <v-list-item-title class="mb-1 d-flex flex-row justify-space-between align-start">
                                <h4 class="height-auto">{{ chat.usernames.length <= 2 ? chat.usernames.filter(item => item !== username)[0] : chat.name }}</h4>
                                <p v-if="chat.usernames.length > 2">{{'(' + chat.usernames.join(', ') + ')'}}</p>
                                <v-chip v-if="getNumberOfUnread(chat)" class="height-auto">{{getNumberOfUnread(chat)}}</v-chip>
                            </v-list-item-title>
                            <v-list-item-subtitle class="d-flex flex-row justify-space-between align-start">
                                <v-subheader class="ml-0 pl-0 height-auto">{{getLastMessageText(chat)}}</v-subheader>
                                <v-subheader class="height-auto pr-0">{{getLastMessageTime(chat)}}</v-subheader>
                            </v-list-item-subtitle>
                        </v-list-item-content>
                        </v-list-item>
                    </v-list-item-group>
                    </v-list>

                    <v-sheet v-else class="mt-8 d-flex flex-column justify-space-around align-center">
                        <h3>No chats yet!</h3>
                        <v-btn large @click="$router.push('/users')" text>Find user</v-btn>
                        <v-subheader>or</v-subheader>
                        <CreateChatDialog></CreateChatDialog>
                    </v-sheet>
            </v-container>
        </v-main>
    </div>
</template>

<script lang="ts">
    import DefaultAppBar from "@/components/DefaultAppBar.vue";
    import CommonNotification from "@/components/CommonNotification.vue";
    import CreateChatDialog from "@/views/ChatsList/subcomponents/CreateChatDialog.vue";
    import Vue from 'vue';
    import {Chat, Message, User} from '@/interfaces/main';

    export default Vue.extend( {
        name: "ChatsList",
        components: {DefaultAppBar, CommonNotification, CreateChatDialog},
        methods: {
            setChat(chat: Chat) {
                this.$store.dispatch('setCurrentChat', {data: chat, isNew: false})
                this.$router.push('/chat/' + chat.id)
            },
            getLastMessageText(chat: Chat): string {
                const len = chat.messages.length
                if (len > 0) {
                    return chat.messages[len - 1].sender + ': ' + chat.messages[len - 1].text
                } else {
                    return ''
                }
            },
            getLastMessageTime(chat: Chat): string {
                const len = chat.messages.length
                if (len > 0) {
                    return this.formatDate(chat.messages[len - 1].date)
                } else {
                    return ''
                }
            },
            getNumberOfUnread(chat: Chat): string {
                const num = chat.messages.filter(item => (item.state !== 3 && item.sender !== this.username)).length
                if (num > 0) {
                    return num.toString()
                } else {
                    return ''
                }
            },
            formatDate(timestamp: number): string {
                const date = new Date(timestamp)
                return (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':' + (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes())
            },
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

<style>
.height-auto {
    height: auto;
}
</style>