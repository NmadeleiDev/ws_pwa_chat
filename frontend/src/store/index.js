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
        }
    },
    actions: {
        LOAD_USER_DATA: (context) => {
            const response = api.get("get_data");
            response.then(data => {
                context.commit("SET_USER", data);
            });
        },
        LOAD_ALL_USERS: (context) => {
            const response = api.get("all_users");
            response.then(data => {
                context.commit("SET_ALL_USERS", data);
            });
        }
    },
});