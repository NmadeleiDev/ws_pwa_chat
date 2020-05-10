const sidebar = {
    state: () => ({
        isShown: false,
    }),
    mutations: {
        SET_SIDEBAR_SHOWN_STATE: (state, payload) => {
            state.isShown = payload;
        },
    },
    actions: {
        SET_SIDEBAR_OPENED_STATE: (context, payload) => {
            context.commit("SET_SIDEBAR_SHOWN_STATE", payload)
        },
    },
    getters: {
        IS_SIDEBAR_SHOWN: state => {
            return state.isShown;
        }
    }
}

export default sidebar;