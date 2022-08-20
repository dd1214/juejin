<template>
        <div class="mgt8 h200w240" >
            <i class="el-icon-close" v-show="item1" @click="click1"></i>
            <img :src="adv1" class="h200w240" v-show="item1"> 

            <i class="el-icon-close" v-show="item2" @click="click2"></i>
            <img :src="adv2" class="h200w240" v-show="item2"> 

            <i class="el-icon-close" v-show="item3" @click="click3"></i>
            <img :src="adv3" class="h200w240" v-show="item3">
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
        }
    },
    methods:{
        click1(){
            this.item1=false
        },
        click2(){
            this.item2=false
        },
        click3(){
            this.item3=false
        }
    }
}
</script>
<style scoped>

</style>