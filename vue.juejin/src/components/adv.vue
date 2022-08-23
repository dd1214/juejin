<template>
  <div class="flex mgr8">
    <!-- 广告组件 -->
    <div class="flex1">
      <div class="mgt8 ">
        <div class="outsideDiv" >
          <img :src="adv1" class="" v-show="item1" />
          <div class="insideDiv">
            <i class="el-icon-close" v-show="item1" @click="click1"></i>
          </div>
        </div>
        
        <div class="outsideDiv mgt8" >
          <img :src="adv2" class="" v-show="item2" onclick="window.open('https://juejin.cn/pin/7129334097113006116?utm_source=slide&utm_medium=banner&utm_campaign=reading')"/>
          <div class="insideDiv">
            <i class="el-icon-close" v-show="item2" @click="click2"></i>
          </div>
        </div>

        <div class="outsideDiv mgt8">
          <img :src="adv3" class="" v-show="item3"  onclick="window.open('https://juejin.cn/book/7126538479051210766?utm_source=web_banner&utm_medium=banner&utm_campaign=Book_SK_0817')"/>
          <div class="insideDiv">
            <i class="el-icon-close" v-show="item3" @click="click3"></i>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'adv',
  data() {
    return {
      adv1: require('..//assets/adv1.jpg'),
      adv2: require('..//assets/adv2.jpg'),
      adv3: require('..//assets/adv3.jpg'),
      item1: true,
      item2: true,
      item3: true,
    }
  },
  methods: {
    click1() {
      this.item1 = false
      this.css2 = {}
      if (this.item2 == false) {
        this.css3 = {}
      } else {
        this.css3 = { top: '340px' }
      }
    },
    click2() {
      this.item2 = false
      if (this.item1 == false) {
        this.css3 = {}
      } else if (this.item1 == true) {
        this.css3 = { top: '340px' }
      }
    },
    click3() {
      this.item3 = false
    },
  },
}
</script>
<style scoped>
.outsideDiv {
  position: relative;
}
.insideDiv {
  position: absolute;
  right: 0;
  top: 0;
}
img{
    height: 200px;
}
i{
    cursor: pointer;
}
</style>
