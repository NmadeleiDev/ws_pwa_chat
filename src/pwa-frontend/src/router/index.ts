import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Login from '../views/Login/Login.vue'

Vue.use(VueRouter)

  const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import(/* webpackChunkName: "about" */ '../views/ChatsList/ChatsList.vue')
  },
  {
    path: '/users',
    name: 'All users',
    component: () => import(/* webpackChunkName: "about" */ '../views/UsersList/UsersList.vue')
  },
  {
    path: '/chat/:id',
    name: 'Chat here',
    component: () => import(/* webpackChunkName: "about" */ '../views/Chat/Chat.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
