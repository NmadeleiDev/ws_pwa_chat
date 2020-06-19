<template>
    <b-container fluid>
        <b-alert v-model="showAlert" variant="danger" dismissible>
            {{ alertText }}
        </b-alert>
        <b-row class="m-4 mt-5">
            <b-col>
                <h1>Welcome back</h1>
            </b-col>
        </b-row>
        <b-row class="m-4">
            <b-col>
                <label>
                    Имя
                    <b-form-input size="lg" v-model="user.username" placeholder="Enter your nick"></b-form-input>
                </label>
            </b-col>
        </b-row>
        <b-row class="m-4">
            <b-col>
                <label>
                    Пароль
                    <b-form-input size="lg" v-model="user.password" :state="passwordOk" type="password" placeholder="Enter your password"></b-form-input>
                </label>
            </b-col>
        </b-row >
        <b-collapse id="collapse-1" class="mt-2">
        <b-row class="m-4">
            <b-col>
                <label>
                    Пароль еще раз
                    <b-form-input size="lg" v-model="passwordVerif" :state="passwordOk" type="password" placeholder="Enter password again"></b-form-input>
                </label>
            </b-col>
        </b-row>
        </b-collapse>
        <b-row class="m-4 mt-5">
            <b-col>
                <b-button variant="success" @click="isCreate ? signUp() : signIn()" size="lg" :block="isMobile">
                    {{ isCreate ? "Создать аккаут" : "Войти" }}
                </b-button>
            </b-col>
        </b-row>
        <b-row>
            <b-col>
                <b-button v-b-toggle.collapse-1 variant="outline-secondary" @click="changeAction">
                    {{ isCreate ? "У меня есть аккаут" : "У меня нет аккаутна" }}
                </b-button>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
    import api from "../../api/api";

    export default {
        name: "Login",
        data() {
            return {
                user: {
                    username: '',
                    password: '',
                },
                passwordVerif: '',
                isCreate: false,
                showAlert: false,
                alertText: false,
                isMobile: true,
            }
        },
        created() {
            this.isMobile = window.screen.width > 500;
        },
        methods: {
            signIn() {
                const that = this;
                const success = api.post("signin", this.user);
                success.then(data => {
                    if (data.status === true) {
                        that.$router.push("/chat");
                    } else {
                        this.showAlert = true;
                        this.alertText = "Имя и пароль не совпадают."
                    }
                });
            },
            signUp() {
                const that = this;
                const success = api.post("signup", this.user);
                success.then(data => {
                    if (data.status === true) {
                        that.$router.push("/chat");
                    } else {
                        this.showAlert = true;
                        this.alertText = "Не удалось создать аккаунт. Попробуйте позже."
                    }
                });
            },
            changeAction() {
                this.isCreate = !this.isCreate;
            }
        },
        computed: {
            passwordOk() {
                if (!this.isCreate || (this.user.password.length === 0)) {
                    return null;
                } else if (this.user.password === this.passwordVerif) {
                    return true;
                }
                return false;
            }
        }
    }
</script>

<style scoped>

</style>