<template>
    <b-sidebar width="380px" id="sidebar-1" :visible="isOpened" :title="userData.username + ' menu'"
               :no-header-close="isNoHeader"
               shadow
               backdrop>
        <b-container class="p-0">
            <b-row class="mb-3 mt-4">
                <b-col>
                <b-list-group class="w-100" v-if="userData.chats.length > 0">
                    <ChatItem :chat="chatItem" @click="emitChat(chatItem)" v-for="chatItem in userData.chats" :key="chatItem.chat_id"></ChatItem>
                </b-list-group>
                <p v-else>Чатов пока нет...</p>
                </b-col>
            </b-row>
            <b-row>
                <b-col class="m-3">
                    <b-button class="mb-1" @click="isCreateModalOpened = true" variant="primary" block>Create chat</b-button>
                    <CreateChatModal
                            @close="isCreateModalOpened = false"
                            :username="userData.username"
                            :show-modal="isCreateModalOpened">
                    </CreateChatModal>
                    <b-button class="mb-1" v-b-toggle.all-users-cont variant="info" block>Chat with new user</b-button>
                    <b-collapse id="all-users-cont">
                        <b-input v-model="findUserText" placeholder="Поиск пользователей..."></b-input>
                        <b-list-group v-if="allUsers.length > 0">
                            <b-list-group-item class="flex-row align-items-start" @click="emitContact(contactItem)" v-for="contactItem in allUsers" :key="contactItem.username">
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
        <EditChatModal
            :all-users="allUsers"
            :is-modal-opened="isEditModalOpened"
            :opened-chat="openedChat"
            @close="isEditModalOpened = false"
            @contact="emitContact($event.data)"></EditChatModal>
    </b-sidebar>
</template>

<script>
    import api from "../../api/api";
    import CreateChatModal from "./modals/CreateChatModal";
    import EditChatModal from "./modals/EditChatModal";
    import ChatItem from "./items/ChatItem";

    export default {
        name: "MobileSidebar",
        components: {ChatItem, EditChatModal, CreateChatModal},
        data() {
            return {
                findUserText: '',
                openedChat: null,
                isEditModalOpened: false,
                isCreateModalOpened: false,
                editedChatName: '',
            }
        },
        methods: {
            emitContact(user) {
                this.$emit('contact', {data: user});
            },
            signOut() {
                api.get("signout");
                this.$router.push("/login");
            },
            openChatSettingsModal(chat) {
                this.openedChat = chat;
                this.isEditModalOpened = true;
                console.log("opened chat: ", chat)
            },

        },
        computed: {
            userData() {
                return this.$store.getters.GET_USER;
            },
            allUsers() {
                let all = this.$store.getters.GET_ALL_USERS;
                if (all === undefined || all === null) {
                    return [];
                }
                return this.$store.getters.GET_ALL_USERS.filter(item => item.username !== this.userData.username);
            },
            isOpened: {
                get() {
                    return this.$store.getters.IS_SIDEBAR_SHOWN;
                },
                set() {
                }
            },
            isNoHeader() {
                return this.$store.getters.IS_INTERFACE_EMPTY;
            }
        }
    }
</script>

<style scoped>
.chat-name {
    padding: 0 0 2px 0;
    color: black;
    font-size: larger;
}
.chat-last-message {
    color: gray;
    font-size: medium;
}
    .chat-options-icon {
        position: absolute;
        top: 6px;
        right: 6px;
    }
</style>