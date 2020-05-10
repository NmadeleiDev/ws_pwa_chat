const editChatModal = {
    state: () => ({
        isShown: false,
        openedChat: null,
    }),
    mutations: {
        SET_SHOWN_STATE: (state, payload) => {
            state.isShown = payload;
        },
        SET_CHAT_TO_EDIT: (state, payload) => {
            state.openedChat = payload;
        }
    },
    actions: {
        SET_EDIT_CHAT_OPENED_STATE: (context, payload) => {
            context.commit("SET_SHOWN_STATE", payload)
        },
        SET_CHAT_TO_EDIT: (context, payload) => {
            context.commit("SET_CHAT_TO_EDIT", payload)
        }
    },
    getters: {
        CHAT_TO_EDIT: state => {
            return state.openedChat;
        },
        IS_EDIT_CHAT_SHOWN: state => {
            return state.isShown;
        }
    }
}

export default editChatModal;