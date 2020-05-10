import Vue from 'vue';
import Vuex from 'vuex';
import user from "./modules/user.store";
import operational from "./modules/operational.store";
import allUsers from "./modules/allUsers.store";
import editChatModal from "./modules/components/editChat.store";
import sidebar from "./modules/components/sidebar.store";

Vue.use(Vuex);

export default new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    modules: {
        user: user,
        operational: operational,
        allUsers: allUsers,
        editChatModal: editChatModal,
        sidebar: sidebar,
    },
    state: {
        socket: {
            isConnected: false,
            message: '',
            reconnectError: false,
        },
    },
    getters: {

    },
    mutations: {
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
        // mutations for reconnect methods
        SOCKET_RECONNECT(state, count) {
            console.info(state, count)
        },
        SOCKET_RECONNECT_ERROR(state) {
            state.socket.reconnectError = true;
        },
    },
    actions: {

    },
});