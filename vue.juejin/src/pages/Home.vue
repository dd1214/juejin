<template>
<div>
  <!-- 上半区 -->
  <Navbar/>
  <LinkPop/>
  <!-- 注册登陆 -->
  <!-- <Land/>
  <Register/> -->
  <!-- 下半区 -->
  <LowerHalf/>
  

</div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import LowerHalf from '@/components/LowerHalf.vue'
import LinkPop from '@/components/Linkpop.vue'
import Land from '@/components/Land.vue'
import Register from '@/components/Register.vue'


export default {
  name: 'Home',
  components: {
   Navbar,LowerHalf,LinkPop,Land,Register
  }
}
</script>

<style>
</style>
