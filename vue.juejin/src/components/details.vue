<template>
  <div class="icon-sources-wrapper wrapper">
    <Navbar/>
    <el-main>
      <h1 id="0">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h2 id="1">二级目录</h2>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h3 id="1">三级目录</h3>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h1 id="0">一级目录</h1>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <h4 id="1">四级目录1</h4>
      <p>
        日光之下无新事，世间处处是围城。有的人死了他还活着，有的人活着他已经死了，方鸿渐30岁以后的生活，就符合后一种。
        事业上，方鸿渐自命不凡却一事无成，每份工作，都是在糊里糊涂中就遭了嫌弃。
        爱情上，他茫然被动，不喜欢苏文纨，却没勇气拒绝，倾心唐晓芙，却不会争取。
        朋友提醒他，孙柔嘉心机深，对他煞费苦心，他却觉得“承她瞧得起，应当珍惜她”。
        最后，方鸿渐落入孙柔嘉精心编织的网中，顺水推舟结了婚，种下了他们往后余生，日复一日争吵的因。
        有一句话说：“生命绚烂的人，都是受到某些力量所驱使。”
      </p>
      <h4 id="">四级目录2</h4>
      <p>
        而有些人，睡着和醒着，都是在无意义地游荡人间。
        不知所因，不知所终，任由人世洪流裹挟着前行，最终遭受到了生活的重击。
        没有梦，没有想法的一生，就无异于死的样品。
      </p>
    </el-main>
    <el-aside>
      <div style="margin:0 0 0 180px;font-size:18px;font-weight:bold">目录</div>
      <el-tabs @tab-click="handleClick" v-model="activeName" :tab-position="tabPosition" style="height: auto;">
        <el-tab-pane :name="'tab'+index"
        :class="item.lev"
        :v-for="(irem,index)in navList"
        :key="index"
        :label="item.name">
        </el-tab-pane>
      </el-tabs>
    </el-aside>
  </div>
</template>
<script>
import Navbar from './Navbar.vue';
  export default{
    data() {
        return {
            activeName: "tab0",
            tabPosition: "right",
            scroll: "",
            navList: [],
        };
    },
    methods: {
        handClick(tab, event) {
            this.jump(tab.index);
        },
        dataScroll: function () {
            this.scroll = document.documentElement.scrollTop || document.body.scrollTop;
        },
        jump(index) {
            //获取所有标题
            let jump = document.querySelectorAll("h1,h2,h3,h4,h5,h6");
            //获取需要滚动的距离
            let total = jump[index].offsetTop - 80;
            //兼容性测试
            //Chrome
            document.body.scrollTop = total;
            //firefox
            document.documentElement.scrollTop = total;
            //safari
            window.pageYOffset = total;
        },
        loadScroll: function () {
            let self = this;
            let navs = document.querySelectorAll(".el-tabs__item");
            for (var i = self.navList.lenth - 1; i >= 0; i--) {
                if (self.scroll >= self.navList[i].offsetTop - 120) {
                    self.activeName = "tab" + i;
                    break;
                }
            }
        },
        selectAllTitle() {
            let title = document.querySelectorAll("h1,h2,h3,h4,h5,h6");
            this.navList = Array.from(title);
            this.navList.forEach(item => {
                item.name = item.innerHTML;
            });
            this.navList.forEach(el => {
                let index = el.localName.indexOf("h");
                el.lev = "lev" + el.localName.substring(index + 1, el.localName.length);
            });
        }
    },
    watch: {
        scroll: function () {
            this.loadScroll();
        }
    },
    created() {
    },
    mounted() {
        //scroll表示滚动条距离页面顶部的距离
        window.addEventListener("scroll", this.dataScroll);
        this.selectAllTitle();
        this.$nextTick(() => {
            setTimeout(() => {
                let navs = document.querySelectorAll(".el-tabs__item");
                for (let i = navs.length - 1; i >= 0; i--) {
                    //从lev1到lev5分别添加不同的样式
                    document.querySelector("#" + navs[i].id).style.padding = "0";
                    if (this.navList[i].lev == "lev1") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "20px";
                    }
                    else if (this.navList[i].lev == "lev2") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "30px";
                    }
                    else if (this.navList[i].lev == "lev3") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "40px";
                    }
                    else if (this.navList[i].lev == "lev4") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "50px";
                        document.querySelector("#" + navs[i].id).style.fontWeight = "400";
                    }
                    else if (this.navList[i].lev == "lev5") {
                        document.querySelector("#" + navs[i].id).style.paddingLeft = "60px";
                        document.querySelector("#" + navs[i].id).style.fontWeight = "400";
                    }
                }
            });
        });
    },
    components: { Navbar, Navbar }
}
</script>
<style lang="scss">
.el-main{
  width: 900px;
}
.el-tabs__header.is-right{
  height: 500px !important;
}
.el-aside{
  position: fixed;
  top: 108px;
  right: 160px;
  width: 220px;
  height: auto;
}
.icon-sources-wrapper.wrapper .el-tabs__nav.is-right{
  box-sizing: content-box !important;
}
</style>
