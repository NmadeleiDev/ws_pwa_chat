import Vue from "nativescript-vue";
import Store from "./src/store"
import routes from "./src/routes";
import RadSideDrawer from "nativescript-ui-sidedrawer/vue";
Vue.use(RadSideDrawer);

// Vue.config.silent = (TNS_ENV === 'production');
Vue.config.silent = false;

new Vue({
    store:Store,
    RadSideDrawer,
    // render: h => h("frame", [h(Store.getters.GET_USER.secretKey.length === 0 ? routes.login : routes.chat)])
    render: h => h("frame", [h(routes.login)])
}).$start();
