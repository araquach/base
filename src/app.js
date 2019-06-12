import Vue from 'vue'
import Home from './components/Home.vue'
import Info from './components/Info.vue'
import Register from './components/Register.vue'

Vue.component('home-component', Home)
Vue.component('info-component', Info)
Vue.component('register-component', Register)

const app = new Vue({
    el: '#app'
});