<template>
    <v-alert dismissible dense :value="show" :type="type" class="ml-3 mr-3">
        {{text}}
    </v-alert>
</template>

<script lang="ts">
    import Vue from 'vue'


    export default Vue.extend({
        name: "CommonNotification",
        computed: {
            show: {
                get(): boolean {
                    return this.$store.getters.isShown()
                },
                set(val: boolean) {
                    this.$store.dispatch('setCommonNotificationValue', val)
                }
            },
            text() {
                return this.$store.getters.notificationText()
            },
            type() {
                return this.$store.getters.notificationType()
            }
        }
    })
</script>

<style scoped>

</style>