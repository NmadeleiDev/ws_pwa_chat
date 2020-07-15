<template>
    <v-app  class="position-relative">
        <DefaultAppBar :name="chat !== undefined ? chat.name : 'Error loading chat'"></DefaultAppBar>
        <v-main v-if="chat !== undefined" class="fill-height">
            <v-row class="fill-height">
                <v-col class="fill-height">
                    <div class="fill-height w-100 d-flex flex-column justify-space-between align-center">
                        <div v-if="chat.messages.length > 0" class="mt-0 fill-height fill-width d-flex flex-column">
                            <v-sheet
                                    rounded
                                    color="grey lighten-2"
                                    :class="'pl-3 pb-1 mb-2 d-flex flex-column justify-space-around align-start ' +
                                    (message.sender === username ? 'align-self-end mr-4' : 'align-self-start ml-4')"
                                    v-for="(message, i) in chat.messages"
                                    :key="i">
                                <div class="d-flex flex-row justify-space-between mb-0">
                                    <v-subheader>{{chat.usernames.length > 2 ? message.sender: ''}}</v-subheader>
                                    <v-subheader>{{formatDate(message.date)}}</v-subheader>
                                </div>
                                <div class="wrap">
                                    {{message.text}}
                                </div>
                                <v-icon small
                                        class="align-self-end pr-2"
                                        :color="getCheckColor(message)">check</v-icon>
                            </v-sheet>
                        </div>
                        <v-sheet v-else class="mt-8 d-flex flex-column justify-space-around align-center">
                            <h4>No messages yet!</h4>
                        </v-sheet>
                        <v-sheet dark class="message-input">
                            <v-textarea
                                    class="w-100 pl-4 mr-4 ml-0"
                                    filled
                                    auto-grow
                                    type="text"
                                    color="green"
                                    rows="1"
                                    v-model="message"
                            >
                                <template v-slot:append-outer>
                                    <v-btn icon @click="sendMessage()">
                                        <v-icon large :color="message === '' ? '' : 'green'">send</v-icon>
                                    </v-btn>
                                </template>
                            </v-textarea>
                        </v-sheet>
                    </div>
                </v-col>
            </v-row>
        </v-main>
    </v-app>
</template>

<script lang="ts">
    import {Chat, Message, User} from "@/interfaces/main.ts";
    import DefaultAppBar from "@/components/DefaultAppBar.vue";
    import Vue from 'vue';

    export default Vue.extend({
        name: "Chat",
        components: {
            DefaultAppBar
        },
        data: () => {
            return {
                message: '',
            }
        },
        methods: {
            sendMessage() {
                this.$store.dispatch('sendMessageInChat', this.message)
                this.message = ''
            },
            formatDate(timestamp: number): string {
                const date = new Date(timestamp)
                return (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':' + (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes())
            },
            getCheckColor(message: Message): string {
                switch (message.state) {
                    case 1:
                        return 'grey lighten-2'
                    case 2:
                        return 'grey darken-4'
                    case 3:
                        return 'green'
                    default:
                        return 'grey lighten-2'
                }
            }
        },
        computed: {
            chat(): Chat | undefined {
                return this.$store.getters.getCurrentChat()
            },
            username(): string {
                return this.$store.getters.username();
            }
        }
    })
</script>

<style scoped>
.message-input {
    position: fixed;
    bottom: 0;
    width: 100%;
}

    .fill-width {
        width: 100%;
    }
</style>