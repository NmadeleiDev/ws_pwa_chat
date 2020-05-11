<template>
    <b-modal v-if="openedChat !== null" id="chat-settings-modal" :title="openedChat.name"
             v-model="isOpened"
             size="lg"
             scrollable
             centered
             button-size="sm">
        <div class="d-flex flex-column justify-content-around align-items-center">
            <div class="m-2 d-flex flex-column justify-content-between align-items-center">
                <b-button size="sm" class="mb-1" v-b-toggle.chat-name-edit variant="secondary">Edit name</b-button>
                <b-collapse id="chat-name-edit">
                    <b-input-group>
                        <b-input placeholder="Enter new name..." v-model="editedChatName"></b-input>
                        <b-input-group-append>
                            <b-button @click="saveChatName" variant="primary">Save</b-button>
                        </b-input-group-append>
                    </b-input-group>
                </b-collapse>
            </div>
            <b-list-group class="w-100">
                <b-list-group-item v-for="user in openedChat.usernames" :key="user"
                                   class="d-flex justify-content-between align-items-center">
                    <h5>{{ user }}</h5>
                    <b-badge v-if="user === openedChat.admin" variant="warning" pill>Admin</b-badge>
                </b-list-group-item>
            </b-list-group>
            <div v-if="allUsers.length > 0">
                <b-button size="sm" class="m-3" v-b-toggle.all-users-collapse variant="primary"><b-icon-plus></b-icon-plus> Add users</b-button>
                <b-collapse id="all-users-collapse">
                    <b-input v-model="findUserText" placeholder="Поиск пользователей..."></b-input>
                    <b-list-group v-if="allUsers.length > 0 && openedChat !== null">
                        <b-list-group-item class="d-flex justify-content-between align-items-center" @click="emitContact(contactItem)" v-for="contactItem in allUsers.filter(item => !openedChat.usernames.includes(item.username) && item.username.indexOf(findUserText) >= 0)" :key="contactItem.username">
                            <div>
                                {{ contactItem.username }}
                            </div>
                            <b-button @click="addUserToChat(contactItem)" variant="info"><b-icon-plus></b-icon-plus> Add</b-button>
                        </b-list-group-item>
                    </b-list-group>
                </b-collapse>
            </div>
            <div>
                <b-button @click="leaveChat" variant="outline-danger">Leave chat</b-button>
            </div>
        </div>
    </b-modal>
</template>

<script>
    export default {
        name: "EditChatModal",
        data() {
            return {
                editedChatName: '',
                findUserText: '',
            }
        },
        methods: {
            saveChatName() {
                let chatWithNewName = this.openedChat;
                chatWithNewName.name = this.editedChatName;
                this.$store.dispatch("SAVE_CHAT_NAME", chatWithNewName);
            },
            addUserToChat(user) {
                if (this.openedChat === null)
                    return;
                const packet = {
                    chat: this.openedChat,
                    user: user
                }
                this.$store.dispatch("ADD_USER_TO_CHAT", packet);
            },
            emitContact(contact) {
                this.$emit('contact', {data: contact});
            },
            leaveChat() {
                this.$store.dispatch("LEAVE_CHAT", this.openedChat);
            }
        },
        computed: {
            isOpened: {
                get() {
                    return this.$store.getters.IS_EDIT_CHAT_SHOWN;
                },
                set(value) {
                    this.$store.dispatch("SET_EDIT_CHAT_OPENED_STATE", value);
                }
            },
            allUsers() {
                return this.$store.getters.GET_ALL_USERS;
            },
            openedChat() {
                return this.$store.getters.CHAT_TO_EDIT;
            }
        }
    }
</script>

<style scoped>

</style>