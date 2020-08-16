import { Getters, Mutations, Actions, Module } from 'vuex-smart-module'
import api from "@/api/api";
import keysGenerator from '@/keys/keysGenerator'
import {Chat, Message, User} from '@/interfaces/main';
import {store} from '@/store'

class UserDataState {
    id: string = ''
    username: string = ''
    email: string = ''
    pool: string = ''
    token: string = ''
    userSecret: string = ''

    allUsers: Array<User> = []
}

class UserDataGetters extends Getters<UserDataState> {
    username(): string {
        return this.state.username
    }
    email(): string {
        return this.state.email
    }
    pool(): string {
        return this.state.pool
    }
    allUsers(): Array<User> {
        return this.state.allUsers
    }

    getNewToken(timestamp: string) {
        return keysGenerator.generateToken({
            username: this.state.username,
            timeStamp: timestamp,
            sessionKey: this.state.token,
            userSecret: this.state.userSecret
        })
    }
}

class UserDataMutations extends Mutations<UserDataState> {
    setToken(payload: string | null) {
        if (payload !== null)
            this.state.token = payload
        else {
            console.log("Error!!! Token not found!");
            store.dispatch('showCommonNotification', {text: 'Your authorization is not valid. Please, try to logout and login again.', type: 'error'}).catch(console.error)
        }
    }
    setUserSecret(payload: string | null) {
        if (payload !== null)
            this.state.userSecret = payload
        else {
            store.dispatch('showCommonNotification', {text: 'Your authorization is not valid. Please, try to logout and login again.', type: 'error'}).catch(console.error)
            console.log("Error!!! User Secret not found!");
        }
    }
    setUsername(payload: string | null) {
        if (payload !== null)
            this.state.username = payload
    }
    setEmail(payload: string) {
        this.state.email = payload
    }
    setPool(payload: string) {
        this.state.pool = payload
    }
    setAllUsers(payload: Array<User>) {
        this.state.allUsers.splice(0, this.state.allUsers.length)

        payload.forEach(item => {
            this.state.allUsers.push(item)
        })
    }
}

class UserDataActions extends Actions<
    UserDataState,
    UserDataGetters,
    UserDataMutations,
    UserDataActions
    > {
        loadDataLocalStorage() {
            this.commit('setUsername', localStorage.getItem('username'))
            this.commit('setToken', localStorage.getItem('sessionKey'))
            this.commit('setUserSecret', localStorage.getItem('userSecret'))
        }

        async initUserState() {
            this.dispatch('loadDataLocalStorage')

            const timeStamp: string = Date.now().toString()

            const result = await api.post('user', {data: null, auth: {username: this.state.username, token: keysGenerator.generateToken({
                username: this.state.username,
                timeStamp: timeStamp,
                sessionKey: this.state.token,
                userSecret: this.state.userSecret
            })}}, timeStamp)

            if (result.status !== true) {
                console.log("Failed to get user data!");
                store.dispatch('showCommonNotification', {text: 'Error loading your data. Please, try to login again.', type: 'error'}).catch(console.error)
                return                
            }

            this.commit('setEmail', result.data.email)
            this.commit('setPool', result.data.poolId)
            this.dispatch('loadAllUsers')
            await store.dispatch('setAllChats', result.data.chats)
        }

        async loadAllUsers() {
            this.dispatch('loadDataLocalStorage')

            const timeStamp: string = Date.now().toString()

            const result = await api.post('all_users', {data: null, auth: {username: this.state.username, token: keysGenerator.generateToken({
                        username: this.state.username,
                        timeStamp: timeStamp,
                        sessionKey: this.state.token,
                        userSecret: this.state.userSecret
                    })}}, timeStamp)

            if (result.status !== true) {
                console.log("Failed to get users!");
                store.dispatch('showCommonNotification', {text: 'Error loading users', type: 'error'}).catch(console.error)
                return
            }

            this.commit('setAllUsers', result.data)
        }

        async logInToPool(payload: {poolId: string, poolPassword: string}) {
            const timeStamp: string = Date.now().toString()

            const result = await api.post('pool', {data: payload, auth: {username: this.state.username, token: keysGenerator.generateToken({
                        username: this.state.username,
                        timeStamp: timeStamp,
                        sessionKey: this.state.token,
                        userSecret: this.state.userSecret
                    })}}, timeStamp)

            if (result.status !== true) {
                console.log("Failed to log in to pool");
                store.dispatch('showCommonNotification', {text: 'Error logging to pool.', type: 'error'}).catch(console.error)
                return
            }

            this.commit('setPool', payload.poolId)
        }

    async createPool(payload: {poolId: string, poolPassword: string}) {
        const timeStamp: string = Date.now().toString()

        const result = await api.put('pool', {data: payload, auth: {username: this.state.username, token: keysGenerator.generateToken({
                    username: this.state.username,
                    timeStamp: timeStamp,
                    sessionKey: this.state.token,
                    userSecret: this.state.userSecret
                })}}, timeStamp)

        if (result.status !== true) {
            console.log("Failed to log in to pool");
            store.dispatch('showCommonNotification', {text: 'Error logging to pool.', type: 'error'}).catch(console.error)
            return
        }

        this.commit('setPool', payload.poolId)
    }
}

export const UserData = new Module({
    namespaced: false,
    state: UserDataState,
    getters: UserDataGetters,
    mutations: UserDataMutations,
    actions: UserDataActions
})