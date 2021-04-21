import Vue from 'vue';
import VueRouter from "vue-router";
import App from './App.vue';
import vuetify from './plugins/vuetify';
import VueParticles from 'vue-particles';
import routers from "@/routers";

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueParticles)

const router = new VueRouter({
    mode: 'history',
    routes: routers,
})

new Vue({
    vuetify,
    router,
    render: h => h(App),
    created() {
        localStorage.clear()
    }
}).$mount('#app')
