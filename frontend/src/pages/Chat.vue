<template>
    <div>
        <h3>{{ chat.name }}</h3>
        <div class="chatroom-container">
            <div class="sidebar-container">
                <b-button class="main-info-button" size="lg" v-b-toggle.sidebar-1 variant="primary"><b-icon-chat-dots></b-icon-chat-dots></b-button>
                <b-sidebar id="sidebar-1" :title="userData.username" shadow>
                    <div class="sidebar-content">
                        <div class="contacts-container">
                            <h4>My chats:</h4>
                            <!--                а вот список контактов потом дополнится не просто список имен, а еще и есть ли непрочитанные сообщения и все в этом роде-->
                            <p @click="setExistingChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">{{ chatItem.name }}</p>
                            <div class="chat-item" @click="setExistingChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">
<!--                                <h5>{{ chatItem.name }}</h5>-->
                                <div>{{ chatItem.usernames.filter(item => item !== userData.username).join(", ") }}</div>
                            </div>
                        </div>
                        <div class="contacts-container">
                            <h4>All users:</h4>
                            <p @click="setNewChat(contactItem)" v-for="contactItem in newUsers" :key="contactItem.username">{{ contactItem.username }}</p>
                        </div>
                    </div>
                </b-sidebar>
            </div>
            <div class="chat-screen">
                <div v-for="message in messages" :key="message.date" :class="(message.sender === userData.username ? 'right-message' : 'left-message') + ' message-container'">
                    <div class="message-header">
                        <div class="message-sender">
                            {{ message.sender }}
                        </div>
                        <div class="message-date"> {{ convertToDatetime(message.date) }} </div>
                    </div>
                    <div class="message-text">{{ message.text }}</div>
                </div>
                <div class="input-container">
                    <input class="message-input" type="text" v-model="messageText">
                    <button @click="sendMessage">Send</button>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    import api from "../api/api";

    export default {
        name: "Chat",
        data() {
            return {
                chat: {
                    chat_id: null,
                    name: '',
                    usernames: [],
                    admin: '',
                },
                messages: [],
                messageText: '',
                connection: {},
            }
        },
        created() {
            this.$store.dispatch("LOAD_USER_DATA");
            this.$store.dispatch("LOAD_ALL_USERS");
            this.connection = this.setConnection();
        },
        methods: {
            sendMessage() {
                const message = {
                    sender: this.userData.username,
                    chat_id: this.chat.chat_id,
                    meta: this.chat.name,
                    date: Date.now(),
                    state: 0,
                    text: this.messageText,
                    attached_file_path: null,
                };
                console.log("Message to send: ", message);
                this.connection.send(JSON.stringify(message));
                this.messageText = "";
                this.messages.push(message);
            },
            setConnection() {
                const that = this;
                if (window["WebSocket"]) {
                    const conn = new WebSocket("ws://" + window.location.host + "/ws/connect");
                    conn.onclose = function () {
                        console.log("Connection closed");
                    };
                    conn.onmessage = function (evt) {
                        console.log("Got message: ", evt);
                        let messageBody = JSON.parse(evt.data);
                        console.log("Body: ", messageBody);
                        if (messageBody.sender !== that.$store.getters.GET_USER.username) {
                            that.messages.push(messageBody);
                        }
                    };
                    setInterval(() => conn.send(JSON.stringify({meta: "pong"})), 20000);

                    return conn;
                }
                else {
                    alert("You are a fucking looser, your browser does not support websockets");
                }
            },
            setExistingChat(contactItem) {
                this.chat = contactItem;

                const response = api.get('get_messages/' + contactItem.chat_id);
                const that = this;
                this.messages = [];

                response.then(data => {
                    if (data === undefined || data.length === 0) {
                        return;
                    }
                    data.forEach(item => {
                            that.messages.push(item)
                        }
                    );
                });
            },
            setNewChat(recipientName) {
                this.chat = {
                    chat_id: null,
                    name: recipientName.username,
                    usernames: [recipientName.username, this.userData.username],
                    admin: this.userData.username,
                };
                this.messages = [];
            },
            convertToDatetime(timestamp) {
                const date = new Date(timestamp);
                const hours = date.getHours();
                const minutes = "0" + date.getMinutes();
                const seconds = "0" + date.getSeconds();

                return hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);
            }
        },
        computed: {
            userData() {
                return this.$store.getters.GET_USER;
            },
            allUsers() {
                return this.$store.getters.GET_ALL_USERS;
            },
            newUsers() {
                return this.$store.getters.GET_ALL_USERS.filter(item => {
                    return item.username !== this.userData.username &&
                        (this.userData.contacts == null ? true : !this.userData.contacts.includes(item.username))
                });
            }
        },
    }
</script>

<style scoped>
    p {
        padding: 0;
        margin: 0;
    }

    .main-info-button {
        position: fixed;
        top: 1em;
        left: 1em;
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
        border: solid green 1px;
        width: 95%
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