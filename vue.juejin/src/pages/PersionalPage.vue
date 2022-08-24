<template>
  <div>
    <Navbar/>
    <HomepageCenter/>
    <HomepageRight/>
    <PersonalEssays/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import HomepageCenter from '@/components/HomepageCenter.vue'
import HomepageRight from '@/components/HomepageRight.vue'
import PersonalEssays from '@/components/PersonalEssays.vue'


export default {
  name: 'PersionalPage',
  components: {
    Navbar,HomepageCenter,HomepageRight,PersonalEssays
  },
}
</script>

<style>
</style>
