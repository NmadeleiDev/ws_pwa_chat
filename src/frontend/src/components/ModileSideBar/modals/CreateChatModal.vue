<template>
    <b-modal id="create-chat-modal" v-model="isModalOpened">
        <b-alert v-model="showAlert" :variant="alertVariant" dismissible>
            {{ alertText }}
        </b-alert>
        <div v-if="allUsers.length > 2" class="d-flex flex-column justify-content-between align-items-center">
            <h4 class="m-2">Create chat</h4>
            <b-input class="m-2" placeholder="Enter chat name..." v-model="chatName"></b-input>
            <div class="users-container mb-2 mt-2 w-100" >
                <b-input class="mb-2 w-100" placeholder="Search user..." v-model="findUserName"></b-input>
                <b-list-group>
                    <b-list-group-item :active="chosenUsers.includes(user)" @click="toggleUser(user)"
                                       v-for="user in allUsers.filter(item => (item.username.includes(findUserName) || chosenUsers.includes(item)))"
                                       :key="user.username">
                        {{ user.username }}
                    </b-list-group-item>
                </b-list-group>
            </div>
            <b-button variant="success" @click="createChat" class="m-2">Create</b-button>
        </div>
        <b-alert show v-else>There is not enough users to create group chat</b-alert>
    </b-modal>
</template>

<script>
    export default {
        name: "CreateChatModal",
        props: {
            username: String,
            showModal: Boolean,
        },
        data() {
            return {
                chatName: '',
                findUserName: '',
                chosenUsers: [],

                showAlert: false,
                alertVariant: 'success',
                alertText: '',
            }
        },
        methods: {
            createChat() {
                if (this.chosenUsers.length < 2) {
                    this.showAlert = true;
                    this.alertVariant = 'danger';
                    this.alertText = 'Chat must contain more than one person';
                    return;
                }
                let users = this.chosenUsers.map(item => item.username).concat([this.username]);
                const chat = {
                    name: 'group (' + users.join(', ') + ')',
                    admin: this.username,
                    usernames: users,
                }
                this.$store.dispatch('CREATE_CHAT', chat).then(data => {
                    if (data === null) {
                        console.log('error creating chat');
                        this.showAlert = true;
                        this.alertVariant = 'warning';
                        this.alertText = 'Error creating chat';
                    } else {
                        console.log('created chat: ', data);
                        this.showAlert = true;
                        this.alertVariant = 'success';
                        this.alertText = 'Chat created!';
                    }
                });
                this.isModalOpened = false;
            },
            toggleUser(user) {
                let index = this.chosenUsers.indexOf(user)
                if (index >= 0) {
                    this.chosenUsers.splice(index, 1);
                } else {
                    this.chosenUsers.push(user);
                }
            }
        },
        computed: {
            isModalOpened: {
                get() {
                    return this.showModal;
                },
                set() {
                    this.$emit("close");
                }
            },
            allUsers() {
                return this.$store.getters.GET_ALL_USERS;
            }
        }
    }
</script>

<style scoped>
.users-container {
    max-height: 80%;
    overflow-x: hidden;
    overflow-x: auto;
}
</style>