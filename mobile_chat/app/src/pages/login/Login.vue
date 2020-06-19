<template>
    <Page actionBarHidden="true" @loaded="passLogin">
        <FlexboxLayout class="page">
            <StackLayout class="form">
                <Label class="header" text="ENC & TALK"></Label>

                <GridLayout rows="auto, auto, auto">
                    <StackLayout row="0" class="input-field">
                        <TextField class="input" hint="username" :isEnabled="!processing"
                            keyboardType="username" autocorrect="false"
                            autocapitalizationType="none" v-model="user.username"
                            returnKeyType="next" @returnPress="focusPassword"></TextField>
                        <StackLayout class="hr-light"></StackLayout>
                    </StackLayout>

                    <StackLayout row="1" class="input-field">
                        <TextField class="input" ref="password" :isEnabled="!processing"
                            hint="Password" secure="true" v-model="user.password"
                            :returnKeyType="isLoggingIn ? 'done' : 'next'"
                            @returnPress="focusConfirmPassword"></TextField>
                        <StackLayout class="hr-light"></StackLayout>
                    </StackLayout>

                    <StackLayout row="2" v-show="!isLoggingIn" class="input-field">
                        <TextField class="input" ref="confirmPassword" :isEnabled="!processing"
                            hint="Confirm password" secure="true" v-model="user.confirmPassword"
                            returnKeyType="done"></TextField>
                        <StackLayout class="hr-light"></StackLayout>
                    </StackLayout>

                    <ActivityIndicator rowSpan="3" :busy="processing"></ActivityIndicator>
                </GridLayout>

                <Button :text="isLoggingIn ? 'Log In' : 'Sign Up'" :isEnabled="!processing"
                    @tap="submit" class="btn btn-primary m-t-20"></Button>
                <Button text="init" @tap="check">Init</Button>
                <Label *v-show="isLoggingIn" text="Forgot your password?"
                    class="login-label" @tap="forgotPassword()"></Label>
            </StackLayout>

            <Label class="login-label sign-up-label" @tap="toggleForm">
                <FormattedString>
                    <Span :text="isLoggingIn ? 'Don’t have an account? ' : 'Back to Login'"></Span>
                    <Span :text="isLoggingIn ? 'Sign up' : ''" class="bold"></Span>
                </FormattedString>
            </Label>
        </FlexboxLayout>
    </Page>
</template>

<script>
    // import Chat from "../chat/Chat";
    import Menu from "~/src/pages/menu/Menu";

    export default {
        data() {
            return {
                isLoggingIn: true,
                processing: false,
                user: {
                    username: "",
                    password: "",
                    confirmPassword: ""
                }
            };
        },
        created() {
            this.$navigateTo(Menu, { clearHistory: true });
        },
        methods: {
            passLogin() {
                // this.$navigateTo(Menu, { clearHistory: true });
                this.$store.dispatch("INIT_SQLITE_CONN").then(res => {
                    if (res === false) {
                        this.alert("Failed to connect to local memory storage");
                        return;
                    }
                    this.$store.dispatch("LOAD_ACTIVE_USER_SECRET_KEYS").then(ok => {
                        if (ok === false) {
                            console.log("No secret keys found");
                            return;
                        }
                        this.$store.dispatch("LOAD_USER_DATA").then(result => {
                            if (result === false) {
                                console.log("Failed to load user data")
                                return
                            }
                            console.log("Navigating to chat!")
                            this.$navigateTo(Menu, { clearHistory: true });
                        }).catch(e => console.log("ERROR: ", e));
                    }).catch(e => console.log("ERROR: ", e));
                }).catch(e => console.log("ERROR: ", e));
            },
            toggleForm() {
                this.isLoggingIn = !this.isLoggingIn;
            },
            check() {
                // this.$navigateTo(Menu, { clearHistory: true });

                this.$store.dispatch("INIT_SQLITE_CONN");
            },
            submit() {
                if (!this.user.username || !this.user.password) {
                    this.alert(
                        "Please provide both an username and password."
                    );
                    return;
                }

                this.processing = true;
                if (this.isLoggingIn) {
                    this.login();
                } else {
                    this.register();
                }
            },

            login() {
                console.log("Login logic");
                this.$store.dispatch("SIGN_IN", {username: this.user.username, password: this.user.password}).then(res => {
                    if (res === true) {
                        this.$navigateTo(Menu, { clearHistory: true });
                    } else {
                        console.log("Signin failed");
                    }
                });
                this.processing = false;
            },

            register() {
                console.log("Register logic");
                if (this.user.password !== this.user.confirmPassword) {
                    this.alert("Your passwords do not match.");
                    this.processing = false;
                    return;
                }
                // здесь обращение в store для регистрации
                this.$store.dispatch("SIGN_UP", {username: this.user.username, password: this.user.password}).then(res => {
                    if (res === true) {
                        console.log("Signup successed, moving to chat.");
                        this.$navigateTo(Menu, { clearHistory: true });
                    } else {
                        console.log("Signup failed");
                    }
                });
                this.processing = false;
            },

            forgotPassword() {
                // возмонжо сделаю обработку восстановления
            },

            focusPassword() {
                this.$refs.password.nativeView.focus();
            },
            focusConfirmPassword() {
                if (!this.isLoggingIn) {
                    this.$refs.confirmPassword.nativeView.focus();
                }
            },

            alert(message) {
                return alert({
                    title: "ENC & TALK",
                    okButtonText: "OK",
                    message: message
                });
            }
        }
    };
</script>

<style scoped>
    .page {
        align-items: center;
        flex-direction: column;
    }

    .form {
        margin-left: 30;
        margin-right: 30;
        flex-grow: 2;
        vertical-align: middle;
    }

    .logo {
        margin-bottom: 12;
        height: 90;
        font-weight: bold;
    }

    .header {
        horizontal-align: center;
        font-size: 25;
        font-weight: 600;
        margin-bottom: 70;
        text-align: center;
        color: #D51A1A;
    }

    .input-field {
        margin-bottom: 25;
    }

    .input {
        font-size: 18;
        placeholder-color: #A8A8A8;
    }

    .input:disabled {
        background-color: white;
        opacity: 0.5;
    }

    .btn-primary {
        margin: 30 5 15 5;
    }

    .login-label {
        horizontal-align: center;
        color: #A8A8A8;
        font-size: 16;
    }

    .sign-up-label {
        margin-bottom: 20;
    }

    .bold {
        color: #000000;
    }
</style>
