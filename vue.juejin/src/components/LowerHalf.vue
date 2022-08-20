<template>
    <div class="flex flex1" >
    <!-- 左边占位 -->
      <div style="flex:0.8;"></div>
        <!-- 文章栏 -->
      <div style="flex:4" class="flex column mg8">
          <!-- 推荐按钮 -->
           <div class="flex white shadow mgt8">
            <el-link href="" target="_blank" class="flex mg14 mgl20 font-size14">推荐</el-link>
            <el-link href="" target="_blank" class="flex mg14 font-size14">最新</el-link>
            <el-link href="" target="_blank" class="flex mg14 font-size14">热榜</el-link>
          </div>
          <!-- 无限滚动文章 -->
          <Essay EssayTitle=文章1 />
          <Essay EssayTitle=文章2 />
          <Essay EssayTitle=文章3 />
          <Essay EssayTitle=文章4 />
      </div>
      <!-- 右区 -->
      <div style="flex: 2" class="flex  mgr8 mgt8">
        <!-- 广告组件 -->
        <div  class="flex1">
        <Adv/>
        </div>
        <!-- 右边占位 -->
        <div  class="flex flex1 mg8"></div>
      </div>
    </div>
</template>

<script>
import Essay from './Essay'
import Adv from './adv.vue'
export default {
    name:'LowerHalf',
    data(){
        return {}
    },
    components:{Essay,Adv}
}
</script>