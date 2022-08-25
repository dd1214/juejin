<template>
  <router-view></router-view>
</template>

<script>

export default {
  name: 'App',

  mounted(){
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
