import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import '@/css/base.css'
import '@/css/element-variables.scss'
import css from './assets/css.css'

import store from '@/store'
Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
