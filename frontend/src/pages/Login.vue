<template>
    <div>
        <div class="form-container">
            <label class="form-item">
                <h3>Name:</h3>
                <input type="text" name="username" v-model="user.username">
            </label>
            <label class="form-item">
                <h3>Password:</h3>
                <input type="password" name="password" v-model="user.password">
            </label>
            <div class="buttons-container">
                <button @click="signIn">Sign In</button>
                <button @click="signUp">Sign up</button>
            </div>
        </div>
    </div>
</template>

<script>
    import api from "../api/api";

    export default {
        name: "Login",
        data() {
            return {
                user: {
                    username: '',
                    password: '',
                }
            }
        },
        methods: {
            signIn() {
                const that = this;
                const success =api.post("signin", this.user);
                success.then(data => {
                    if (data.status === true) {
                        that.$router.push("/chat");
                    }
                });
            },
            signUp() {
                const that = this;
                const success =api.post("signup", this.user);
                success.then(data => {
                    if (data.status === true) {
                        that.$router.push("/chat");
                    }
                });
            }
        }
    }
</script>

<style scoped>

    .form-container {
        display: flex;
        flex-direction: column;
        justify-content: center;
        width: 60%;
        border: solid darkred;
        background-color: lightgray;
        margin: 4em auto;
    }

    .form-item {
        margin: 1em;
    }

    .buttons-container {
        width: 40%;
        margin: 2em auto;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }
</style>