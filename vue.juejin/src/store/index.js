import Vue from 'vue'
import Vuex from 'vuex'
import loginOptions from './login'
import registerOptions from './register'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    loginOptions,
    registerOptions,
  }
})