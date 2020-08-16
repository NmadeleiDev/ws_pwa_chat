import { Getters, Mutations, Actions, Module } from 'vuex-smart-module'

// State
class CommonNotificationState {
    show: boolean = false
    type: string = 'success'
    text: string = ''
}

class CommonNotificationGetters extends Getters<CommonNotificationState> {
    isShown(): boolean {
        return this.state.show
    }

    notificationText(): string {
        return this.state.text
    }

    notificationType(): string {
        return this.state.type
    }
}

class CommonNotificationMutations extends Mutations<CommonNotificationState> {

    setCommonNotificationValue(payload: boolean) {
        this.state.show = payload
    }

    setCommonNotificationText(payload: string) {
        this.state.text = payload
    }

    setCommonNotificationType(payload: string) {
        this.state.type = payload
    }
}

class CommonNotificationActions extends Actions<
    CommonNotificationState,
    CommonNotificationGetters,
    CommonNotificationMutations,
    CommonNotificationActions
    > {

    showCommonNotification(payload: {text: string, type: string}) {
        this.commit('setCommonNotificationText', payload.text)
        this.commit('setCommonNotificationType', payload.type)
        this.commit('setCommonNotificationValue', true)
    }

    setCommonNotificationValue(payload: boolean) {
        this.commit('setCommonNotificationValue', payload)
    }
}

export const CommonNotification = new Module({
    namespaced: false,
    state: CommonNotificationState,
    getters: CommonNotificationGetters,
    mutations: CommonNotificationMutations,
    actions: CommonNotificationActions,
})
