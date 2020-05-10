<template>
    <b-list-group-item class="d-flex justify-content-between align-items-center position-relative">
        <div @click="setChat" class="d-flex justify-content-around flex-row align-content-center align-items-center">
            <b-icon-person font-scale="2" v-if="chat.usernames.length <= 2"></b-icon-person>
            <b-icon-people font-scale="2" v-else></b-icon-people>
            <div class="chat-name-and-message">
                <div class="chat-name">
                    {{ chat.name }}
                </div>
                <div class="chat-last-message" v-if="chat.messages.length > 0">
                    {{chat.messages[chat.messages.length - 1].sender}}: {{ chat.messages[chat.messages.length - 1].text }}
                </div>
            </div>
        </div>
        <b-icon-plus @click="openChatSettingsModal(chat)" font-scale="1.5" class="chat-options-icon"></b-icon-plus>
        <b-badge variant="primary" pill v-if="getNumberOfUnread > 0">
            {{ getNumberOfUnread }}
        </b-badge>
    </b-list-group-item>
</template>

<script>
    export default {
        name: "ChatItem",
        props: {
            chat: Object,
        },
        data() {
            return {
                openedChat: {}
            }
        },
        created() {
            console.log("CHAT ITEM: ", this.chat);
        },
        methods: {
            openChatSettingsModal(chat) {
                this.$store.dispatch("SET_CHAT_TO_EDIT", chat);
                this.$store.dispatch("SET_EDIT_CHAT_OPENED_STATE", true);
                console.log("opened chat: ", chat)
            },
            setChat() {
                this.$store.dispatch("CHANGE_CURRENT_CHAT", this.chat.chat_id);
                this.$store.dispatch("SET_SIDEBAR_OPENED_STATE", false);
            },
        },
        computed: {
            getNumberOfUnread() {
                let lastReadIndex = this.chat.messages.findIndex(item => item.id === this.chat.lastReadMessageId);
                if (lastReadIndex < 0) {
                    return this.chat.messages.length;
                } else {
                    return this.chat.messages.length - lastReadIndex - 1;
                }
            },
        }
    }
</script>

<style scoped>
    .chat-name-and-message {
        display: flex;
        flex-direction: column;
        text-align: left;
        justify-items: start;
        align-items: start;
        padding: 0 0 0 1em;
    }
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