import { Getters, Mutations, Actions, Module } from 'vuex-smart-module'
import api from "@/api/api";
import keysGenerator from '@/keys/keysGenerator'
import {store} from "@/store";

// State
class LoginState {
    isLogged = false
}

// Getters
// Extend 'Getters' class with 'LoginState' type
class LoginGetters extends Getters<LoginState> {
    // You can declare both getter properties or methods
    get isLogged() {
        // Getters instance has 'state' property
        return this.state.isLogged;
    }
}

// Mutations
// Extend 'Mutations' class with 'LoginState' type
class LoginMutations extends Mutations<LoginState> {
    setLoginState(payload: boolean) {
        // Mutations instance has 'state' property.
        // You update 'this.state' by mutating it.
        this.state.isLogged = payload
    }
}

// Actions
// Extend 'Actions' class with other module asset types
// Note that you need to specify self action type (LoginActions) as a type parameter explicitly
class LoginActions extends Actions<
    LoginState,
    LoginGetters,
    LoginMutations,
    LoginActions
    > {

    async signUp(payload: { login: string; password: string }) {
        payload.password = keysGenerator.getSha224(payload.password)
        
        let result = await api.post('signup', {data: {username: payload.login, password: payload.password}, auth: null}, "")
        if (result.status === true) {
            localStorage.setItem('username', payload.login);
            localStorage.setItem('sessionKey', result.data.token);
            localStorage.setItem('userSecret', keysGenerator.generateUserSecret({username: payload.login, password: payload.password}));

            this.commit('setLoginState', true);

            console.log("Signed up successfully: ", localStorage.getItem('sessionKey'), localStorage.getItem('userSecret'));
            return true;
        }
        store.dispatch('showCommonNotification', {text: 'Sign up failed. Please, try again.', type: 'error'}).catch(console.error)
        return false
    }

    async signIn(payload: { login: string; password: string }) {
        payload.password = keysGenerator.getSha224(payload.password)

        let result = await api.post('signin', {data: {username: payload.login, password: payload.password}, auth: null}, "")
        if (result.status === true) {
            localStorage.setItem('username', payload.login);
            localStorage.setItem('sessionKey', result.data.token);
            localStorage.setItem('userSecret', keysGenerator.generateUserSecret({username: payload.login, password: payload.password}));

            this.commit('setLoginState', true);

            console.log("Signed in successfully: ", localStorage.getItem('sessionKey'), localStorage.getItem('userSecret'));
            return true;
        }
        store.dispatch('showCommonNotification', {text: 'Sign in failed. Please, try again.', type: 'error'}).catch(console.error)
        return false
    }

    async findLocalKeys() {

        let sessionKey = localStorage.getItem('sessionKey');
        let userSecret = localStorage.getItem('userSecret');
        if (sessionKey && userSecret && userSecret !== '' && sessionKey !== '') {
            console.log("Keys found in storage")
            this.commit('setLoginState', true);
            return true;
        }
        console.log("Keys not found in storage")
        return false
    }
}

export const Login = new Module({
    namespaced: false,
    state: LoginState,
    getters: LoginGetters,
    mutations: LoginMutations,
    actions: LoginActions,
})