export interface User {
    username: string
    isOnline: boolean
}

export interface Chat {
    id: string
    name: string
    usernames: Array<string>
    admin: string
    messages: Array<Message>
    storePeriod: number
}

export interface Message {
    id: string
    sender: string
    chatId: string
    date: number
    state: number
    text: string
    attachedFileIds: Array<string>
}

export interface WebSocketChatMessage {
    type: number
    message: Message
}

export interface ServerToClientMessage {
    type: number
    data: object
    error: string
}