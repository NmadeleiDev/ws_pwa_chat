<template>
    <Page>
        <StackLayout ref="screen" orientation="vertical" width="*" height="*"
                     backgroundColor="lightgray">
            <Label text="Label 1" width="70" height="50" backgroundColor="red" />
            <Button text="Button" @tap="onButtonTap" />
            <ScrollView id="list" ref="list" orientation="vertical">
                <StackLayout height="auto" orientation="vertical">
                    <FlexboxLayout v-for="(message, key) in messages" :key="key"
                                   flexDirection="column" height="auto" width="80%">
                        <Label :text="message.user" />
                        <TextView editable="false">
                            <FormattedString>
                                <Span :text="message.text" />
                            </FormattedString>
                        </TextView>
                    </FlexboxLayout>
                </StackLayout>
            </ScrollView>
        </StackLayout>
    </Page>
</template>

<script>
    export default {
        data() {
            return {
                name: "greg",
                messages: [
                    {
                        user: "greg",
                        text: "Hey! How are you?"
                    },
                    {
                        user: "MiKe",
                        text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
                    },
                    {
                        user: "greg",
                        text: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur"
                    },
                    {
                        user: "Frank",
                        text: "Hey! Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
                    },
                    {
                        user: "Mila",
                        text: "Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
                    },
                    {
                        user: "Fred",
                        text: "Hi! Fine!"
                    },
                    {
                        user: "MiKe",
                        text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
                    },
                    {
                        user: "greg",
                        text: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur"
                    },
                    {
                        user: "Frank",
                        text: "Hey! Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
                    },
                    {
                        user: "Mila",
                        text: "Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
                    },
                    {
                        user: "Fred",
                        text: "Hi! Fine!"
                    }
                ],
                messageList: {}
            };
        },
        mounted() {
            let screen = this.$refs.screen.nativeView;
            console.log("SCREEEEEN: ", screen);
            this.messageList = this.$refs.list.nativeView;

            this.messageList.scrollToVerticalOffset(this.messageList
                    .scrollableHeight,
                true);
        },
        methods: {
            onButtonTap() {
                this.messageList.scrollToVerticalOffset(this.messageList
                    .scrollableHeight, true);
                console.log("Button was pressed ", this.messageList.scrollableHeight);
            }
        }
    };
</script>

<style>
</style>
