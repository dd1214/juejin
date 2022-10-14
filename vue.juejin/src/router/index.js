import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "@/pages/Home"
import Details from '@/pages/Details';
import PersionalPage from '@/pages/PersionalPage';
import Edit from '@/pages/Edit'
import MarketPage from '@/pages/MarketPage'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    {//重定向
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/details',
      name: 'xiangqingye',
      component: Details,
    },
    {
      path: '/PersionalPage',
      name: 'gerenzhuye',
      component: PersionalPage,
    },
    {
      path:'/Edit',
      name:'bianji',
      component:Edit,
    },
    {
      path: '/MarketPage',
      name:'shangcheng',
      component:MarketPage,
    }
  ]
})

export default router
