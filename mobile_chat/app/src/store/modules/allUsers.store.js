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
            response.then(res => {
                if (res.error === null) {
                    context.commit("SET_ALL_USERS", res.data);
                } else {
                    console.log("Error loading users: ", res.error);
                }
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
