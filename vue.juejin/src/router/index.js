import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "@/App.vue"
import Details from '@/pages/Details';

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    {//重定向
      path: '/',
      name: 'home',
      component: Home,
      children: [
        {
          path: 'details',
          name: 'xiangqingye',
          component: Details,
        }
      ]
    },
    // {
    //   path: '/details',
    //   name: 'xiangqingye',
    //   component: Details,
    // }
  ]

})

export default router
