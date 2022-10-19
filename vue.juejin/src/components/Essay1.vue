<!-- 内容区的每一项 -->
<!-- 数据有：发稿人、发稿时间、分类、文章标题、文章内容前 x 个字、点赞、评论、查看 -->
<template>
  <div class="e-container">
    <!-- 稿件信息 -->
    <div class="info">
      <div>
        <span>{{ essay.info.author }}</span>
      </div>
      <div>{{ dateFromNow }}</div>
      <div>
        <span v-for="(item, index) in essay.info.category" :key="index">{{
          item
        }}</span>
      </div>
    </div>
    <!-- 稿件内容预览 -->
    <div class="content">
      <div class="text">
        <div class="title">{{ essay.content.title }}</div>
        <div class="preview">{{ essay.content.preview }}</div>
        <!-- 查看、点赞、评论 -->
        <div class="feedback">
          <ul>
            <li>
              <img src="@/assets/view.png" />
              <span>{{ feedbackProcess(essay.feedback.view) }}</span>
            </li>
            <li>
              <img src="@/assets/like.png" />
              <span>{{ feedbackProcess(essay.feedback.like) }}</span>
            </li>
            <li>
              <img src="@/assets/comments.png" />
              <span>{{ feedbackProcess(essay.feedback.comments) }}</span>
            </li>
          </ul>
        </div>
      </div>
      <!-- 可能存在的快照 -->
      <img v-if="essay.snapshot" :src="essay.snapshot" class="snapshot" />
    </div>
  </div>
</template>

<script>
import moment from 'moment'
moment.locale('zh-cn')
export default {
  name: 'Essay',
  data() {
    return {}
  },
  props: {
    // 文章
    essay: {
      // 文章第一行
      info: {
        author: String,
        date: String,
        category: Array,
      },
      // 文章内容区域
      content: {
        title: String,
        preview: String,
      },
      // 点赞，查看，评论
      feedback: {
        view: Number,
        like: Number,
        comments: Number,
      },
      // 图片快照（不一定有）
      snapshot: String,
    },
    index: Number
  },
  methods: {
    // 大于 10000 就切换显示（3.4w）
    feedbackProcess(value) {
      if (value >= 10000) {
        return `${(value / 1000).toFixed(0) / 10}w`
      } else {
        return value
      }
    },
  },
  computed: {
    dateFromNow() {
      return moment(this.essay.info.date, moment.ISO_8601).fromNow()
    },
  },
  mounted(){
    console.log(this.index)
  }
}
</script>

<style scoped lang="less">
// 最外部容器，控制位置
.e-container {
  // 网格布局
  display: grid;
  // 三行：24px 自动 20px
  grid-template-rows: 24px 80px 20px;
  // 两列：480px 20px
  grid-template-columns: 1fr 120px;
  // 先行后列排布
  grid-auto-flow: row;
  // 整体网格在水平方向平均分布
  justify-content: start;
  // // 网格内的内容水平方向从头开始
  // justify-items: start;

  padding: 5px 10px;
  max-width: 700px;

  background-color: white;
  border-bottom: 1px solid #e1e1e1;
  box-shadow: rgb(0 0 0 / 0%) 0px 2px 1px -1px, rgb(0 0 0 / 1%) 0px 1px 1px 0px,
    rgb(0 0 0 / 12%) 0px 1px 3px 0px;

  &:hover {
    background-color: #ededed;
    cursor: pointer;
  }

  // 信息区
  .info {
    grid-column-start: span 2;

    display: flex;
    justify-content: flex-start;
    align-items: center;

    font-size: 13px;
    color: #869093;
    white-space: nowrap;
    & > div:first-child {
      color: #4e5969;
    }

    div {
      margin: 7px 0;
      padding: 0 7px;
      border-right: 1px solid rgb(163, 163, 148);
      span:hover {
        color: #1e80ff;
      }
      // 中间的有点连接，最后一个没有点
      span::after {
        content: '·';
        padding: 0 5px;
      }
      & > span:last-child {
        &::after {
          content: '';
          padding: 0;
        }
      }
    }
    & > div:nth-child(3) {
      border-right: none;
    }

    // border: 1px solid royalblue;
  }
  // 文章内容区
  .content {
    grid-column-start: span 2;
    grid-row-start: span 1;

    display: flex;
    justify-content: space-between;
    align-content: center;
    padding: 10px;
    .text {
      .title {
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 1;
        text-overflow: ellipsis;
        overflow: hidden;
        margin-bottom: 5px;
        font-size: 16px;
        font-weight: bold;
        color: #1d2129;
      }
      .preview {
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 1;
        text-overflow: ellipsis;
        margin: 10px 0;
        overflow: hidden;
        color: #80969c;
        font-size: 13px;
      }
      // 用户反馈区
      .feedback {
        grid-row-start: 3;
        grid-column-start: 1;

        // padding: 0 10px;
        height: 100%;
        width: 100%;

        ul {
          display: flex;
          justify-content: flex-start;
          align-items: center;
          list-style: none;
          li {
            display: flex;
            justify-content: flex-start;
            align-items: center;
            color: #4e5969;
            img {
              width: 20px;
              height: 20px;
            }
            span {
              padding-left: 5px;
              padding-right: 10px;
              height: 16px;
              line-height: 16px;
            }
          }
        }

        // border: 1px solid #55efc4;
      }
    }
    // 缩略图区
    .snapshot {
      grid-column-start: 2;
      grid-row-start: 2;
      grid-row-end: 4;

      width: 120px;
      height: 80px;
      margin-left: 10px;
      object-fit: cover;

      // background-size: contain;

      // border: 1px solid green;
    }
  }
}
@media screen and (max-width: 1000px) {
  .e-container {
    width: 100%;
    max-width: 1000px;
    grid-template-columns: 1fr;
  }
}
</style>
