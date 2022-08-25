<template>
  <div class="container">
    <div class="atc-main">
      <div class="article-title">标题</div>
      <div class="author-info flex">
        <el-avatar>User</el-avatar>
        <div class="author-info-box">
          <div class="author-name">用户名</div>
          <div class="meta-box">
            <time>2022年08月22日 1:23</time
            ><span>·&nbsp;&nbsp;阅读 114514</span>
          </div>
        </div>
        <button class="follow-button">+ 关注</button>
      </div>
      <div style="padding-top:10px">
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h2 id="-1">二级目录1</h2>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h3 id="-1">三级目录1</h3>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      </div>
      <!-- <div class="comment-box">评论
        <div class="comment">

        </div>
      </div> -->
    </div>

    <div class="aside">
      <!-- 作者信息 -->
      <div class="author"></div>
      <!--  掘金app二维码-->
      <div class="app-link">
        <img src="../assets/QRCode.jpg" style="max-width: 21%" />
        <div class="app-card">
          <div class="app-card-download">下载稀土掘金APP</div>
          <div class="app-card-text">一个帮助开发者成长的社区</div>
        </div>
      </div>
      <!-- 广告 -->
      <!-- 目录 -->
      <div class="catalog-box">
        <div
          class="mulu"
          style="font-size: 18px; font-weight: bold; padding: 16px"
        >
          目录
        </div>
        <hr style="opacity: 25%" />
        <div class="catalog">
          <el-tabs
            @tab-click="handleClick"
            v-model="activeName"
            :tab-position="tabPosition"
            style="height: auto"
          >
            <el-tab-pane
              :name="'tab' + index"
              :class="item.lev"
              v-for="(item, index) in navList"
              :key="index"
              :label="item.name"
            ></el-tab-pane>
          </el-tabs>
        </div>
      </div>
      
    </div>
  </div>
</template>
<script>
import { Tabs } from "element-ui"

export default {
  data() {
    return {
      activeName: "tab0",
      tabPosition: "right",
      scroll: "",
      navList: [],
    }
  },
  methods: {
    handleClick(tab, event) {
      this.jump(tab.index)
    },
    dataScroll: function () {
      this.scroll =
        document.documentElement.scrollTop || document.body.scrollTop
    },
    jump(index) {
      let jump = document.querySelectorAll("h1,h2,h3,h4,h5,h6")
      // 获取需要滚动的距离
      let total = jump[index].offsetTop - 80
      // Chrome
      document.body.scrollTop = total
      // Firefox
      document.documentElement.scrollTop = total
      // Safari
      window.pageYOffset = total
      // $('html, body').animate({
      // 'scrollTop': total
      // }, 400);
    },
    loadScroll: function () {
      let self = this
      let navs = document.querySelectorAll(".el-tabs__item")
      // var sections = document.getElementsByClassName('section');
      for (var i = self.navList.length - 1; i >= 0; i--) {
        if (self.scroll >= self.navList[i].offsetTop - 120) {
          self.activeName = "tab" + i
          break
        }
      }
    },
    selectAllTitle() {
      //获取h1-6标题
      let title = document.querySelectorAll("h1,h2,h3,h4,h5,h6")
      this.navList = Array.from(title) //将获取的title存储到navList数组中
      this.navList.forEach((item) => {
        //遍历navList数组，将每个title存储为item.name
        item.name = item.innerHTML
      })
      this.navList.forEach((el) => {
        let index = el.localName.indexOf("h")
        el.lev = "lev" + el.localName.substring(index + 1, el.localName.length) //截取下标，获取title的级别
      })
    },
    logScrollHeight() {
      //获取catalog元素
      let catalog = document.getElementsByClassName("catalog")
      //获取catalog-box元素
      let catalog_box = document.getElementsByClassName("catalog-box") 
      //获取catalog距离浏览器顶部的高度
      let cTop = catalog[0].offsetTop
      // console.log(window.pageYOffset)
      //如果浏览器页面滚动高度大于目录距离页面顶部的高度，将外部box的position设置为fixed
      if (window.pageYOffset > cTop) {
        catalog_box[0].style.position = "fixed"
        //调整box的宽度
        catalog_box[0].style.width = "20%"
      } else if (window.pageYOffset < cTop) {
        // 当鼠标滚轮网上滚，页面滚动高度小于目录距离页面顶部的高度时，重置外部box的position跟width
        catalog_box[0].style.position = ""
        catalog_box[0].style.width = ""
      }
    },
  },
  watch: {
    scroll: function () {
      this.loadScroll()
    },
  },
  created() {},
  mounted() {
    // scroll代表滚动条距离页面顶部距离
    window.addEventListener("scroll", this.dataScroll)
    window.addEventListener("scroll", this.logScrollHeight)
    this.selectAllTitle()
    this.$nextTick(() => {
      setTimeout(() => {
        let navs = document.querySelectorAll(".el-tabs__item")
        for (let i = navs.length - 1; i >= 0; i--) {
          // console.log($('#'+navs[i].id))
          // 从lev1到lev5分别添加不同到样式
          document.querySelector("#" + navs[i].id).style.padding = "0"
          if (this.navList[i].lev == "lev1") {
            document.querySelector("#" + navs[i].id).style.paddingLeft = "20px"
          } else if (this.navList[i].lev == "lev2") {
            document.querySelector("#" + navs[i].id).style.paddingLeft = "35px"
          } else if (this.navList[i].lev == "lev3") {
            document.querySelector("#" + navs[i].id).style.paddingLeft = "50px"
          } else if (this.navList[i].lev == "lev4") {
            document.querySelector("#" + navs[i].id).style.paddingLeft = "65px"
            document.querySelector("#" + navs[i].id).style.fontWeight = "400"
          } else if (this.navList[i].lev == "lev5") {
            document.querySelector("#" + navs[i].id).style.paddingLeft = "80px"
            document.querySelector("#" + navs[i].id).style.fontWeight = "400"
          }
        }
      })
    })
  },
}
</script>
<style lang="scss">
* {
  font-weight: bold;
}

.article-title {
  font-size: 32px;
  margin-bottom: 20px;
}

.author-info-box {
  margin-bottom: 20px;
  align-items: center;
  cursor: pointer;
}
.el-avatar {
  margin-right: 12px;
  cursor: pointer;
}
.author-name {
  font-size: 16px;
  cursor: pointer;
}
.meta-box {
  margin-top: 2px;
  color: rgb(134, 134, 134);
}
.follow-button {
  margin: 15px 20px 15px auto;
  cursor: pointer;
  height: 34px;
  font-size: 14px;
  background: rgb(30, 128, 255, 0.05);
  border: 1px solid rgb(30, 128, 255, 0.3);
  border-radius: 4px;
  align-items: center;
  width: 52px;
  color: #1e80ff;
}
.follow-button:hover {
  background: rgb(30, 128, 255, 0.15);
  transition: 300ms;
}
.atc-main {
  width: 50%;
  margin: 0px 15px 200px 200px;
  background-color: white;
  padding: 25px 20px 20px 20px;
}

.el-tabs__header.is-right {
  height: auto;
  width: 100%;
}

.aside {
  width: 20%;
  right: 15%;
  top: 18.75rem;
  height: auto;
  border: 20px;
  border-radius: 5px;
}
.app-link {
  margin: 5px 0;
  cursor: pointer;
  display: inline-block;
  background-color: white;
  padding: 10px;
}
.app-card {
  display: inline-block;
}
.app-card-download {
  font-size: 16px;
  padding-bottom: 24px;
}
.app-card-text {
  color: #8a919f;
  border: 4px;
}
.mulu {
  background-color: white;
}
.catalog {
  max-height: 500px;
  background-color: white;
  overflow: hidden;
  // width: 20%;
  // position: fixed;
  border-radius: 5px;
}
// .catalog::-webkit-scrollbar {
//   width: 0;
//   height: 0;
//   color: transparent;
// }
.container {
  display: flex;
  width: 100%;
  height: 100%;
  padding-top: 20px;
  justify-content: center;
  align-items: flex-start;
}

@media screen and (max-width: 1000px) {
    .aside {
      display: none;
    }
  }

</style>
