<template>
    <v-dialog
            v-model="dialog"
            fullscreen
    >
        <template v-slot:activator="{ on, attrs }">
            <v-btn light class="pb-3" icon
                   v-bind="attrs"
                   v-on="on">
                <v-icon light large color="white">attach_file</v-icon>
            </v-btn>
        </template>

        <v-card>
            <v-card-title class="headline grey lighten-2">
                {{chat ? chat.name : 'Send file'}}
            </v-card-title>

            <v-card-text class="fill-width d-flex flex-column justify-space-around align-center">
                <div>Files: {{files}}</div>
                <v-file-input v-model="files"
                              prepend-icon="mail"
                              class="fill-width"
                              full-width
                              accept="image/*"
                              label="Files to send"
                              placeholder="Choose files"
                              chips
                              show-size
                              counter
                              outlined
                              multiple></v-file-input>
                <v-divider></v-divider>
                <v-textarea
                        v-model="message"
                        class="fill-width"
                        full-width rows="2"
                        placeholder="Type message here"></v-textarea>
            </v-card-text>

            <v-divider></v-divider>

            <v-card-actions class="d-flex flex-row justify-space-around align-center">
                <v-btn
                        dark
                        @click="dialog = false"
                >
                    Close
                </v-btn>
                <v-btn
                        color="green"
                        @click="sendMessage(); dialog = false"
                >
                    Send
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
    import  Vue from 'vue';
    import {Chat} from "@/interfaces/main";


    export default Vue.extend( {
        name: "SendFileDialog",
        data: () => {
            return {
                files: [] as Array<Blob>,
                message: <string>'',
                dialog: false
            }
        },
        methods: {
            sendMessage() {
                this.$store.dispatch('sendMessageInChat', {message: this.message, files: this.files})
            }
        },
        computed: {
            chat(): Chat | undefined {
                return this.$store.getters.getCurrentChat()
            },
            username(): string {
                return this.$store.getters.username()
            }
        }
    })
</script>

<style scoped>

</style>