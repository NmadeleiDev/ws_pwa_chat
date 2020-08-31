<template>
    <div class="position-relative d-flex flex-column height-screen">
        <DefaultAppBar :name="chat !== undefined ? chat.name : 'Error loading chat'"></DefaultAppBar>
        <CommonNotification></CommonNotification>
        <v-main>
            <v-container fluid>
                <div v-if="chat !== undefined" class="d-flex flex-column">
                    <div class="d-flex flex-column flex-grow-1">
                        <div class="d-flex flex-column flex-grow-1">
                            <div class="flex-grow-1 w-100 d-flex flex-column justify-space-between align-center">
                                <div
                                        v-if="chat.messages.length > 0"
                                        class="mt-0 overflow-y-auto flex-grow-1 fill-width d-flex flex-column-reverse"
                                        id="messages-cont">
                                    <v-sheet
                                            rounded
                                            color="grey lighten-2"
                                            :class="'pl-3 pb-1 mb-2 d-flex flex-column justify-space-around align-start ' +
                                            (message.sender === username ? 'align-self-end mr-4' : 'align-self-start ml-4')"
                                            v-for="(message, i) in chat.messages.concat([]).reverse()"
                                            :key="i">
                                        <div class="mt-2 mb-1 d-flex flex-row justify-space-between">
                                            <v-subheader class="pl-0 ml-0 height-auto">{{chat.usernames.length > 2 ? message.sender: ''}}</v-subheader>
                                            <v-subheader class="height-auto">{{formatDate(message.date)}}</v-subheader>
                                        </div>
                                        <div v-if="message.attachedFileIds && message.attachedFileIds.length > 0" class="d-flex flex-column align-center justify-space-around">
                                            <v-img max-width="210px" class="mr-2 mb-2" v-for="id in message.attachedFileIds" :src="getBase64(id)"></v-img>
                                        </div>
                                        <div class="wrap">
                                            {{message.text}}
                                        </div>
                                        <v-icon v-if="message.sender === username" small
                                                class="align-self-end pr-2"
                                                :color="getCheckColor(message)">check</v-icon>
                                    </v-sheet>
                                </div>
                                <v-sheet v-else class="mt-8 flex-grow-1 d-flex flex-column justify-space-around align-center">
                                    <h4>No messages yet!</h4>
                                </v-sheet>
                            </div>
                        </div>
                    </div>
                </div>
            </v-container>
        </v-main>
        <v-footer app class="fill-width pa-0">
            <v-sheet dark class="message-input fill-width">
                <v-textarea
                        class="w-100 pl-4 mr-4 ml-0 input-color"
                        solo
                        light
                        auto-grow
                        type="text"
                        rows="1"
                        v-model="message"
                >
                    <template v-slot:prepend>
                        <SendFileDialog></SendFileDialog>
                    </template>
                    <template v-slot:append-outer>
                        <v-btn light class="pb-1" icon @click="sendMessage()">
                            <v-icon light large :color="message === '' ? 'white' : 'green'">send</v-icon>
                        </v-btn>
                    </template>
                </v-textarea>
            </v-sheet>
        </v-footer>
    </div>
</template>

<script lang="ts">
    import {Chat, Message, User} from "@/interfaces/main.ts";
    import DefaultAppBar from "@/components/DefaultAppBar.vue";
    import CommonNotification from "@/components/CommonNotification.vue";
    import SendFileDialog from "@/views/Chat/subcomponents/SendFileDialog.vue";
    import Vue from 'vue';

    export default Vue.extend({
        name: "Chat",
        components: {
            DefaultAppBar,
            CommonNotification,
            SendFileDialog
        },
        data: () => {
            return {
                message: '',
            }
        },
        created() {
            if (this.chat === undefined && !this.$store.getters.isNew()) {
                const chatId = this.$route.params.id
                this.$store.dispatch('setCurrentChatId', chatId)
            }
        },
        mounted() {
            this.scrollDown()
        },
        updated() {
            this.scrollDown()
        },
        methods: {
            scrollDown() {
                let messDiv = document.getElementById("messages-cont");
                if (messDiv) {
                    let chld = messDiv.firstElementChild;
                    if (chld)
                        chld.scrollIntoView({behavior: "smooth"});
                }
            },
            sendMessage() {
                this.$store.dispatch('sendMessageInChat', {message: this.message, files: null})
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
            },
            getBase64(id: string) {
                let blob = this.loadedFiles.find(item => item.id === id)
                if (!blob)
                    return ''
                let urlCreator = window.URL || window.webkitURL;
                return urlCreator.createObjectURL(blob.file);
            }
        },
        computed: {
            chat(): Chat | undefined {
                // console.log("Current: ", this.$store.getters.getCurrentChat())
                let chat = this.$store.getters.getCurrentChat()
                if (chat) {
                    this.$store.dispatch('loadChatImages', chat)
                    this.scrollDown()
                }
                return chat
            },
            username(): string {
                return this.$store.getters.username()
            },
            loadedFiles(): Array<{id: string, file: Blob}> {
                return this.$store.getters.getLoadedFiles()
            }
        }
    })
</script>

<style scoped>
.message-input {
    /*position: fixed;*/
    /*bottom: 0;*/
    /*right: 0;*/
    /*left: 0;*/
    width: 100%;
    padding: 8px 0 0 0;
    margin: 0;
}

    .fill-width {
        width: 100%;
    }
.height-auto {
    height: auto;
}

.height-screen {
    max-height: 100vh;
}


    .input-color {
        color: black;
    }
</style>