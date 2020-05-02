import Vue from 'vue';
import Vuex from 'vuex';
import api from "../api/api";

Vue.use(Vuex);

const messageType = 1
const chatType = 2

export default new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    state: {
        user: {
            username: '',
            email: '',
            contacts: [],
            chats: [{
                chat_id: '',
                name: '',
                usernames: '',
                admin: '',
                messages: []
            }],
        },
        allUsers: [],
        socket: {
            isConnected: false,
            message: '',
            reconnectError: false,
        },
        currentChatId: 0,
    },
    getters: {
        GET_USER: state => {
            return state.user;
        },
        GET_ALL_USERS: state => {
            return state.allUsers;
        },
        GET_MESSAGES: state => {
            console.log("Message requested");
            let chat = state.user.chats.find(item => item.chat_id === state.currentChatId);
            if (chat !== undefined) {
                console.log("Returning messages from store: ", chat.messages);
                return chat.messages;
            }
            else
                return null;
        },
    },
    mutations: {
        SET_USER: (state, payload) => {
            state.user.username = payload.username;
            state.user.email = payload.email;
            state.user.contacts = payload.contacts;

            // initialize all chats with empty messages to let vuex watch them
            payload.chats.forEach(item => {
                item.messages = [];
            })
            state.user.chats = payload.chats;
            console.log("User set: ", state.user);
        },
        SET_ALL_USERS: (state, payload) => {
            state.allUsers = payload;
        },
        SET_CHAT_MESSAGES: (state, payload) => {
            state.user.chats.find(item => item.chat_id === payload.chat_id).messages = payload.messages;
        },
        SOCKET_ONOPEN (state, event)  {
            Vue.prototype.$socket = event.currentTarget
            state.socket.isConnected = true
            console.log("Connected in store!")
        },
        SOCKET_ONCLOSE (state, event)  {
            state.socket.isConnected = false
        },
        SOCKET_ONERROR (state, event)  {
            console.error("socket error in store: ", state, event)
        },
        SOCKET_ONMESSAGE (state, socketMessage)  {
            switch (socketMessage.type) {
                case messageType: // TODO: message.error handling
                    state.user.chats.find(item => item.chat_id === socketMessage.data.chat_id).messages.push(socketMessage.data);
                    break;
                case chatType:
                    state.user.chats.push(socketMessage.data)
                    break;
                default:
                    console.log("Unknown message type: ", socketMessage)
            }
        },
        // mutations for reconnect methods
        SOCKET_RECONNECT(state, count) {
            console.info(state, count)
        },
        SOCKET_RECONNECT_ERROR(state) {
            state.socket.reconnectError = true;
        },
        SET_CURRENT_CHAT(state, id) {
            state.currentChatId = id;
        }
    },
    actions: {
        LOAD_USER_DATA: (context) => {
            const response = api.get("user");
            response.then(data => {
                let user = data.data;
                context.commit("SET_USER", user);
                user.chats.forEach(item => {
                    const response = api.get('messages/' + item.chat_id);
                    response.then(data => {
                        let messages = [];
                        if (data === undefined || data === null || data.status === false) {
                            console.log("Some error in load messages");
                            return;
                        }
                        data.data.forEach(item => {
                                messages.push(item)
                            }
                        );
                        context.commit("SET_CHAT_MESSAGES", {chat_id: item.chat_id, messages: messages})
                    });
                })
            });
        },
        LOAD_ALL_USERS: (context) => {
            const response = api.get("all_users");
            response.then(data => {
                context.commit("SET_ALL_USERS", data.data);
            });
        },
        CHANGE_CURRENT_CHAT: (context, payload) => {
            context.commit("SET_CURRENT_CHAT", payload);
        },
        SEND_MESSAGE: function(context, message) {
            Vue.prototype.$socket.send(message);
        },
        RECEIVE_MESSAGE: (context, payload) => {
            context.commit("SOCKET_ONMESSAGE", payload);
        }
    },
});