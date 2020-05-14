import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from "../pages/login/Login";
import Chat from "../pages/chat/Chat";

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        redirect: '/login',
    },
    {
        path: '/login',
        name: 'login_page',
        meta: {
            title: 'login page',
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
