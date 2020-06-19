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
        LOAD_ALL_USERS: async (context) => {
            try {
                const response = await api.get("all_users");
                context.commit("SET_ALL_USERS", response.data);
                return true;
            } catch (e) {
                console.log(e);
                return false;
            }
        },
    },
    getters: {
        GET_ALL_USERS: state => {
            return state.allUsers;
        },
    }
}

export default allUsers;