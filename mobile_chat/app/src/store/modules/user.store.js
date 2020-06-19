import api from "../../api/api";
const sha256 = require('js-sha256').sha256;

const messageType = 1
const chatType = 2

const generateUserSecret = function (user) {
    let hash = sha256.create()
    return  hash.update([user.username, user.password, user.username].join("&-")).toString()
}

const generateToken = function (payload) {
    let hash = sha256.create()
    return hash.update([payload.username, payload.timeStamp, payload.secretKey, payload.userSecret].join("")).toString()
}

// client to server messages types
const insertMessage = 1
const updateMessage = 2
const deleteMessage = 3

// message state codes
const messageSent = 1
const messageDelivered = 2
const messageRead = 3

const user = {
    state: () => ({
        username: '',
        email: '',
        contacts: [],
        chats: [{
            chat_id: '',
            name: '',
            usernames: '',
            admin: '',
            messages: [],
            lastReadMessageId: '',
        }],
        poolId: '',
        secretKey: '',
        userSecret: '',
    }),
    mutations: {
        SET_USER: (state, payload) => {
            state.username = payload.username;
            state.email = payload.email;
            state.contacts = payload.contacts;

            // initialize all chats with empty messages to let vuex watch them
            if (payload.chats !== null && Array.isArray(payload.chats)) {
                payload.chats.forEach(item => {
                    item.messages = [];
                });
                state.chats = payload.chats;
            } else {
                state.chats = [];
            }
            state.secretKey = payload.secret_hash;
            console.log("User set: ", state);
        },
        SET_CHAT_MESSAGES: (state, payload) => {
            state.chats.find(item => item.chat_id === payload.chat_id).messages = payload.messages;
        },
        PUSH_MESSAGE (state, message)  {
            let container;
            container = state.chats.find(item => item.chat_id === message.chat_id)
            if (container.messages === undefined || container.messages === null)
                container.messages = []
            container.messages.push(message);
        },
        UPDATE_MESSAGE (state, message) {
            let messageToUpdate = state.chats.find(item => item.chat_id === message.chat_id).messages.find(item => item.id === message.id);
            if (messageToUpdate === undefined) {
                console.log("Message to update is undefined, wtf??")
            }
            messageToUpdate.state = message.state;
            messageToUpdate.text = message.text;
            console.log("Updated message ", messageToUpdate);
        },
        PUSH_CHAT (state, chat) {
            chat.messages = [];
            state.chats.push(chat);
        },
        ADD_USER_TO_CHAT(state, {chat, user}) {
            // да, мы просто добавляем юзера в локальном сторе, обновления об этом с сервера думаю нет смылса делать
            state.chats.find(item => item.chat_id === chat.chat_id).usernames.push(user.username);
        },
        SET_CHAT_LAST_READ_MESSAGE_ID (state, message) {
            state.chats.find(item => item.chat_id === message.chat_id).lastReadMessageId = message.id;
        },
        SET_USER_POOL_ID (state, poolId) {
            state.poolId = poolId;
        },
        SET_USERNAME (state, payload) {
            state.username = payload;
        },
        SET_SECRET_KEY (state, payload) {
            state.secretKey = payload;
        },
        SET_USER_SECRET (state, payload) {
            state.userSecret = payload;
        }
    },
    actions: {
        LOAD_USER_DATA: (context) => {
            let time = Date.now().toString()
            let token = generateToken({
                username: context.getters.GET_USER.username,
                timeStamp: time,
                secretKey: context.state.secretKey,
                userSecret: context.state.userSecret,
            });
            const response = api.post("user", {auth: {token: token, username: context.state.username}}, time);
            response.then(res => {
                if (res.error !== null || res.data === null) {
                    console.log("Failed to load user from api.")
                    return false;
                }
                let user = res.data;
                context.commit("SET_USER", user);
                context.state.chats.forEach(item => {
                    let currentTime = Date.now().toString();
                    let currentToken = generateToken({
                        username: context.getters.GET_USER.username,
                        timeStamp: currentTime,
                        secretKey: context.state.secretKey,
                        userSecret: context.state.userSecret,
                    });
                    const response = api.post('messages', {auth: {token: currentToken, username: context.state.username}, data: {chat_id: item.chat_id}}, currentTime);
                    response.then(res => {
                        let messages = [];
                        if (res.error !== null) {
                            console.log("Some error in load messages: ", res.error);
                            return;
                        }
                        res.data.forEach(item => {
                                messages.push(item)
                            }
                        );
                        context.commit("SET_CHAT_MESSAGES", {chat_id: item.chat_id, messages: messages})
                    }).catch(e => console.log(e));
                });
            }).catch(e => console.log(e));
        },
        RECEIVE_MESSAGE: (context, socketMessage) => {
            let container;

            switch (socketMessage.type) {
                case messageType: // TODO: message.error handling
                    // check if this message already exists
                    container = context.state.chats.find(chat => chat.chat_id === socketMessage.data.chat_id).messages.find(message => message.id === socketMessage.data.id);
                    if (container !== undefined) {
                        context.commit("UPDATE_MESSAGE", socketMessage.data);
                        return;
                    }
                    context.commit("PUSH_MESSAGE", socketMessage.data);
                    if (socketMessage.data.chat_id === context.rootGetters.GET_CURRENT_CHAT_ID && socketMessage.data.sender !== context.state.username && socketMessage.data.status !== messageRead) {
                        socketMessage.data.state = messageRead;
                        context.dispatch("UPDATE_MESSAGE_STATE", socketMessage.data);
                        context.dispatch("UPDATE_LAST_READ_MESSAGE", socketMessage.data);
                    }
                    break;
                case chatType:
                    context.commit("PUSH_CHAT", socketMessage.data);
                    break;
                default:
                    console.log("Unknown message type: ", socketMessage)
            }
        },
        LEAVE_CHAT: (context, payload) => {
            api.post("leave", payload).then(response => {
                    if (response.error !== null) {
                        console.log("Error leaving chat");
                    } else {
                        location.reload();
                    }
                }
            )
        },
        TRY_JOIN_POOL: (context, payload) => {
            api.post('pool_join', payload).then(response => {
                    if (response.error !== null) {
                        console.log("Error joining pool");
                    } else {
                        context.commit("SET_USER_POOL_ID", payload.poolId);
                        context.dispatch("LOAD_ALL_USERS");
                    }
                    return !(response.error === null);
                }
            )
        },
        // TODO: SIGN_IN и SIGN_UP - одинкаковые, соединить.
        SIGN_IN: async (context, payload) => {
            let result
            console.log("Sending signin request: ", payload);
            try {
                result = await api.post("signin", payload);
            } catch (e) {
                console.log("Error sending signin request: ", e);
                return
            }
            if (result.error === null) {
                context.commit("SET_USER", result.data);
                try {
                    payload.userSecret = generateUserSecret(payload)
                    if (await context.dispatch("SAVE_NEW_USER_TO_DB", payload) === false) {
                        console.log("Failed to signin user to db.")
                        return false;
                    }
                    context.commit("SET_USER_SECRET", payload.userSecret)
                    if (await context.dispatch("SAVE_KEY_TO_DB", {key: context.getters.GET_USER_SECRET_KEY, username: context.getters.GET_USER.username}) === true) {
                        return true;
                    } else {
                        return false;
                    }
                } catch (e) {
                    console.log(e)
                    return false;
                }
            } else {
                console.log("Error signin: ", result);
                return false;
            }
        },
        SIGN_UP: async (context, payload) => {
            let result
            console.log("Sending signup request: ", payload);
            try {
                result = await api.post("signup", payload);
            } catch (e) {
                console.log("Error sending signup request: ", e)
                return
            }
            if (result.error === null) {
                context.commit("SET_USER", result.data);
                try {
                    if (await context.dispatch("SAVE_NEW_USER_TO_DB", payload) === false) {
                        console.log("Failed to create user to db.")
                        return false;
                    }
                    context.commit("SET_USER_SECRET", payload.userSecret)
                    if (await context.dispatch("SAVE_KEY_TO_DB", {key: context.getters.GET_USER_SECRET_KEY, username: context.getters.GET_USER.username}) === true) {
                        console.log("User set: ", result.data);
                        return true;
                    } else {
                        return false;
                    }
                } catch (e) {
                    console.log(e)
                    return false;
                }
            } else {
                console.log("Error signup: ", result);
                return false;
            }
        }
    },
    getters: {
        GET_USER: state => {
            return state;
        },
        GET_CURRENT_CHAT: (state, getters) => {
            return state.chats.find(item => item.chat_id === getters.GET_CURRENT_CHAT_ID);
        },
        GET_MESSAGES: (state, getters) => {
            console.log("Message requested");
            let chat = state.chats.find(item => item.chat_id === getters.GET_CURRENT_CHAT_ID);
            if (chat !== undefined) {
                console.log("Returning messages from store: ", chat.messages);
                return chat.messages;
            }
            else
                return null;
        },
        GET_CHAT_NAME: (state, getters) => {
            console.log("ID: ", getters.GET_CURRENT_CHAT_ID)
            let chat = state.chats.find(item => item.chat_id === getters.GET_CURRENT_CHAT_ID);
            if (chat !== undefined)
                return chat.name;
            else
                return state.newContactName;
        },
        GET_USER_SECRET_KEY: state => {
            return state.secretKey;
        }
    }
}

export default user;
