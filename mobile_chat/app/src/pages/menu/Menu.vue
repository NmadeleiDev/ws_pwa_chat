<template>
    <Page class="page" @loaded="appLoaded">
        <ActionBar :title="account + '`s chats'" marginTop="10"/>
        <AbsoluteLayout ref="rootLayout" marginTop="10" marginRight="5">
            <StackLayout height="100%" width="100%">
                <TextField v-model="filter" hint="Search chat"/>
                <ListView v-if="chats.length > 0"
                          @tap="$store.dispatch('CHANGE_CURRENT_CHAT', item.id)"
                          for="item in chats"
                          marginLeft="20"
                          height="100%" width="100%" marginBottom="0">
                    <v-template>
                        <FlexboxLayout alignItems="flex-start"
                                       width="100%"
                                       height="auto"
                                       justifyContent="flex-start"
                                       flexDirection="column">
                            <FlexboxLayout height="30"
                                           width="100%"
                                           class="chat-item-inner"
                                           alignContent="center"
                                           justifyContent="space-between"
                                           flexDirection="row"
                                           flexGrow="6">
                                <Label :text="item.name"
                                       flexGrow="2"
                                       class="chat-name"
                                       alignSelf="flex-start"/>
                                <Label textAlignment="right"
                                       marginRight="30"
                                       :text="item.messages.filter(i => i.isNew).length + 's ' + item.messages.slice(-1)[0].time.getHours() + ':' + item.messages.slice(-1)[0].time.getMinutes()"
                                       flexGrow="1"/>
                            </FlexboxLayout>
                            <FlexboxLayout height="30"
                                           width="auto"
                                           justifyContent="space-between"
                                           alignItems="flex-start"
                                           alignContent="flex-start"
                                           flexDirection="row"
                                           flexGrow="1">
                                <Label height="20">
                                    <FormattedString>
                                        <Span :text="item.messages.slice(-1)[0].sender + ': '"
                                              style="color: black"></Span>
                                        <Span :text="item.messages.slice(-1)[0].text"></Span>
                                    </FormattedString>
                                </Label>
                            </FlexboxLayout>
                        </FlexboxLayout>
                    </v-template>
                </ListView>

                <Label v-else text="You don't have any chats yet"></Label>
            </StackLayout>


            <StackLayout left="0" top="0" height="100%" width="100%" class="backdrop" :class="classBackdrop"/>

            <AbsoluteLayout ref="fabItemPosition" marginTop="87%" marginLeft="80%">
                <GridLayout ref="fabItemContainer" left="8" top="8">
                    <FloatButtonItem row="1" :class="classItem1" color="#E94E77" title="E"/>
                    <FloatButtonItem row="1" :class="classItem2" color="#3FB8AF" title="U"/>
                    <FloatButtonItem row="1" :class="classItem3" color="#FCB653" title="V"/>
                </GridLayout>
                <FloatButton @onButtonTap="onButtonTap" :isActive="isActive"/>
            </AbsoluteLayout>
        </AbsoluteLayout>
    </Page>
</template>

<script>
    import FloatButton from "~/src/widgets/FloatButton";
    import FloatButtonItem from "~/src/widgets/FloatButtonItem";

    const app = require('tns-core-modules/application')
    const platform = require('tns-core-modules/platform')

    export default {
        mounted() {
            // let maxl = this.fish.length - 1;
            // let words = this.fish.split(" ");
            // let maxc = words.length - 1;
            //
            // for (let i = 1; i < 21; i++) {
            //     let date = new Date(Math.floor(Math.random() * Math.floor(20)) + 2000, Math.floor(Math.random() * Math.floor(12)), Math.floor(Math.random() * Math.floor(28)), Math.floor(Math.random() * Math.floor(59)), Math.floor(Math.random() * Math.floor(59)), Math.floor(Math.random() * Math.floor(59)), 0);
            //     let chat = {
            //         id: Math.floor(Math.random() * Math.floor(10000)) + i,
            //         name: words[Math.floor(Math.random() * Math.floor(maxc))],
            //         messages: [{
            //             sender: words[Math.floor(Math.random() * Math.floor(maxc))],
            //             isNew: Math.floor(Math.random() * Math.floor(10)) >= 5,
            //             text: this.fish.substring(Math.random() * Math.floor(maxl / 2), Math.random() * Math.floor(maxl / 2) + maxl / 2),
            //             time: date
            //         }, {
            //             isNew: Math.floor(Math.random() * Math.floor(10)) >= 5,
            //             sender: words[Math.floor(Math.random() * Math.floor(maxc))],
            //             text: this.fish.substring(Math.random() * Math.floor(maxl / 2), Math.random() * Math.floor(maxl / 2) + maxl / 2),
            //             time: date
            //         }]
            //     }
            //     this.chats.push(chat);
            // }
            console.log("Len 1 = ", this.chats.length);
        },
        data() {
            return {
                textFieldValue: "",
                filter: '',
                isActive: false,
                // chats: [],

                fish: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.'
            };
        },
        computed: {
            chats() {
                return this.$store.getters.GET_USER.chats;
            },
            account() {
                return 'Greg';
            },

            classItem1() {
                return this.isActive ? "raiseItem1" : "downItem1"
            },
            classItem2() {
                return this.isActive ? "raiseItem2" : "downItem2"
            },
            classItem3() {
                return this.isActive ? "raiseItem3" : "downItem3"
            },
            classBackdrop() {
                return this.isActive ? "backdrop-visible" : "backdrop-invisible"
            },
        },
        methods: {
            appLoaded(args) {

                let fabItemContainer = this.$refs.fabItemContainer.nativeView
                let fabItemPosition = this.$refs.fabItemPosition.nativeView
                let rootLayout = this.$refs.rootLayout.nativeView

                // Needed to avoid masking child components on Android
                if (app.android && platform.device.sdkVersion >= '21') {
                    fabItemContainer.android.setClipChildren(false)
                    fabItemPosition.android.setClipChildren(false)
                    rootLayout.android.setClipChildren(false)
                }
            },
            onItemTap(args) {
                console.log("Len 2 = ", this.chats.length);
                console.log('Tapped cell: ' + args)
            },
            onButtonTap(args) {
                this.isActive = !this.isActive
            },
            hourOfDay() {
                let hours = new Date(Date.now()).getHours();
                console.log("HOUR: ", hours);
                return hours;
            }
        },
        components: {FloatButton, FloatButtonItem}
    }

</script>

<style scoped>
    label {
        margin: 0;
        padding: 0;
    }

    .chat-item-container {
        /*margin: 2px 10px 0 15px;*/
    }

    .chat-item-inner {

    }

    .chat-name {
        font-size: 20px;
        font-weight: bold;
        color: black;
    }

    .last-message {
        font-size: 14px;
    }


    ListView Label {
        height: 48;
        min-height: 48;
    }

    .backdrop {
        background-color: rgba(29, 30, 35, .90);
        opacity: 0;
    }

    .backdrop-visible {
        animation-name: activateBackdrop;
        animation-duration: .25;
        animation-fill-mode: forwards;
    }

    .backdrop-invisible {
        animation-name: activateBackdrop;
        animation-duration: .25;
        animation-fill-mode: forwards;
        animation-direction: reverse;
    }

    .raiseItem1 {
        opacity: 1;
        animation-name: raiseItem1;
        animation-duration: .25;
        animation-timing-function: cubic-bezier(0.165, 0.840, 0.440, 1.000);
        animation-fill-mode: forwards;
    }

    .raiseItem2 {
        opacity: 1;
        animation-name: raiseItem2;
        animation-duration: .25;
        animation-timing-function: cubic-bezier(0.165, 0.840, 0.440, 1.000);
        animation-fill-mode: forwards;
    }

    .raiseItem3 {
        opacity: 1;
        animation-name: raiseItem3;
        animation-duration: .25;
        animation-timing-function: cubic-bezier(0.165, 0.840, 0.440, 1.000);
        animation-fill-mode: forwards;
    }

    .downItem1 {
        animation-name: raiseItem1;
        animation-duration: .20;
        animation-timing-function: cubic-bezier(0.895, 0.030, 0.685, 0.220);
        animation-fill-mode: forwards;
        animation-direction: reverse;
    }

    .downItem2 {
        animation-name: raiseItem2;
        animation-duration: .20;
        animation-timing-function: cubic-bezier(0.895, 0.030, 0.685, 0.220);
        animation-fill-mode: forwards;
        animation-direction: reverse;
    }

    .downItem3 {
        animation-name: raiseItem3;
        animation-duration: .20;
        animation-timing-function: cubic-bezier(0.895, 0.030, 0.685, 0.220);
        animation-fill-mode: forwards;
        animation-direction: reverse;
    }

    @keyframes activateBackdrop {
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }

    @keyframes raiseItem1 {
        from {
            opacity: 1;
            transform: translate(0, 0);
        }
        to {
            opacity: 1;
            transform: translate(0, -64);
        }
    }

    @keyframes raiseItem2 {
        from {
            opacity: 1;
            transform: translate(0, 0);
        }
        to {
            opacity: 1;
            transform: translate(0, -128);
        }
    }

    @keyframes raiseItem3 {
        from {
            opacity: 1;
            transform: translate(0, 0);
        }
        to {
            opacity: 1;
            transform: translate(0, -192);
        }
    }
</style>
