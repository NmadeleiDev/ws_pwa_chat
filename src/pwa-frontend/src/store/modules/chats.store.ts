import { Getters, Mutations, Actions, Module } from 'vuex-smart-module'
import api from "@/api/api";
import {Chat, Message, User} from "@/interfaces/main";
import {store} from '@/store';
// @ts-ignore
import md5 from 'md5';

// State
class ChatsState {
    isNew: boolean = false
    chats: Array<Chat> = []
    currentChat: Chat = {id: '', admin: '', name: '', usernames: [], messages: [], storePeriod: 24}

    generateMessageId = function () {
        return md5(Date.now() + Math.random())
    }
}

class ChatsGetters extends Getters<ChatsState> {
    getCurrentChat(): Chat | undefined {
        if (this.state.isNew) {
            return this.state.currentChat
        }
        return this.state.chats.find((chat: Chat) => chat.id === this.state.currentChat.id)
    }

    getAllChats(): Array<Chat> {
        return this.state.chats
    }

    isNew() {
        return this.state.isNew
    }

    getChatById(chatId: string): Chat {
        const data = this.state.chats.find(item => item.id === chatId)
        if (data === undefined) {
            return <Chat>{id: "", name: "Chat", messages: [], usernames: [store.getters.username()], admin: store.getters.username(), storePeriod: 24}
        } else {
            return data
        }
    }
}

class ChatsMutations extends Mutations<ChatsState> {
    sortChats() {
        this.state.chats.sort((a: Chat, b: Chat) => {
            if (a.messages.length > 0 && b.messages.length > 0) {
                return b.messages[b.messages.length - 1].date - a.messages[a.messages.length - 1].date
            }
            return 1
        })
    }

    setChats(payload: Array<Chat>) {
        this.state.chats.splice(0, this.state.chats.length)
        payload.forEach(chat => {
            this.state.chats.push(chat)
        })
    }

    addMessageToItsChat(payload: Message) {
        this.state.chats.forEach((chat: Chat) => {
            if (chat.id === payload.chatId) {

                let index = chat.messages.findIndex(item => item.id === payload.id)
                if (index < 0) {
                    chat.messages.push(payload);
                } else {
                    chat.messages[index].text = payload.text
                    chat.messages[index].state = payload.state
                }
            }
        })
    }

    addMessageToCurrentChat(payload: Message) {
        try {
            this.state.currentChat.messages.push(payload)
        } catch (e) {
            console.log("Failed to push to current chat: ", e)
        }
    }

    addChat(payload: Chat) {
        console.log("Adding chat: ", payload)
        if (payload.messages === undefined) {
            payload.messages = new Array<Message>()
        }
        if (this.state.chats.find(item => item.id === payload.id) === undefined) {
            this.state.chats.push(payload)
        }
    }

    setCurrentChat(payload: Chat) {
        this.state.currentChat = payload
    }

    setCurrentChatId(id: string) {
        this.state.currentChat.id = id
    }

    setChatData(payload: Chat) {
        const chat = this.state.chats.find(item => item.id === payload.id)
        if (chat !== undefined) {
            chat.name = payload.name
            chat.usernames = payload.usernames
            chat.admin = payload.admin
        } else {
            this.state.chats.push(payload)
        }
    }

    setCurrentChatName(name: string) {
        if (this.state.isNew) {
            this.state.currentChat.name = name
            return
        }
        const chat = this.state.chats.find(item => item.id === this.state.currentChat.id)
        if (chat !== undefined) {
            chat.name = name
        }
    }

    setNewValue(payload: boolean) {
        this.state.isNew = payload;
    }

    setMessageState(payload: {message: Message, state: number}) {
        payload.message.state = payload.state
    }

    setCurrentChatStorePeriod(payload: number) {
        this.state.currentChat.storePeriod = payload;
    }
}

class ChatsActions extends Actions<
    ChatsState,
    ChatsGetters,
    ChatsMutations,
    ChatsActions
    > {

    async sendMessageInChat(payload: string) {
        const message: Message = {
            id: this.state.generateMessageId(),
            sender: store.getters.username(),
            chatId: this.state.currentChat.id,
            date: Date.now(),
            state: 1,
            text: payload,
            attachedFilePath: ''
        }

        if (this.state.isNew) {
            this.commit('addMessageToCurrentChat', message)

            let newChat = this.getters.getCurrentChat()
            if (newChat !== undefined) {
                try {
                    newChat = await this.dispatch('createChat', newChat as Chat)
                } catch (e) {
                    console.log('Error creating chat: ', e)
                    store.dispatch('showCommonNotification', {text: 'Error creating chat.', type: 'error'}).catch(console.error)
                    return
                }
                this.commit('setCurrentChat', newChat as Chat)
                this.commit('setNewValue', false);
            } else {
                console.log('Current chat not found, wtf?')
                store.dispatch('showCommonNotification', {text: 'Error creating chat.', type: 'error'}).catch(console.error)
            }
        } else {
            try {
                await store.dispatch('sendSocketMessage', {type: 1, message: message})
            } catch (e) {
                console.log('Error sending message: ', e)
                store.dispatch('showCommonNotification', {text: 'Error sending message.', type: 'error'}).catch(console.error)
            }
        }

        this.commit('addMessageToItsChat', message)
    }

    async createChat(payload: Chat): Promise<Chat> {
        const timeStamp: string = Date.now().toString()

        let result = await api.post("chat", {data: payload, auth: {username: store.getters.username(), token: store.getters.getNewToken(timeStamp)}}, timeStamp)
        if (result.status !== true) {
            console.log("Failed to send init message! Failed to create chat!")
            store.dispatch('showCommonNotification', {text: 'Chat is not created, please, try again.', type: 'error'}).catch(console.error)
            return payload
        }
        const newChat: Chat =  {
            id: result.data.id,
            admin: result.data.admin,
            name: result.data.name,
            usernames: result.data.usernames,
            messages: payload.messages as Array<Message>,
            storePeriod: 24,
        }
        this.commit('addChat', newChat)
        return newChat;
    }

    async setCurrentChat(payload: {data: Chat | User, isNew: boolean}) {
        if (payload.isNew) {
            let user: User = <User>payload.data
            let chat: Chat = {
                id: '',
                name: store.getters.username() + ' and ' + user.username,
                usernames: [store.getters.username(), user.username],
                admin: store.getters.username(),
                messages: [],
                storePeriod: 24
            }
            this.commit('setNewValue', true);
            this.commit('setCurrentChat', chat);
        } else {
            const chat: Chat = <Chat>payload.data
            this.commit('setCurrentChat', chat);
            this.commit('setNewValue', false);
            chat.messages.forEach(message => {
                if (message.sender !== store.getters.username() && message.state !== 3) {
                    this.commit('setMessageState', {message: message, state: 3})
                    store.dispatch('sendSocketMessage', {type: 2, message: message})
                }
            })
        }
    }

    async setAllChats(payload: Array<Chat>) {
        await Promise.all(payload.map(async (chat: Chat) => {
            const timeStamp: string = Date.now().toString()

            const result = await api.post('messages', {data: chat, auth: {username: store.getters.username(), token: store.getters.getNewToken(timeStamp)}}, timeStamp)
            if (result.status !== true) {
                console.log("Error loading messages")
                return
            }
            chat.messages = !result.data ? [] : <Array<Message>>result.data
            chat.messages.forEach(item => item.chatId = chat.id)
        }))
        this.commit('setChats', payload)
        this.commit('sortChats')
    }

    saveCurrentChatName(name: string) {
        this.commit('setCurrentChatName', name)
        const currentChat = this.getters.getCurrentChat()
        if (currentChat === undefined)
            return
        const timeStamp: string = Date.now().toString()
        api.post("name", {data:
        {
            id: currentChat.id,
            name: currentChat.name,
        },
        auth: {
            username: store.getters.username(), token: store.getters.getNewToken(timeStamp)
        }}, timeStamp).catch(console.warn)
    }

    async addUserToCurrentChat(username: string) {
        const timeStamp: string = Date.now().toString()

        const result = await api.post("add", {auth: {
                username: store.getters.username(),
                token: store.getters.getNewToken(timeStamp)
            }, data: {
                user: {username: username},
                chat: this.getters.getCurrentChat()
            }}, timeStamp)
        if (result.status !== true) {
            console.log("Error add to chat: ", result.data)
            return false
        }
        return true
    }

    setCurrentChatId(id: string) {
        this.commit('setCurrentChatId', id)
    }

    async setCurrentChatStorePeriod(period: number) {
        const timeStamp: string = Date.now().toString()
        const currentChat = this.getters.getCurrentChat()
        if (currentChat === undefined) {
            console.log("Modifying undef chat")
            return
        }

        const result = await api.post("period", {auth: {
                username: store.getters.username(),
                token: store.getters.getNewToken(timeStamp)
            }, data: {
                id: currentChat.id,
                name: currentChat.name,
                storePeriod: period
            }}, timeStamp)
        if (result.status !== true) {
            console.log("Error save period: ", result.data)
            return false
        }
        this.commit('setCurrentChatStorePeriod', period);
    }

}

export const Chats = new Module({
    namespaced: false,
    state: ChatsState,
    getters: ChatsGetters,
    mutations: ChatsMutations,
    actions: ChatsActions,
})