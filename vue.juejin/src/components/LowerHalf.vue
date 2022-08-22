<template>
  <div class="lh-container">
    <!-- 左边文章区 -->
    <Essays />
    <!-- 右区广告区 -->
    
    <Adv class="ads"/>
  </div>
</template>

<script>
import Essays from './Essays'
import Adv from './adv.vue'

export default {
  name: 'LowerHalf',
  data() {
    return {}
  },
  components: { Essays, Adv },
}
</script>

<style scoped lang="less">
.lh-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  overflow-y: hidden;
  // .ads {
  // }
  @media screen and (max-width: 1000px) {
    .ads {
      display: none;
    }
  }
}
</style>
