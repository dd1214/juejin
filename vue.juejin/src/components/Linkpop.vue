<!-- 弹出框，被用在主导航栏的下半部分的弹出显示 -->
s
<template>
  <div :class="{ container: true, isFixed: isFixed }" ref="linkpop">
    <ul>
      <LinkpopItem
        v-for="item in lists"
        :key="item.id"
        :index="item.id"
        :info="item.info"
      />
    </ul>
  </div>
</template>

<script>
import LinkpopItem from '@/components/LinkpopItem'
export default {
  name: 'Linkpop',
  components: {
    LinkpopItem,
  },
  data() {
    return {
      active_index: 1,
      isFixed: false,
      // 测试数据
      lists: [
        {
          id: 1,
          info: {
            main: '综合',
          },
        },
        {
          id: 2,
          info: {
            main: '关注',
          },
        },
        {
          id: 3,
          info: {
            main: '后端',
            sub: ['后端', 'java', 'go'],
          },
        },
        {
          id: 4,
          info: {
            main: '前端',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 5,
          info: {
            main: 'Android',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 6,
          info: {
            main: 'iOS',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 7,
          info: {
            main: '人工智能',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 8,
          info: {
            main: '开发工具',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 9,
          info: {
            main: '代码人生',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
        {
          id: 10,
          info: {
            main: '阅读',
            sub: ['前端', 'JavaScript', 'Vue.js', 'React.js', 'css'],
          },
        },
      ],
    }
  },
  methods: {
    handleScroll() {
      // 最上方下半部导航吸附在顶部
      const l = this.$refs.linkpop
      let scrollTop = document.documentElement.scrollTop
      this.isFixed = scrollTop >= 64 ? true : false
    },
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
}
</script>

<style scoped lang="less">
.container {
  display: flex;
  justify-content: flex-start;
  border-top: 1px solid #e6e6e6;
  width: 100%;
  background-color: white;
  padding-left: 100px;
  ul {
    display: flex;
    justify-content: flex-start;
    align-items: stretch;
    width: fit-content;
    height: 46px;

    padding: 10px 20px;
    background-color: white;
    list-style: none;
    font-size: 14px;
    color: #71777c;

    .active {
      color: #1e80ff;
    }
  }
}
@media screen and (max-width: 960px) {
  .container {
    padding-left: 0 !important;
  }
}
@media screen and (max-width: 1100px) {
  .container {
    padding-left: 20px;
  }
}
@media screen and (max-width: 1200px) {
  .container {
    padding-left: 50px;
  }
}
.isFixed {
  position: fixed;
  top: 0;
  z-index: 9999;
}
</style>
