import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "../views/Home.vue"
import Article from '../views/Article.vue';

Vue.use(VueRouter)

const router = new VueRouter({
  mode:'history',
  routes:[
    {//重定向
      path:'/',
      name:'home',
      component: Home,
    },
    {
      path:'/article',
      name:'xiangqingye',
      component: Article,
    }
  ]



})

export default router
