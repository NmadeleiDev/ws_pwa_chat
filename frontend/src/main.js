import Vue from 'vue';
import {store} from './store/index';
import App from './App.vue';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import router from "./router/router";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.config.productionTip = false;
export const alertNotifierChannel = new Vue();

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
