<template>
    <b-sidebar id="sidebar-1" v-model="isOpened" title="Chats" shadow backdrop>
        <b-container>
            <b-row class="m-3 mt-2">
                <b-col cols="4">
                    <b-avatar variant="primary" text="BV"></b-avatar>
                </b-col>
                <b-col cols="6">
                    <h3>{{ userData.username }}</h3>
                </b-col>
            </b-row>
            <b-row>
                <b-col class="m-3">
                    <b-button class="mb-1" v-b-toggle.all-chats-cont variant="info" block>Активные чаты</b-button>
                    <!--                а вот список контактов потом дополнится не просто список имен, а еще и есть ли непрочитанные сообщения и все в этом роде-->
                    <!--                            <p @click="setExistingChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">{{ chatItem.name }}</p>-->
                    <b-collapse id="all-chats-cont">
                        <b-list-group v-if="userData.chats.length > 0">
                            <b-list-group-item @click="emitChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id">
                                {{ chatItem.name }}
                            </b-list-group-item>
                        </b-list-group>
                        <p v-else>Чатов пока нет...</p>
                    </b-collapse>
                </b-col>
            </b-row>
            <b-row>
                <b-col class="m-3">
                    <b-button class="mb-1" v-b-toggle.all-users-cont variant="info" block>Найти пользователя</b-button>
                    <b-collapse id="all-users-cont">
                        <b-input v-model="findUserText" placeholder="Поиск пользователей..."></b-input>
                        <b-list-group v-if="newUsers.length > 0">
                            <b-list-group-item class="flex-row align-items-start" @click="emitContact(contactItem)" v-for="contactItem in newUsers.filter(item => item.username.includes(findUserText) > 0)" :key="contactItem.username">
                                <div class="d-flex flex-row align-content-center justify-content-around">
                                    <div class="align-self-lg-start"><b-avatar></b-avatar></div>
                                    <div><h4>{{ contactItem.username }}</h4></div>
                                </div>
                            </b-list-group-item>
                        </b-list-group>
                    </b-collapse>
                </b-col>
            </b-row>
        </b-container>
        <template v-slot:footer="{ hide }">
            <div class="d-flex bg-dark text-light justify-content-around align-items-center px-3 py-2">
                <b-button variant="outline-warning" @click="signOut">Выйти</b-button>
                <b-button size="md" @click="hide">Закрыть</b-button>
            </div>
        </template>
    </b-sidebar>
</template>

<script>
    import api from "../api/api";

    export default {
        name: "MobileSidebar",
        props: {
            userData: {
                type: Object,
                required: true,
            },
            newUsers: {
                type: Array,
                default: Array,
            }
        },
        data() {
            return {
                findUserText: '',
                isOpened: false,
            }
        },
        methods: {
            emitChat(chosenChat) {
                this.$emit('chat', {data: chosenChat});
                this.isOpened = false;
            },
            emitContact(user) {
                this.$emit('contact', {data: user});
                this.isOpened = false;
            },
            signOut() {
                api.get("signout");
                this.$router.push("/login");
            }
        }
    }
</script>

<style scoped>

</style>