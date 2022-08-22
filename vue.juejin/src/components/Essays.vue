<!-- 全部文章页，主要做到无限向下滚动 -->
<template>
  <div class="es-container">
    <ul class="nav">
      <li @click="activate(1)" :class="{ active: activeIndex == 1 }">推荐</li>
      <li @click="activate(2)" :class="{ active: activeIndex == 2 }">最新</li>
      <li @click="activate(3)" :class="{ active: activeIndex == 3 }">热榜</li>
    </ul>
    <div class="content" ref="essaylist">
      <Essay v-for="(essay, index) in essays" :key="index" :essay="essay" @click.native="jump2Atc('xiangqingye')" />
    </div>
  </div>
</template>

<script>
import Essay from '@/components/Essay'

export default {
  name: 'Essays',
  components: {
    Essay,
  },
  data() {
    return {
      // nav 栏激活 index
      activeIndex: 1,
      // 判断当前是否正在加载（正忙）
      busy: false,
      // 记录当前文章列表长度
      curEssaysLength: 0,
      // IntersectionObserver 对象，用来监听元素实现无限滚动
      observer: new IntersectionObserver(
        (entries) => {
          if (!this.busy && entries[0].intersectionRatio > 0.75) {
            this.busy = true
            const p = new Promise((resolve, reject) => {
              this.loadMore()
              resolve()
            })
            p.then(() => {
              // 执行异步操作后的内容（目前没有）
            })
          }
        },
        {
          threshold: [0.75],
        }
      ),

      // 初始文章信息
      //#region
      essays: [
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '掘金安东尼',
            date: '2022-08-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
      ],
      //#endregion

      // 测试无限滚动，用来新增的新文章数据
      newEssays: [
        {
          info: {
            author: '湖人科比',
            date: '2022-03-10',
            category: ['前端', '面试', '瞎玩'],
          },
          content: {
            title: '如何打破 2-3 联防',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“投篮”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: '大鲨鱼',
            date: '2022-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: ' 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
        {
          info: {
            author: '骑士詹姆斯',
            date: '2022-09-10',
            category: ['前端', '面试'],
          },
          content: {
            title: 'JavaScript 中如何取消请求',
            preview: '本篇带来 XMLHttpRequest、Fetch 和 axios 分别是怎样“取消请求”的。闲话少说，冲就完事了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',
          },
          feedback: {
            view: 9364,
            like: 527,
            comments: 337,
          },
        },
        {
          info: {
            author: 'Big shark@LX',
            date: '2021-11-27',
            category: ['前端', 'JavaScript', 'TypeScript'],
          },
          content: {
            title: '最全的 TypeScript 学习指南',
            preview: '前言 Hello 大家好 我是鲨鱼哥 这次给大家带来的是我曾经非常嫌弃 如今却爱不释手的 TS 技术 哈哈',
          },
          feedback: {
            view: 61407,
            like: 1428,
            comments: 87,
          },
          snapshot: 'https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e15e04fd93f4eb18d310e19526ce181~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?',
        },
      ],

      //
    }
  },
  computed: {
    availableEssays() {
      if (this.essays.length <= 24) {
        return this.essays
      }
      // 当加载的文章过多时会有明显卡顿，这种情况下只拿出一部分来渲染
    },
  },
  methods: {
    activate(index) {
      this.activeIndex = index
    },
    // 模拟滚动到最下方时加载新数据
    loadMore() {
      setTimeout(() => {
        this.essays = [...this.essays, ...this.newEssays]
      }, 1000)
    },
    // 点击文章跳转到详情页
    jump2Atc(pageName) {
      // 路由跳转之前校验是否重复跳转到当前路由，是的话则不跳转
      if (this.$route.name != pageName) {
        let newRoute = this.$router.resolve({
          name: pageName,
        })
        console.log(newRoute)
        console.log(window.location.origin)
        window.open(newRoute.href, '_blank')
      }
      // window.open('localhost:8080/src/details', '_blank')
    },

    // getEssays(){
    //   // 发送 ajax 请求获取 essays
    // }
  },
  mounted() {
    // 监听当前列表的最下部元素
    this.observer.observe(this.$refs.essaylist.lastElementChild)
  },
  updated() {
    // 更新后，如果使得文章列表长度发生变化，则重新监听最下方的元素
    if (this.$refs.essaylist.childNodes.length != this.curEssaysLength) {
      this.observer.disconnect()
      this.observer.observe(this.$refs.essaylist.lastElementChild)
      this.busy = false
    }
  },
  beforeDestroy() {
    this.observer.unobserve(this.$refs.essaylist.lastElementChild)
  },
}
</script>

<style scoped lang="less">
.es-container {
  margin-right: 15px;
  .nav {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    list-style: none;

    padding: 20px;

    background-color: white;
    color: #909090;
    font-size: 14px;
    white-space: nowrap;
    li {
      padding: 0 15px;
      &:hover {
        color: #1e80ff;
        cursor: pointer;
      }
      &.active {
        color: #1e80ff;
      }
    }
  }
  content {
    width: 100%;
  }
}
@media screen and (max-width: 1000px) {
  .es-container {
    width: 100%;
    margin-right: 0;
  }
}
</style>
