import Vue from 'vue'
import Home from './components/Home.vue'
import Info from './components/Info.vue'
import Registera from './components/Registera.vue'
import Registerf from './components/Registerf.vue'
import Success from './components/Success.vue'

Vue.component('home-component', Home)
Vue.component('info-component', Info)
Vue.component('registera-component', Registera)
Vue.component('registerf-component', Registerf)
Vue.component('success-component', Success)

const app = new Vue({
    el: '#app'
});