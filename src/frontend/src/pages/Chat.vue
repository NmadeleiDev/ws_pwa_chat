<template>
    <b-container>
        <b-sidebar id="sidebar-1" title="Chats" shadow backdrop>
            <b-container>
                <b-row class="m-3 mt-5">
                    <b-col cols="4">
                        <b-avatar variant="primary" text="BV"></b-avatar>
                    </b-col>
                    <b-col cols="6">
                        <h3>{{ userData.username }}</h3>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col class="m-3">
                        <b-button class="mb-1" v-b-toggle.all-chats-cont variant="info" block>Активные чаты</b-button>
                        <!--                а вот список контактов потом дополнится не просто список имен, а еще и есть ли непрочитанные сообщения и все в этом роде-->
                        <!--                            <p @click="setExistingChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">{{ chatItem.name }}</p>-->
                        <b-collapse id="all-chats-cont">
                        <b-list-group v-if="userData.chats.length > 0">
                            <b-list-group-item @click="setExistingChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">
                                {{ chatItem.name }}
                            </b-list-group-item>
                        </b-list-group>
                        <p v-else>Чатов пока нет...</p>
                        </b-collapse>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col class="m-3">
                        <b-button class="mb-1" v-b-toggle.all-users-cont variant="info" block>Найти пользователя</b-button>
                        <b-collapse id="all-users-cont">
                            <b-input v-model="findUserText" placeholder="Поиск пользователей..."></b-input>
                            <b-list-group v-if="newUsers.length > 0">
                                <b-list-group-item class="flex-row align-items-start" @click="setNewChat(contactItem)" v-for="contactItem in newUsers.filter(item => item.username.includes(findUserText) > 0)" :key="contactItem.username">
                                    <div class="d-flex flex-row align-content-center justify-content-around">
                                        <div class="align-self-lg-start"><b-avatar></b-avatar></div>
                                        <div><h4>{{ contactItem.username }}</h4></div>
                                    </div>
                                </b-list-group-item>
                            </b-list-group>
                        </b-collapse>
                    </b-col>
                </b-row>
            </b-container>
            <template v-slot:footer="{ hide }">
                <div class="d-flex bg-dark text-light justify-content-around align-items-center px-3 py-2">
                    <b-button variant="outline-warning" @click="signOut">Выйти</b-button>
                    <b-button size="md" @click="hide">Закрыть</b-button>
                </div>
            </template>
        </b-sidebar>
        <b-button squared class="position-absolute fixed-top" variant="primary" v-b-toggle.sidebar-1>Меню</b-button>
        <b-row>
            <b-col>
                <h4>Чат с {{ chat.usernames.join(", ") }}</h4>
            </b-col>
        </b-row>
        <b-row>
            <b-col>
                <div v-for="message in messages" :key="message.date" :class="(message.sender === userData.username ? 'right-message' : 'left-message') + ' message-container'">
                    <div class="message-header">
                        <div class="message-sender">
                            {{ message.sender }}
                        </div>
                        <div class="message-date"> {{ convertToDatetime(message.date) }} </div>
                    </div>
                    <div class="message-text">{{ message.text }}</div>
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
                findUserText: '',
                showSidebar: false,
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
                this.messageText = "";
                this.$store.dispatch("SEND_MESSAGE", JSON.stringify(message));
                // this.messages.push(message);
            },
            setConnection() {
                this.$connect();
                this.$options.sockets.onmessage = (data) => {
                    let message = JSON.parse(data.data)
                    this.$store.dispatch("RECEIVE_MESSAGE", message);
                    this.messages.push(message);
                };
            },
            setExistingChat(contactItem) {
                this.chat = contactItem;

                const response = api.get('get_messages/' + contactItem.chat_id);
                const that = this;
                this.messages = [];

                response.then(data => {
                    if (data === undefined || data === null || data.status === false || data.data.length === 0) {
                        return;
                    }
                    data.data.forEach(item => {
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
            },
            signOut() {
                api.get("signout");
                this.$router.push("/login");
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
                if (this.$store.getters.GET_ALL_USERS.length < 1) {
                    return [];
                }
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