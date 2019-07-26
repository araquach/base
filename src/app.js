import Vue from 'vue'
import Home from './components/Home.vue'
import Info from './components/Info.vue'
import Register from './components/Registerf.vue'
import Success from './components/Success.vue'

Vue.component('home-component', Home)
Vue.component('info-component', Info)
Vue.component('register-component', Register)
Vue.component('success-component', Success)

const app = new Vue({
    el: '#app'
});