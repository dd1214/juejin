<template>
  <div class="outside">
    <Navbar/>
    <div class="outside-row">
        
      
      <div class="outside">
    <HomepageCenter/>
    <PersonalEssays/>
    </div>
    
    <HomepageRight/>
     
      </div>
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

<style scoped>
.outside{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  
}
.outside-row{
  display: flex;
  flex-direction: row;
  justify-content: flex-star;
  align-items:flex-star;
}


</style>
