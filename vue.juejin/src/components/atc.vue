<template>
  <div class="container">
    <el-main>
      <div class="article-title">标题</div>
      <div class="author-info flex">
        <el-avatar>User</el-avatar>
        <div class="author-info-box">
          <div class="author-name">仙贝</div>
          <div class="meta-box">
            <time></time><span>·&nbsp;&nbsp;阅读 114514</span>
          </div>
        </div>
        <button class="follow-button">+ 关注</button>
      </div>
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
      <h1 id="-">一级目录</h1>
      <br />
      <br />
      <br />
      <br />

    </el-main>
    <!-- 作者信息 -->
    <div class="author">

    </div>
    <!--  掘金app二维码-->
    <div class="app-link">

    </div>
    <!-- 广告 -->
    <el-aside>
      <div style=" font-size: 18px; font-weight: bold;padding: 16px;">
        目录
      </div>
      <hr />
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
    </el-aside>
  </div>
</template>
<script>
import Adv from './adv.vue';


export default {
    data() {
        return {
            activeName: "tab0",
            tabPosition: "right",
            scroll: "",
            navList: [],
        };
    },
    methods: {
        handleClick(tab, event) {
            this.jump(tab.index);
        },
        dataScroll: function () {
            this.scroll =
                document.documentElement.scrollTop || document.body.scrollTop;
        },
        jump(index) {
            let jump = document.querySelectorAll("h1,h2,h3,h4,h5,h6");
            // 获取需要滚动的距离
            let total = jump[index].offsetTop - 80;
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
        loadScroll: function () {
            let self = this;
            let navs = document.querySelectorAll(".el-tabs__item");
            // var sections = document.getElementsByClassName('section');
            for (var i = self.navList.length - 1; i >= 0; i--) {
                if (self.scroll >= self.navList[i].offsetTop - 120) {
                    self.activeName = "tab" + i;
                    break;
                }
            }
        },
        selectAllTitle() {
            let title = document.querySelectorAll("h1,h2,h3,h4,h5,h6");
            this.navList = Array.from(title);
            this.navList.forEach((item) => {
                item.name = item.innerHTML;
            });
            this.navList.forEach((el) => {
                let index = el.localName.indexOf("h");
                el.lev = "lev" + el.localName.substring(index + 1, el.localName.length);
            });
        },
    },
    watch: {
        scroll: function () {
            this.loadScroll();
        },
    },
    created() { },
    mounted() {
        // scroll代表滚动条距离页面顶部距离
        window.addEventListener("scroll", this.dataScroll);
        this.selectAllTitle();
        this.$nextTick(() => {
            setTimeout(() => {
                let navs = document.querySelectorAll(".el-tabs__item");
                for (let i = navs.length - 1; i >= 0; i--) {
                    // console.log($('#'+navs[i].id))
                    // 从lev1到lev5分别添加不同到样式
                    document.querySelector("#" + navs[i].id).style.padding = "0";
                    if (this.navList[i].lev == "lev1") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "20px";
                    }
                    else if (this.navList[i].lev == "lev2") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "35px";
                    }
                    else if (this.navList[i].lev == "lev3") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "50px";
                    }
                    else if (this.navList[i].lev == "lev4") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "65px";
                        document.querySelector("#" + navs[i].id).style.fontWeight = "400";
                    }
                    else if (this.navList[i].lev == "lev5") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "80px";
                        document.querySelector("#" + navs[i].id).style.fontWeight = "400";
                    }
                }
            });
        });
    },
    components: { Adv }
};
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
  margin: 0 0 0 auto;
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
.follow-button:hover{
    background: rgb(30, 128, 255, 0.15);
    transition: 300ms;
}
.el-main {
  width: 60%;
  margin: auto;
  background-color: white;
}

.el-tabs__header.is-right {
  height: auto;
  width: 299px !important;

}

.el-aside {
  width:300px !important;
  right: 200px;
  margin: 20px;
  float: right;
  position: fixed;
  top: 108px;
  height: auto;
  background-color: white;
  border: 20px;
  border-radius: 5px;
}
.container {
  width: 80%;
  padding-top: 20px;
}


</style>
