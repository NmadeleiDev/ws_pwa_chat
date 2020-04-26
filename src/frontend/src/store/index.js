import Vue from 'vue';
import Vuex from 'vuex';
import api from "../api/api";

Vue.use(Vuex);

export const store = new Vuex.Store({
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
            }],
        },
        allUsers: [],
        socket: {
            isConnected: false,
            message: '',
            reconnectError: false,
        },
    },
    getters: {
        GET_USER: state => {
            return state.user;
        },
        GET_ALL_USERS: state => {
            return state.allUsers;
        }
    },
    mutations: {
        SET_USER: (state, payload) => {
            state.user = payload;
            console.log("User set: ", state.user);
        },
        SET_ALL_USERS: (state, payload) => {
            state.allUsers = payload;
            console.log("All users set: ", state.allUsers);
        },
        SOCKET_ONOPEN (state, event)  {
            Vue.prototype.$socket = event.currentTarget
            state.socket.isConnected = true
        },
        SOCKET_ONCLOSE (state, event)  {
            state.socket.isConnected = false
        },
        SOCKET_ONERROR (state, event)  {
            console.error(state, event)
        },
        // default handler called for all methods
        SOCKET_ONMESSAGE (state, message)  {
            console.log("Got message in store: ", message);
            state.socket.message = message
        },
        // mutations for reconnect methods
        SOCKET_RECONNECT(state, count) {
            console.info(state, count)
        },
        SOCKET_RECONNECT_ERROR(state) {
            state.socket.reconnectError = true;
        },
    },
    actions: {
        LOAD_USER_DATA: (context) => {
            const response = api.get("get_data");
            response.then(data => {
                context.commit("SET_USER", data.data);
            });
        },
        LOAD_ALL_USERS: (context) => {
            const response = api.get("all_users");
            response.then(data => {
                context.commit("SET_ALL_USERS", data.data);
            });
        },
        SEND_MESSAGE: function(context, message) {
            Vue.prototype.$socket.send(message);
        },
        RECEIVE_MESSAGE: (injectee, payload) => {
            injectee.commit("SOCKET_ONMESSAGE", payload);
        }
    },
});