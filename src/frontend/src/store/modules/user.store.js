import api from "../../api/api";

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
    }),
    mutations: {
        SET_USER: (state, payload) => {
            state.username = payload.username;
            state.email = payload.email;
            state.contacts = payload.contacts;

            // initialize all chats with empty messages to let vuex watch them
            payload.chats.forEach(item => {
                item.messages = [];
            })
            state.chats = payload.chats;
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
        }
    },
    actions: {
        LOAD_USER_DATA: async (context) => {
            const response = await api.get("user");
            if (response.status !== true) {
                console.log("Error loading user data.")
                return
            }
            let user = response.data;
            context.commit("SET_USER", user);
            user.chats.forEach(item => {
                const response = api.post('messages', {chat_id: item.chat_id});
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
            });
            return true;
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
                    if (response.status === false) {
                        console.log("Error leaving chat");
                    } else {
                        location.reload();
                    }
                }
            )
        },
        TRY_JOIN_POOL: (context, payload) => {
            api.post('pool_join', payload).then(response => {
                    if (response.status === false) {
                        console.log("Error joining pool");
                    } else {
                        context.commit("SET_USER_POOL_ID", payload.poolId);
                        context.dispatch("LOAD_ALL_USERS");
                    }
                    return response.status;
                }
            )
        },
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
    }
}

export default user;