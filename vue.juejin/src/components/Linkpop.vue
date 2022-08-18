<!-- 弹出框，被用在主导航栏的下半部分的弹出显示 -->s
<template>
<ul>
  <LinkpopItem v-for="(item) in lists" :key="item.id" :index="item.id" :info="item.info"/>
</ul>
</template>

<script>
import LinkpopItem from '@/components/LinkpopItem';
export default {
  name: 'Linkpop',
  components: {
    LinkpopItem
  },
  data(){
    return {
      active_index: 1,
      // 测试数据
      lists: [
        {
          id: 1,
          info: {
            main: '综合'
          }
        },
        {
          id: 2,
          info: {
            main: '前端',
            sub: ['前端', 'JavaScript', 'Vue.js']
          }
        }
      ]
    }
  }
}
</script>

<style scoped lang="less">
ul{
  display: flex;
  justify-content: flex-start;
  align-items: stretch;

  width: auto;
  height: 46px;
  padding: 10px 20px;  
  background-color: white;
  list-style: none;
  font-size: 14px;
  color: #71777c;

  .active{
    color: #1e80ff;
  }
}

</style>