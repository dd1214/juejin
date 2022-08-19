<!-- 内容区的每一项 -->
<!-- 数据有：发稿人、发稿时间、分类、文章标题、文章内容前 x 个字、点赞、评论、查看 -->
<template>
  <div class="container">
    <!-- 第一行：稿件信息 -->
    <div class="info">

    </div>
    <!-- 第二行：稿件内容预览 -->
    <div class="content">

    </div>
    <!-- 第三行：查看、点赞、评论 -->
    <div class="feedback">

    </div>
  </div>
</template>

<script>
export default {
  name: 'Essay',
  props: ['EssayTitle'],
  data() {
    return {}
  },
}
</script>

<style scoped lang="less">
// 最外部容器，控制位置
.container {
  .info{

    font-size: 13px;
    color: #4e5969;
  }
}
</style>
