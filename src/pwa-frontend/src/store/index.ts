import Vue from 'vue'
import * as Vuex from 'vuex'
import { createStore, Module } from 'vuex-smart-module'
import { Login } from './modules/login.store'
import { UserData } from './modules/user.store'
import {Chats} from "@/store/modules/chats.store";
import {WebSocket} from "@/store/modules/socket.store";
import {CommonNotification} from "@/store/modules/notification.store";

Vue.use(Vuex)

const root = new Module({
  modules: {
    Login,
    UserData,
    Chats,
    WebSocket,
    CommonNotification,
  }
})

export const store = createStore(
    root,
    {
      strict: process.env.NODE_ENV !== 'production'
    }
)