import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "@/pages/Home"
import Details from '@/pages/Details';
import PersionalPage from '@/pages/PersionalPage';

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
    }
  ]
})

export default router
