import api from "../../api/api";
import Vue from "vue";

const md5 = require('md5');

const messageType = 1
const chatType = 2

// client to server messages types
const insertMessage = 1
const updateMessage = 2
const deleteMessage = 3

// message state codes
const messageSent = 1
const messageDelivered = 2
const messageRead = 3

const generateMessageId = function () {
    return md5(Date.now() + Math.random())
}

const operational = {
    state: () => ({
        newContactName: '',
        currentChatId: 0,
    }),
    mutations: {
        SET_CURRENT_CHAT_ID(state, id) {
            state.currentChatId = id;
            console.log("Current chat set: ", id);
        },
        SET_NEW_CONTACT_NAME(state, name) {
            state.newContactName = name;
        },
    },
    actions: {
        SEND_SOCKET_MESSAGE: (context, payload) => {
            Vue.prototype.$socket.send(JSON.stringify(payload));
        },
        CHANGE_NEW_CONTACT_NAME: (context, payload) => {
            context.commit("SET_NEW_CONTACT_NAME", payload);
            console.log("Contact name set: ", context.state.newContactName);
        },
        CREATE_CHAT: async function(context, chat) {
            const res = await api.post("chat", chat);
            if (res.error !== null) {
                console.log("server error creating chat: ", res.error)
                return null;
            }
            return res.data;
        },
        CHANGE_CURRENT_CHAT: (context, payload) => {
            context.commit("SET_CURRENT_CHAT_ID", payload);
            if (payload === null) {
                return;
            }
            // marking all opened messages as read
            let messages = context.getters.GET_MESSAGES;
            if (messages === null || messages.length === 0)
                return;
            messages.forEach(message => {
                if (message.sender !== context.state.username && message.state !== messageRead) {
                    let updatedMessage = Object.assign({}, message);
                    updatedMessage.state = messageRead;
                    updatedMessage.chat_id = context.getters.GET_CURRENT_CHAT_ID;
                    context.dispatch("UPDATE_MESSAGE_STATE", updatedMessage);
                }
            });
            // обновляем последнее прочитанное сообщение
            let lastMessage = messages[messages.length - 1];
            if (lastMessage.sender !== context.state.username && lastMessage.id !== context.getters.GET_CURRENT_CHAT.lastReadMessageId) {
                lastMessage.chat_id = context.getters.GET_CURRENT_CHAT_ID;
                context.dispatch("UPDATE_LAST_READ_MESSAGE", lastMessage);
                context.commit("SET_CHAT_LAST_READ_MESSAGE_ID", lastMessage)
            }
        },
        SEND_MESSAGE: function(context, message) {
            let messageObj = {
                type: insertMessage,
                message: {
                    id: generateMessageId(),
                    sender: context.getters.GET_USER.username,
                    chat_id: context.getters.GET_CURRENT_CHAT_ID,
                    date: Date.now(),
                    state: messageSent,
                    text: message,
                    meta: null,
                    attached_file_path: "",
                }
            }
            if (context.getters.GET_CURRENT_CHAT_ID === null) {
                let chat = {
                    name: context.getters.GET_USER.username + " with " + context.getters.GET_NEW_CONTACT_NAME,
                    admin: context.getters.GET_USER.username,
                    usernames: [context.getters.GET_USER.username, context.getters.GET_NEW_CONTACT_NAME]
                }
                console.log("Creating chat: ", chat);
                context.dispatch("CREATE_CHAT", chat).then(newChat => {
                    if (newChat === null)
                        return;
                    console.log("Created chat: ", newChat);
                    messageObj.message.chat_id = newChat.chat_id;
                    context.commit("SET_CURRENT_CHAT_ID", newChat.chat_id)
                    context.dispatch("SEND_SOCKET_MESSAGE", messageObj)
                    console.log("init message sent: ", messageObj);
                    context.commit("PUSH_MESSAGE", messageObj.message)
                    context.dispatch("UPDATE_LAST_READ_MESSAGE", messageObj.message)
                });
            } else {
                context.dispatch("SEND_SOCKET_MESSAGE", messageObj)
                console.log("reg message sent: ", messageObj);
                context.commit("PUSH_MESSAGE", messageObj.message)
                context.dispatch("UPDATE_LAST_READ_MESSAGE", messageObj.message)
            }
        },
        UPDATE_MESSAGE_STATE: (context, payload) => {
            let message = {
                type: updateMessage,
                message: payload
            }
            context.dispatch("SEND_SOCKET_MESSAGE", message)
        },
        UPDATE_LAST_READ_MESSAGE: (context, message) => {
            api.post('last', message).then(res => {
                    if (res.error !== null) {
                        console.log("server error updating last read message: ", res.error)
                    }
                }
            )
            context.commit("SET_CHAT_LAST_READ_MESSAGE_ID", message)
        },
        ADD_USER_TO_CHAT: (context, payload) => {
            api.post("add", payload).then(res => {
                    if (res.error !== null) {
                        console.log("server error creating chat: ", res.error)
                    }
                }
            )
            context.commit("ADD_USER_TO_CHAT", payload);
        },
        SAVE_CHAT_NAME:  (context, payload) => {
            api.post("name", payload).then(res => {
                    if (res.error !== null) {
                        console.log("server error editing chat: ", res.error)
                    }
                }
            )
        },
        CREATE_POOL: (context, payload) => {
            return api.post('pool', payload);
        },
    },
    getters: {
        GET_CURRENT_CHAT_ID: state => {
            return state.currentChatId;
        },
        GET_NEW_CONTACT_NAME: state => {
            return state.newContactName;
        },
        // используется при первом заходе на страницу, чтобы не дать закрыть sidebar, пока не выбран чат
        IS_INTERFACE_EMPTY: state => {
            return state.newContactName.length === 0 && state.currentChatId === 0;
        }
    }
}

export default operational;
