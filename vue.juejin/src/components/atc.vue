<template>
  <div class="container">
    <div class="suspended-panel">
      
    </div>
    <div class="suspended-panel">
      
    </div>
    <div class="atc-main">
      <div class="article-title">前端设计模式|青训营笔记</div>
      <div class="author-info flex">
        <el-avatar>User</el-avatar>
        <div class="author-info-box">
          <div class="author-name">李泽宇</div>
          <div class="meta-box">
            <time>2022年08月26日 22:41</time><span>·&nbsp;&nbsp;阅读 1111</span>
          </div>
        </div>
        <button class="follow-button">+ 关注</button>
      </div>
      <div class="markdown-body">
        <h2>一、什么是设计模式？</h2>

        <p>设计模式（Design
          pattern）代表了最佳的实践，通常被有经验的面向对象的软件开发人员所采用。设计模式是软件开发人员在软件开发过程中面临的一般问题的解决方案。这些解决方案是众多软件开发人员经过相当长的一段时间的试验和错误总结出来的。
        </p>

        <p>设计模式是一套被反复使用的、多数人知晓的、经过分类编目的、代码设计经验的总结。使用设计模式是为了重用代码、让代码更容易被他人理解、保证代码可靠性。
          毫无疑问，设计模式于己于他人于系统都是多赢的，设计模式使代码编制真正工程化，设计模式是软件工程的基石，如同大厦的一块块砖石一样。项目中合理地运用设计模式可以完美地解决很多问题，每种模式在现实中都有相应的原理来与之对应，每种模式都描述了一个在我们周围不断重复发生的问题，以及该问题的核心解决方案，这也是设计模式能被广泛应用的原因。
        </p>
        <h2>二、为什么要学设计模式？</h2>
        <p>* 写出优雅的代码</p>
        <p>* 写出优雅的代码</p>
        <p>* 写出优雅的代码</p>
        <h2>三、设计模式的类型</h2>
        <h3>创建型：如何创建一个对象</h3>
        <h3>结构型：如何灵活地将对象组装成较大的结构</h3>
        <h3>行为型：负责对象间的高效通信和职责划分</h3>
        <img
          src="https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/47116068a681424492ccfa3b9dd94fec~tplv-k3u1fbpfcp-watermark.image"
          style="width:100%">
        <h2>四、前端涉及到的设计模型</h2>
        <h3>浏览器中的设计模型</h3>
        >**单例模式：** 定义一个全局唯一访问对象
        <br>应用场景：缓存、全局状态管理等。
        >**发布订阅模式：** 一种订阅机制，可在被订阅对象发生变化时通知订阅者
        <br>应用场景：从系统架构之间的解耦，到业务中一些现实模式，像邮件订阅，上线订阅等等，应用广泛。
        <h3>JavaScript中的设计模式：</h3>
        >**原型模式：** 复制已有对象来创建新的对象<br />
        应用场景：JS中对象创建的基本模式<br>
        >**代理模式：** 可自定义控制对原对象的访问方式，并允许在更新前后做一些额外处理<br />
        应用场景：监控，代理工具，前端框架实现等等。<br />
        >**迭代器模式：** 在不暴露数据类型的情况下访问集合中的数据<br />
        应用场景：数据结构有多种数据类型：列表、树等，提供通用操作接口。<br />
        <h3>前端框架中的设计模式：</h3>
        >**代理模式：** 可自定义控制对原对象的访问方式，并允许在更新前后做一些额外处理<br/>
        应用场景：监控，代理工具，前端框架实现等等。- Web 技术栈<br/>


        >**组合模式：** 可多个对象组合使用，也可单个对象独立使用。<br/>
        应用场景：DOM，前端组件，文件目录，部门。<br/>
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
        <div class="mulu" style="font-size: 18px; font-weight: bold; padding: 16px">
          目录
        </div>
        <hr style="opacity: 25%" />
        <div class="catalog">
          <el-tabs @tab-click="handleClick" v-model="activeName" :tab-position="tabPosition" style="height: auto">
            <el-tab-pane :name="'tab' + index" :class="item.lev" v-for="(item, index) in navList" :key="index"
              :label="item.name"></el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
    <el-backtop :bottom="100" :visibility-height="50"></el-backtop>
    <el-backtop :bottom="100" :visibility-height="50"></el-backtop>
  </div>


</template>
<script>
// import Like from './like.vue'

export default {
  name: 'Atc',
  components: {
    // Like
  },
  name: 'Atc',
  components: {
    // Like
  },
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
  created() {  },
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

.markdown-body {
  word-break: break-word;
  line-height: 1.75;
  font-weight: 400;
  font-size: 16px;
  overflow-x: hidden;
  color: #333;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  line-height: 5;
}

.suspended-panel {
  z-index: 2;
  width: 32px;
  height: auto;
  position: fixed;
  top: 140px;
  left: 40px;
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

@media screen and (max-width: 1140px) {
  .suspended-panel {
    display: none;
  }
}
</style>
