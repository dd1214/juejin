<template>
<div>
  <!-- 上半区 -->
  <Navbar/>
  <!-- 下半区 -->
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
        <Adv advTitle=广告1 />
        <Adv advTitle=广告2 />
        <Adv advTitle=广告3 />
        </div>
        <!-- 右边占位 -->
        <div  class="flex flex1 mg8"></div>
      </div>
    </div>
</div>
</template>

<script>
import Essay from './components/Essay'
import Adv from './components/adv.vue'
import Navbar from '@/components/Navbar.vue'

export default {
  name: 'App',
  components: {
    Essay,Adv,Navbar
  }
}
</script>

<style>
body {
    margin: 0;
    height: 100vh;
    background-color: #eee;
  }
  /* * {
    border: 1px solid #333;
  }    */
  .flex {
    display: flex;
  }
  .flexCenter {
    justify-content: center;
    align-items: center;
  }
  /* 占满比例 */
  .flex1 {
    flex: 1;
  }
  /* 竖向 */
  .column {
    flex-direction: column;
  }
  /* 阴影 */
  .shadow {
    box-shadow: rgba(0, 0, 0, 0) 0px 2px 1px -1px,
      rgb(0 0 0 / 1%) 0px 1px 1px 0px, rgb(0 0 0 / 12%) 0px 1px 3px 0px;
  }
  .font-size14{
    font-size:14px
  }   
  /* 白色背景 */
  .white {
    background-color: #fff;
  }
  /* 上下左右外边距 */
  .mg8 {
    margin: 8px;
  }
  .mg17 {
    margin: 11px;
    margin-left:17px;
  }
  .mg10{
    margin: 10px;
  }
  .mg14{
    margin: 14px;
  }
  .mg25{
    margin-left:25px
  }
  /* 左外边距 */
  .mgl20 {
    margin-left: 20px;
  }
  /* 右外边距 */
  .mgr8 {
    margin-right: 8px;
  }
  /* 上外边距 */
  .mgt8 {
    margin-top: 8px;
  }
  /* 下外边距 */
  .mgb1 {
    margin-bottom: 1px;
  }
  /* 文章 */
  .article {
    height: 160px;
  }
  
 .none{
    display: none;
 }
 .h200w240{
    height:200px;
    width: 240px;
 }
</style>
