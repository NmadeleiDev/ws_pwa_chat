import Vue from 'vue';
import {store} from './store/index';
import App from './App.vue';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import router from "./router/router";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import VueNativeSock from 'vue-native-websocket'
Vue.use(VueNativeSock, "ws://" + window.location.host + "/ws/connect", {
  connectManually: true,
  reconnection: true,
  reconnectionAttempts: 5,
}, {store: store}, );


Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
