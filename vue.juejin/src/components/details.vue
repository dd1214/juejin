<template>
  <div class="container">
    <div class="wrapper">
      <div
        class="section"
        style="height: 500px; width: 100%"
        v-for="(item, index) in list"
        :key="index"
        :style="{ height: index == 0 ? '1000px' : '500px' }"
      >
        <div
          style="
            width: 100%;
            height: 100%;
            font-size: 30px;
            text-align: center;
            font-weight: bold;
            color: #fff;
          "
          :style="{ background: item.backgroundcolor }"
        >
          {{ item.name }}
        </div>
      </div>
    </div>
    <div id="nac" style="height: 500px"></div>
    <nav style="position: fixed; right: 30px; top: 300px">
      <div style="margin-left: 20px; font-size: 16px">目录</div>
      <el-tabs
        @tab-click="handleClick"
        v-model="activeName"
        :tab-position="tabPosition"
        style="height: 200px"
      >
        <el-tab-pane
          :name="'tab' + index"
          :class="index == 0 ? 'current' : ''"
          v-for="(item, index) in list"
          :key="index"
          :label="item.name"
        ></el-tab-pane>
      </el-tabs>
    </nav>
  </div>
</template>
<script>
export default {
  data() {
    return {
      activeName: "tab0",
      tabPosition: "right",
      scroll: "",
      list: [
        {
          name: "第一条",
          backgroundcolor: "#90B2A3",
        },
        {
          name: "第二条",
          backgroundcolor: "#A593B2",
        },
        {
          name: "第三条",
          backgroundcolor: "#A7B293",
        },
        {
          name: "第四条",
          backgroundcolor: "#0F2798",
        },
        {
          name: "第五条",
          backgroundcolor: "#0A464D",
        },
      ],
      navList: [1, 2, 3, 4, 5],
    };
  },
  methods: {
    //这里传入的tab相当于item
    handleClick(tab, event) {
      console.log(tab.index);
      this.jump(tab.index);
    },
    dataScroll: function () {
      this.scroll =
        document.documentElement.scrollTop || document.body.scrollTop;
    },
    jump(index) {
      let jump = document.getElementsByClassName("section");
      // 获取需要滚动的距离
      let total = jump[index].offsetTop;
      // Chrome
      document.body.scrollTop = total;
      // Firefox
      document.documentElement.scrollTop = total;
      // Safari
      window.pageYOffset = total;
      // $('html, body').animate({
      // 'scrollTop': total
      // }, 400);
    },
    // 用循环的方式将每个标题离顶部的距离与滚动条当前位置对比来确定哪个标题需要变为高亮
    loadScroll: function () {
      var self = this;
      var sections = document.getElementsByClassName("section");
      for (var i = sections.length - 1; i >= 0; i--) {
        if (self.scroll >= sections[i].offsetTop - 100) {
          //我在上面规定了每个el-tab-pane标签的name属性值为tab+该标签的index值
          self.activeName = "tab" + i;
          break;
        }
      }
    },
  },
  watch: {
    //监听scroll变量，只要滚动条位置变化就会执行方法loadScroll
    scroll: function () {
      this.loadScroll();
    },
  },
  mounted() {
    // scroll代表滚动条距离页面顶部距离
    window.addEventListener("scroll", this.dataScroll);
  },
};
</script>

<style lang="scss" scoped>
* {
  padding: 0;
  margin: 0;
}
.nav1 {
  display: block;
  width: 120px;
  height: 40px;
  text-align: left;
  line-height: 40px;
  color: #fff;
  /* background: #eee; */
  margin: 10px 0;
  cursor: pointer;
  padding-left: 18px;
  &:hover {
    color: #0177ff;
  }
}
.current {
  color: #0578fc !important;
  cursor: pointer;
}
nav {
  // border-left: 1px solid #eee;
  a {
    border-left: 3px solid #0177ff;
  }
}
</style>
<style>
.el-tabs__header.is-right {
  margin-right: 100px !important;
}
</style>
