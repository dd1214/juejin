import Vue from 'vue'
import Vuex from 'vuex'
import loginOptions from './login'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    loginOptions
  }
})