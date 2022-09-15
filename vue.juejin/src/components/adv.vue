<template>
  <div class="flex mgr8">
    <!-- 广告组件 -->
    <div class="flex1">
      <div class="sign column flexCenter mgt8">
        <div class="flex flexCenter ">
          <svg width="60" 
          height="30" 
          viewBox="0 0 24 24" >
          <path d="M8 2C8 1.72386 7.77614 1.5 7.5 1.5H6.5C6.22386 1.5 6 1.72386 6 2L5.9995 
          3H3C2.44772 3 2 3.47259 2 4.05556V7H22V4.05556C22 3.47259 21.5523 3 21 3H18V2C18 
          1.72386 17.7761 1.5 17.5 1.5H16.5C16.2239 1.5 16 1.72386 16 2V3H8V2ZM22 8.5H2V20
          .9444C2 21.5274 2.44772 22 3 22H21C21.5523 22 22 21.5274 22 20.9444V8.5Z" 
          fill="#FFCF8B" ></path>
              <rect x="5" y="12" width="3" height="2" rx="1" fill="#FF7D00" ></rect>
              <rect x="10.5" y="12" width="3" height="2" rx="1" fill="#FF7D00" ></rect>
              <rect x="5" y="16" width="3" height="2" rx="1" fill="#FF7D00" >
                </rect><rect x="10.5" y="16" width="3" height="2" rx="1" fill="#FF7D00" ></rect>
                <rect x="16" y="12" width="3" height="2" rx="1" fill="#FF7D00" ></rect>
                <rect x="16" y="16" width="3" height="2" rx="1" fill="#FF7D00" ></rect>
                </svg>
             <span style="font-size:18px;font-weight: 400">下午好！</span>
    
       <el-button round class="round-button">去签到</el-button>
       </div>
       <br>
       <span style="font-weight: 300;font-size:15px">点亮你在社区的每一天</span>
       </div>
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
.flex1{
  padding: 7px;
}
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
.sign{
  background-color:white;
  display:flex;
  height: 100px;
  width:240px;
}
.round-button{
  height:40px;
  border-color:#1e80ff;
  color:#1e80ff;
  font-size:14px;
}
</style>
