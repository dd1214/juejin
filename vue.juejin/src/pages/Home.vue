<template>
<div>
  <!-- 上半区 -->
  <Navbar/>
  <LinkPop/>
  <!-- 下半区 -->
  <LowerHalf/>
</div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import LowerHalf from '@/components/LowerHalf.vue'
import LinkPop from '@/components/Linkpop.vue'


export default {
  name: 'Home',
  components: {
   Navbar,LowerHalf,LinkPop
  }
}
</script>

<style>
</style>
