<template>
    <b-modal title="You are not logged in to any messaging pool" :visible="isShown" no-close-on-esc no-close-on-backdrop hide-header-close cancel-disabled ok-disabled>
        <b-tabs>
            <b-alert :type="alert.type" :show="alert.isShown" dismissible>{{ alert.text }}</b-alert>
            <b-tab title="Join">
                <div class="d-flex flex-column align-items-center justify-content-center">
                    <h4 class="m-3">Enter your pool:</h4>
                    <b-input class="m-3" placeholder="Enter your pool..." v-model="poolId"></b-input>
                    <b-input class="m-3" placeholder="Enter pool password" type="password" v-model="poolPasswd"></b-input>
                    <b-button class="m-3" @click="joinPool">Join pool</b-button>
                </div>
            </b-tab>
            <b-tab title="Create">
                <div class="d-flex flex-column align-items-center justify-content-center">
                    <h4 class="m-3">Create pool id and password:</h4>
                    <b-input class="m-3" placeholder="Enter new pool id..." v-model="poolId"></b-input>
                    <b-input class="m-3" placeholder="Enter pool password" type="password" v-model="poolPasswd"></b-input>
                    <b-input class="m-3" placeholder="Verify password" type="password" :state="poolPasswd === poolPasswdVerif" v-model="poolPasswdVerif"></b-input>
                    <b-button class="m-3" @click="createPool">Create pool</b-button>
                </div>
            </b-tab>
        </b-tabs>

    </b-modal>
</template>

<script>
    export default {
        name: "EnterPoolModal",
        data() {
            return {
                poolId: '',
                poolPasswd: '',
                poolPasswdVerif: '',
                alert: {
                    type: '',
                    text: '',
                    isShown: false,
                }
            }
        },
        methods: {
            joinPool() {
                this.$store.dispatch("TRY_JOIN_POOL", {poolId: this.poolId, password: this.poolPasswd}).then(result => {
                    console.log(result);
                    if (result.status) {
                        this.alert.text = "Success";
                        this.alert.type = "success";
                        this.alert.isShown = true;
                    } else {
                        this.alert.text = "Error";
                        this.alert.type = "danger";
                        this.alert.isShown = true;
                    }
                });
                this.poolId = this.poolPasswd = this.poolPasswdVerif = '';
            },
            createPool() {
                this.$store.dispatch("CREATE_POOL", {poolId: this.poolId, password: this.poolPasswd}).then(result => {
                    console.log(result);
                    if (result.status) {
                        this.alert.text = "Success";
                        this.alert.type = "success";
                        this.alert.isShown = true;
                    } else {
                        this.alert.text = "Error";
                        this.alert.type = "danger";
                        this.alert.isShown = true;
                    }
                });
                this.poolId = this.poolPasswd = this.poolPasswdVerif = '';
            }
        },
        computed: {
            isShown() {
                return this.$store.getters.GET_USER.poolId === null || this.$store.getters.GET_USER.poolId.length === 0
            }
        }
    }
</script>

<style scoped>

</style>