import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from "../pages/Login";
import Chat from "../pages/Chat";

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        redirect: '/home',
    },
    {
        path: '/login',
        redirect: '/home',
    },
    {
        path: '/home',
        name: 'home_page',
        meta: {
            title: 'Home page',
        },
        component: Login,
    },
    {
        path: '/chat',
        name: 'chat_page',
        meta: {
            title: 'Chat page',
        },
        component: Chat,
    },
];

const router = new VueRouter({
    mode: 'history',
    base: '/',
    routes,
});

export default router;
