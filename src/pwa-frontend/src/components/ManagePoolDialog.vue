<template>
    <v-dialog
            v-model="dialog"
            width="500"
    >
        <template v-slot:activator="{ on, attrs }">
            <v-btn
                    color="black"
                    class="text--white"
                    text
                    dark
                    v-bind="attrs"
                    v-on="on"
            >
                {{pool === '' ? 'Common pool' : pool}}
            </v-btn>
        </template>

        <v-card>
            <v-card-title>{{'You are in ' + (pool === '' ? 'common pool' : pool)}}</v-card-title>
            <v-card-subtitle>Here you can log to another pool</v-card-subtitle>
            <v-card-text>
                <v-text-field placeholder="Pool name"
                v-model="poolId"></v-text-field>

                <v-text-field placeholder="Pool password"
                type="password"
                v-model="poolPassword"></v-text-field>
            </v-card-text>
            <v-card-actions>
                <v-btn @click="isCreating ? createPool() : logIn()">{{isCreating ? 'Create pool' : 'Log in to poll'}}</v-btn>
                <v-btn text @click="isCreating = !isCreating">{{isCreating ? 'Log in to poll' : 'Create pool'}}</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend({
        name: "ManagePoolDialog",
        data: () => {
            return {
                isCreating: false,
                poolId: '',
                poolPassword: '',

                dialog: false,
            }
        },
        methods: {
            logIn() {
                this.$store.dispatch('logInToPool', {poolId: this.poolId, poolPassword: this.poolPassword}).then(() => {
                    this.dialog = false;
                }).catch(console.error)
            },
            createPool() {
                this.$store.dispatch('createPool', {poolId: this.poolId, poolPassword: this.poolPassword}).then(() => {
                    this.dialog = false;
                }).catch(console.error)
            }
        },
        computed: {
            pool(): string {
                return this.$store.getters.pool();
            },
        }
    })
</script>

<style scoped>

</style>