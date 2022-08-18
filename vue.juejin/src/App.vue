<template>
<div>
  <!-- 上半区 -->
  <Navbar/>
  <!-- 下半区 -->
  <LowerHalf/>
</div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import LowerHalf from '@/components/LowerHalf.vue'


export default {
  name: 'App',
  components: {
   Navbar,LowerHalf
  }
}
</script>

<style>
</style>
