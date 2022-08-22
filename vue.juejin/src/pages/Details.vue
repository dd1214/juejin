<template>
  <div>
    <!-- 上半区 -->
    <Navbar />
    <!-- 下半区 -->
    <!-- <Atc /> -->
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import Atc from '@/components/atc.vue'

export default {
  name: 'App',
  components: {
    Navbar,
    Atc,
  },
}
</script>

<style></style>
