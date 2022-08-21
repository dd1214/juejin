import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "../App.vue"
import Details from '../components/details.vue';

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
      component: Details,
    }
  ]

})

export default router
