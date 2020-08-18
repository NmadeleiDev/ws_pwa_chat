<template>
    <v-app>
        <DefaultAppBar name="Find users"></DefaultAppBar>
        <v-main>
            <v-row>
                <v-col>
                    <CommonNotification></CommonNotification>
                    <v-list shaped v-if="usersToShow.length > 1">
                        <v-subheader>Users</v-subheader>
                        <v-text-field class="ml-2 mr-2" placeholder="Find user" v-model="filter"></v-text-field>
                        <v-list-item-group color="primary">
                            <v-list-item
                                    v-for="(user, i) in usersToShow"
                                    :key="i"
                                    @click="setChat(user)"
                            >
                                <v-list-item-icon>
                                    <v-icon v-if="user.isOnline">person</v-icon>
                                    <v-icon v-else>person_outline</v-icon>
                                </v-list-item-icon>
                                <v-list-item-content>
                                    <v-list-item-title v-text="user.username"></v-list-item-title>
                                </v-list-item-content>
                            </v-list-item>
                        </v-list-item-group>
                    </v-list>

                    <v-sheet v-else class="mt-8 d-flex flex-column justify-space-around align-center">
                        <h3>Looks like you are the only user yet!</h3>
                        <v-subheader>Try refreshing page :)</v-subheader>
                    </v-sheet>
                </v-col>
            </v-row>
        </v-main>
    </v-app>
</template>

<script lang="ts">
    import Vue from "vue";
    import CommonNotification from "@/components/CommonNotification.vue";
    import {Chat, User} from "@/interfaces/main";
    import DefaultAppBar from "@/components/DefaultAppBar.vue";

    export default Vue.extend({
        name: "UsersList",
        components: {
            DefaultAppBar,
            CommonNotification
        },
        data: () => {
            return {
                filter: ''
            }
        },
        created() {
            if (this.users.length === 0) {
                this.$store.dispatch('loadAllUsers');
            }
        },
        methods: {
            setChat(user: User) {
                this.$store.dispatch('setCurrentChat', {data: user, isNew: true})
                this.$router.push('/chat/' + user.username)
            }
        },
        computed: {
            users(): Array<User> {
                return this.$store.getters.allUsers()
            },
            username(): string {
                return this.$store.getters.username()
            },
            existingPersonalChats(): Array<string> {
                return this.$store.getters.getAllChats().reduce(
                    (acc: Array<string>, curr: Chat) => {
                        if (curr.usernames.length === 2) {
                            return acc.concat(curr.usernames)
                        }
                        return acc
                    }, [] as Array<string>
                )
            },
            usersToShow() {
                return this.$store.getters.allUsers().filter((item: User) => item.username.includes(this.filter) && (!this.existingPersonalChats.includes(item.username)))
            }
        }
    })
</script>

<style scoped>

</style>