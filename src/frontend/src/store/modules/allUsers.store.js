import api from "../../api/api";

const allUsers = {
    state: () => ({
        allUsers: [],
    }),
    mutations: {
        SET_ALL_USERS: (state, payload) => {
            payload.forEach(item => {
                state.allUsers.push(item);
            })
        },
    },
    actions: {
        LOAD_ALL_USERS: (context) => {
            const response = api.get("all_users");
            response.then(data => {
                context.commit("SET_ALL_USERS", data.data);
            });
        },
    },
    getters: {
        GET_ALL_USERS: state => {
            return state.allUsers;
        },
    }
}

export default allUsers;