<template>
    <b-container>
        <MobileSidebar :user-data="user"
                       :new-users="newUsers"
                       v-on:chat="setExistingChat($event.data)"
                       v-on:contact="setNewChat($event.data)"
        ></MobileSidebar>
        <b-button squared class="position-fixed fixed-top" variant="primary" v-b-toggle.sidebar-1>Меню</b-button>
        <b-row>
            <b-col>
                <h4>Чат {{ chosenChat.usernames.join(", ") }}</h4>
            </b-col>
        </b-row>
        <b-row>
            <b-col>
                <div id="messages-cont" class="messages-container">
                    <div v-for="message in chosenChatMessages" :key="message.date" :class="(message.sender === user.username ? 'right-message' : 'left-message') + ' message-container'">
                        <div class="message-header">
                            <div class="message-sender">
                                {{ message.sender }}
                            </div>
                            <div class="message-date"> {{ convertToDatetime(message.date) }} </div>
                        </div>
                        <div class="message-text">{{ message.text }}</div>
                    </div>
                </div>
            </b-col>
        </b-row>
        <b-row align-content="center">
            <b-col align-self="center">
                <div class="input-container w-75 m-auto">
                    <b-input-group>
                        <b-input class="message-input" placeholder="Введите сообщение..." type="text" v-model="messageText"></b-input>
                        <b-input-group-append>
                            <b-button variant="outline-primary" @click="sendMessage"><b-icon-envelope></b-icon-envelope></b-button>
                        </b-input-group-append>
                    </b-input-group>
                </div>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
    import MobileSidebar from "../components/MobileSidebar";

    export default {
        name: "Chat",
        components: {MobileSidebar},
        data() {
            return {
                chosenChat: {
                    chat_id: null,
                    name: '',
                    usernames: [],
                    admin: '',
                    meta: 0,
                },
                messages: [],
                messageText: '',
                findUserText: '',
                showSidebar: false,
            }
        },
        created() {
            this.$store.dispatch("LOAD_USER_DATA");
            this.$store.dispatch("LOAD_ALL_USERS");
            this.connection = this.setConnection();
        },
        mounted() {
            let messageCont = document.getElementById("messages-cont");
            messageCont.scrollTop = messageCont.scrollHeight;
        },
        beforeDestroy() {
            this.$disconnect();
        },
        methods: {
            sendMessage() {
                const message = {
                    sender: this.user.username,
                    chat_id: this.chosenChat.chat_id,
                    meta: this.chosenChat.meta,
                    date: Date.now(),
                    state: 0,
                    text: this.messageText,
                    attached_file_path: null,
                };
                this.messageText = "";
                this.$store.dispatch("SEND_MESSAGE", JSON.stringify(message));
                // this.messages.push(message);
            },
            setConnection() {
                this.$connect();
                // подписываюсь здесь, но вся логика получения сообщения в store
                this.$options.sockets.onmessage = (data) => {
                    let message = JSON.parse(data.data);
                    this.$store.dispatch("RECEIVE_MESSAGE", message);
                    // this.messages.push(message);
                    let messageCont = document.getElementById("messages-cont");
                    messageCont.scrollTop = messageCont.scrollHeight;
                };
            },
            setExistingChat(contactItem) {
                this.chosenChat = contactItem;
                this.chosenChat.meta = 0;
                this.$store.dispatch("CHANGE_CURRENT_CHAT", contactItem.chat_id);
            },
            setNewChat(recipientName) {
                this.chosenChat = {
                    chat_id: recipientName.username,
                    name: recipientName.username,
                    usernames: [recipientName.username, this.user.username],
                    admin: this.user.username,
                    meta: 1,
                };
                this.messages = [];
            },
            convertToDatetime(timestamp) {
                const date = new Date(timestamp);
                const hours = date.getHours();
                const minutes = "0" + date.getMinutes();
                const seconds = "0" + date.getSeconds();

                return hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);
            },
        },
        computed: {
            user() {
                return this.$store.getters.GET_USER;
            },
            allUsers() {
                return this.$store.getters.GET_ALL_USERS;
            },
            newUsers() {
                if (this.$store.getters.GET_ALL_USERS.length < 1) {
                    return [];
                }
                return this.$store.getters.GET_ALL_USERS.filter(item => {
                    return item.username !== this.user.username &&
                        (this.user.contacts == null ? true : !this.user.contacts.includes(item.username))
                });
            },
            chosenChatMessages() {
                let messageCont = document.getElementById("messages-cont");
                if (messageCont !== null)
                    messageCont.scrollTop = messageCont.scrollHeight;
                return this.$store.getters.GET_MESSAGES;
            },
        },
    }
</script>

<style scoped>
    p {
        padding: 0;
        margin: 0;
    }

    .messages-container {
        max-height: 70%;
        overflow-x: hidden;
        overflow-x: auto;
        display: flex;
        flex-direction: column;
        padding-bottom: 2em;
    }

    .chatroom-container {
        margin: 0 auto;
    }
    .contacts-container {
        position: relative;
        display: flex;
        flex-direction: column;
        justify-content: start;
        align-items: flex-start;
        align-content: flex-start;
        background-color: #ffd46f;
        padding: 2em;
    }
    .chat-screen {
        position: relative;
        min-width: 30em;
        display: flex;
        margin: 0 auto;
        flex-direction: column;
        justify-content: start;
        background-color: whitesmoke;
        box-shadow: 0.5em 0.5em 4em rgba(0, 0, 0, .5);
        padding: 1em 0.7em 4em;
    }
    .chat-item {
        display: flex;
        flex-direction: column;
        justify-content: center;
        height: fit-content;
    }

    .input-container {
        position: fixed;
        bottom: 1em;
        height: fit-content;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        padding: 0.5em;
    }

    .message-input {
        max-width: 100%;
        width: 75%;
    }

    .sidebar-content {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
    }

    .active-chat-container {
        display: block;
    }

    .contact-container {
        border: green solid 1px;
        border-radius: 7px;
    }

    .message-container {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        padding: 0.5em;
        margin: 0 0 5px 0;
        max-width: 55%;
        width: fit-content;
    }

    .message-header {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: fit-content;
        align-content: flex-end;
        align-items: flex-end;
    }
    .message-sender {
        padding-right: 1vw;
    }

    .right-message {
        align-self: flex-end;
    }

    .left-message {
        align-self: flex-start;
    }

    .right-message p {
        text-align: right;
    }

    .right-message .message-text {
        margin-left: auto;
    }


    .left-message p {
        text-align: left;
    }

    .left-message .message-text {
        margin-right: auto;
    }

    .message-text {
        background-color: #ffd46f;
        padding: 5px;
        width: 100%;
    }


    @media only screen
    and (max-width : 375px)
    and (max-height : 667px)
    and (-webkit-device-pixel-ratio : 2),
    only screen
    and (max-width : 375px)
    and (max-height : 812px),
    only screen
    and (max-width : 414px)
    and (max-height : 736px)
    and (-webkit-device-pixel-ratio : 3),
    only screen
    and (max-width : 411px)
    and (max-height : 731px),
    only screen
    and (max-width : 411px)
    and (max-height : 823px) {
        .chat-screen {
            position: relative;
            width: 100%;
            display: flex;
            margin: 0 auto;
            flex-direction: column;
            justify-content: start;
            background-color: whitesmoke;
            box-shadow: 0.5em 0.5em 4em rgba(0, 0, 0, .5);
            padding: 1em 0.7em 4em;
        }
    }


</style>