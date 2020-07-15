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
        if (this.state.client !== undefined)
            this.state.client.send(JSON.stringify(payload))
        else {
            console.log("Error! Client undefined!")
        }
    }

    initWebSocket() {
        if (this.state.isConnected) {
            return
        }
        const timeStamp: string = Date.now().toString()
        const username = store.getters.username()

        const token = store.getters.getNewToken(timeStamp)
        const client = new W3CWebSocket('ws://localhost:8080/api/v1/connect?user=' + username + '&token=' + token + '&time=' + timeStamp, 'chat')
        // const client = new W3CWebSocket('wss://enchat.ga/api/v1/connect?user=' + username + '&token=' + token + '&time=' + timeStamp, 'chat')

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
        let data = <ServerToClientMessage>JSON.parse(message)
        if (data.type === MessageType) {
            store.commit('addMessageToItsChat', <WebSocketChatMessage>data.data)
        } else {
            const timeStamp: string = Date.now().toString()

            api.post('messages', {data: <Chat>data.data, auth: {username: store.getters.username(), token: store.getters.getNewToken(timeStamp)}}, timeStamp)
                .then(result => {
                    if (result.status !== true) {
                        console.log("Error loading messages")
                        return
                    }
                    let chat = <Chat>data.data
                    const newChat: Chat =  {
                        id: chat.id,
                        admin: chat.admin,
                        name: chat.name,
                        usernames: chat.usernames,
                        messages: new Array<Message>()
                    }
                    if (Array.isArray(result.data)) {
                        newChat.messages = <Array<Message>>result.data
                    }
                    store.commit('addChat', <Chat>data.data)
                })
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