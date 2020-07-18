import { Getters, Mutations, Actions, Module } from 'vuex-smart-module'
import api from "@/api/api";
import {Chat, Message, ServerToClientMessage, WebSocketChatMessage} from "@/interfaces/main";
import keysGenerator from "@/keys/keysGenerator";
import {store} from "@/store";
const W3CWebSocket = require('websocket').w3cwebsocket;

const MessageType = 1
const ChatType = 2

// State
class WebSocketState {
    client: WebSocket | undefined
    isConnected: boolean = false
}

class WebSocketGetters extends Getters<WebSocketState> {
    isConnected(): boolean {
        return this.state.isConnected
    }
}

class WebSocketMutations extends Mutations<WebSocketState> {

    setClient(client: WebSocket) {
        this.state.client = client
    }

    setConnected(payload: boolean) {
        this.state.isConnected = payload
    }
}

class WebSocketActions extends Actions<
    WebSocketState,
    WebSocketGetters,
    WebSocketMutations,
    WebSocketActions
    > {

    sendSocketMessage(payload: WebSocketChatMessage) {
        if (this.state.client !== undefined && this.state.isConnected)
            this.state.client.send(JSON.stringify(payload))
        else {
            console.log("Error! IsConnected: ", this.state.isConnected, " Client: ", this.state.client)
        }
    }

    initWebSocket() {
        if (this.state.isConnected) {
            return
        }
        const timeStamp: string = Date.now().toString()
        const username = store.getters.username()

        const token = store.getters.getNewToken(timeStamp)
        // const client = new W3CWebSocket('ws://localhost:8080/ws/connect?user=' + username + '&token=' + token + '&time=' + timeStamp, 'chat')
        const client = new W3CWebSocket('wss://enchat.ga/ws/connect?user=' + username + '&token=' + token + '&time=' + timeStamp, 'chat')

        client.onopen = () => {
            this.dispatch('onOpen')
        }

        client.onclose = () => {
            this.dispatch('onClose')
        }

        client.onmessage = (message: any) => {
            this.dispatch('onMessage', message.data)
        }

        client.onerror = () => {
            console.log('Connection Error');
        };

        this.commit('setClient', client)
    }

    onOpen() {
        console.log("Opened ws")
        this.commit('setConnected', true)
    }

    onClose() {
        console.log("Closed ws")
        this.commit('setConnected', false)
    }

    onMessage(message: any) {
        console.log("Got message in ws: ", message)
        let data = JSON.parse(message) as ServerToClientMessage
        if (data.type === MessageType) {
            let message: Message = data.data as Message
            store.commit('addMessageToItsChat', message)
            if (message.sender !== store.getters.username()) {
                message.state = 3
                this.dispatch('sendSocketMessage', {type: 2, message: message})
                this.dispatch('notifyUser', data)
            }
        } else if (data.type === ChatType) {
            const timeStamp: string = Date.now().toString()
            const chat: Chat = <Chat>data.data

            if (store.getters.getChatById(chat.id).id !== "") {
                store.commit('setChatData', chat)
                console.log("Found and updated chat " + chat.name)
                return
            }

            if (chat.admin !== store.getters.username()) {
                this.dispatch('notifyUser', data)
                api.post('messages', {data: chat, auth: {username: store.getters.username(), token: store.getters.getNewToken(timeStamp)}}, timeStamp)
                    .then(result => {
                        if (result.status !== true) {
                            console.log("Error loading messages")
                            return
                        }
                        if (Array.isArray(result.data)) {
                            chat.messages = result.data as Array<Message>
                        }
                        store.commit('addChat', chat)
                    })
            }
        }
        store.commit('sortChats')
    }

    notifyUser(payload: ServerToClientMessage) {
        let title = ''
        let body = ''
        let actions: { action: string; title: string; }[] = []
        if (payload.type === MessageType) {
            let message: Message = payload.data as Message
            let chat: Chat = store.getters.getChatById(message.chatId)
            title = 'New message in ' + chat.name
            body = message.sender + ': ' + message.text
            actions = [
                {
                    action: 'open',
                    title: 'Open ' + chat.name
                }
            ]
        } else if (payload.type === ChatType) {
            let chat: Chat = payload.data as Chat
            title = "You've been added to chat " + chat.name
            body = 'Chat members: ' + chat.usernames.join(', ')
            actions = [
                {
                    action: 'open',
                    title: 'Open ' + chat.name
                }
            ]
        }

        if (Notification.permission === 'granted') {
            navigator.serviceWorker.getRegistration()
                .then((reg) => {
                        if(reg == undefined){
                            console.log("only works online")
                            return
                        }
                        let options = {
                            body: body,
                            // icon: './static/img/notification-flat.png',
                            vibrate: [100, 50, 100],
                            data: {
                                dateOfArrival: Date.now(),
                                primaryKey: 1
                            },
                            actions: actions
                        }
                        reg.showNotification(title, options).catch(console.warn)
                    }
                )
        } else {
            Notification.requestPermission((status) =>  {
                console.log('Notification permission status:', status);
            }).catch(console.warn);
        }
    }

}

export const WebSocket = new Module({
    namespaced: false,
    state: WebSocketState,
    getters: WebSocketGetters,
    mutations: WebSocketMutations,
    actions: WebSocketActions,
})
