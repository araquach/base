import Vue from 'vue'
import App from './App.vue'
import Home from './components/Home.vue'
import Info from './components/Info.vue'

Vue.component('main-component', App)
Vue.component('home-component', Home)
Vue.component('info-component', Info)

const app = new Vue({
    el: '#app'
});