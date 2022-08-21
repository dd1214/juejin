import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "../App.vue"
import atc from '../components/atc.vue';

Vue.use(VueRouter)

const router = new VueRouter({

  routes:[
    {//重定向
      path:'/',
      name:'home',
      component: Home,
    },
    {
      path:'/details',
      name:'xiangqingye',
      component: atc,
    }
  ]

})

export default router
