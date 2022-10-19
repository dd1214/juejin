<template>
  <router-view></router-view>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'  // 安装axios后引入
Vue.prototype.$axios = axios  // 将axios挂载到原型上方便使用

export default {
  name: 'App',
  data() {
    return {
    }
  },
  mounted(){
    this.$axios.get('./article.json').then(res => {   
        window.localStorage.setItem('atc',JSON.stringify(res.data)) //将文章数据存在本地存储
  })
    if(localStorage.getItem('token')){
      this.$store.commit('loginOptions/LOGIN')
    }
    // 监听 localStorage，如果有人删除了 token，就登出
    window.addEventListener('storage', e => {
      if(!localStorage.getItem('token')){
        this.$store.commit('loginOptions/LOGOUT')
      }else{
        this.$store.commit('loginOptions/LOGIN')
      }
    })
  }
}
</script>