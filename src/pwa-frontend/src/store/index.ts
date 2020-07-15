import Vue from 'vue'
import * as Vuex from 'vuex'
import { createStore, Module } from 'vuex-smart-module'
import { Login } from './modules/login.store'
import { UserData } from './modules/user.store'
import {Chats} from "@/store/modules/chats.store";
import {WebSocket} from "@/store/modules/socket.store";

Vue.use(Vuex)

const root = new Module({
  modules: {
    Login,
    UserData,
    Chats,
  WebSocket
  }
})

export const store = createStore(
    root,
    {
      strict: process.env.NODE_ENV !== 'production'
    }
)