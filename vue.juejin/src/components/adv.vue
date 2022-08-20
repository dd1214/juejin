<template>
        <div style="flex: 2" class="flex  mgr8 mgt8">
            <!-- 广告组件 -->
            <div  class="flex1">
            <div class="mgt8 h200w240" >
            <div class="outsideDiv" >
        <img :src="adv1" class="h200w240" v-show="item1">     
            <div class="insideDiv">
            <i class="el-icon-close" v-show="item1" @click="click1"></i>
            </div>
            </div> 
            
        <div class="outsideDiv" :style="css2">
            <img :src="adv2" class="h200w240" v-show="item2">     
            <div class="insideDiv">
            <i class="el-icon-close" v-show="item2" @click="click2"></i>
            </div>
            </div> 
            
        <div class="outsideDiv" :style="css3">
            <img :src="adv3" class="h200w240" v-show="item3">     
            <div class="insideDiv">
            <i class="el-icon-close" v-show="item3" @click="click3"></i>
            </div>
            </div> 
             </div>
        <!-- 右边占位 -->
         <div  class="flex flex1 mg8"></div>
            </div>
             </div>
</template>

<script>
export default {
    name:'adv', 
    data(){
        return {
            adv1:require('..//assets/adv1.jpg'),
            adv2:require('..//assets/adv2.jpg'),
            adv3:require('..//assets/adv3.jpg'),
            item1:true,
            item2:true,
            item3:true,
            css2:{
                top:"340px"
            },
            css3:{
                top:"555px"
            }
        }
    },
    methods:{
        click1(){
            this.item1=false;
            this.css2={};
            if(this.item2==false){this.css3={}}
            else{this.css3={top:"340px"}}
            
        },
        click2(){
            this.item2=false
            if(this.item1==false){this.css3={}}
            else if(this.item1==true){this.css3={top:"340px"}}
        },
        click3(){
            this.item3=false
        }
    }
}
</script>
<style scoped>
.outsideDiv {
    position:absolute;
    }
.insideDiv {
    position:absolute;
    right:0;
    top:0;}
</style>