<!-- 上方主导航栏 -->

<template>
  <div class="bg-container">
    <div class="nav-container">
      <!-- 主导航栏上半部分 - 页面跳转 -->
      <el-row type="flex" class="nav-bar">
        <img src="@/assets/logo-text.svg" alt class="logo-text" />
        <img src="@/assets/logo.svg" alt class="logo-img" />
        <!-- 主导航栏左半部分，包含链接 -->
        <el-col :span="9" class="left">
          <!-- 屏幕较窄时的下拉链接列表 -->
          <el-dropdown trigger="click" class="nav-link-dropdown">
            <span class="el-dropdown-link">
              首页
              <i class="el-icon-caret-bottom"></i>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>沸点</el-dropdown-item>
              <el-dropdown-item>直播</el-dropdown-item>
              <el-dropdown-item>活动</el-dropdown-item>
              <el-dropdown-item>商城</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>

          <!-- 完整的导航链接 -->
          <el-menu default-active="1" class="nav-link" mode="horizontal">
            <el-menu-item index="1"><el-link>首页</el-link></el-menu-item>
            <el-menu-item index="2"><el-link href="https://juejin.cn/pins" target="_blank">沸点</el-link></el-menu-item>
            <el-menu-item index="3"><el-link href="https://juejin.cn/course" target="_blank">课程</el-link></el-menu-item>
            <el-menu-item index="4"><el-link href="https://juejin.cn/live" target="_blank">直播</el-link></el-menu-item>
            <el-menu-item index="5"><el-link href="https://juejin.cn/events/all" target="_blank">活动</el-link></el-menu-item>
            <el-menu-item index="6"><el-link href="https://element.eleme.io" target="_blank">商城</el-link></el-menu-item>
          </el-menu>
        </el-col>

        <!-- 主导航栏右半部分，包含搜索框、创作者中心按钮、消息提示和用户头像 -->
        <el-col :span="15" class="right">
          <!-- 输入框 -->
          <el-input placeholder="探索掘金社区" class="input">
            <i slot="suffix" class="el-icon-search"></i>
          </el-input>

          <!-- 创作者中心按钮 -->
          <el-badge :value="1" class="badge" type="danger">
            <el-dropdown split-button type="primary" size="small" class="originator-drop">
              创作者中心
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item> <img src="@/assets/write-article.svg" /> 写文章 </el-dropdown-item>
                <el-dropdown-item> <img src="@/assets/boiling.svg" /> 发沸点 </el-dropdown-item>
                <el-dropdown-item> <img src="@/assets/coding.svg" /> 写代码 </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-badge>

          <!-- vip 图标 -->
          <img src="@/assets/logo-vip.svg" />

          <!-- 消息图标 -->
          <el-button type="text" icon="el-icon-bell" class="message-btn"></el-button>

          <!-- 用户头像 of 登录按钮 -->
          <div>
            <el-avatar v-if="isLogin">user</el-avatar>
            <el-button v-if="!isLogin" @click="openLoginDialog">登录</el-button>
          </div>
        </el-col>
      </el-row>      
    </div>
  </div>
</template>

<script>
import {mapState, mapMutations} from 'vuex'
export default {
  name: 'Navbar',
  computed: {
    ...mapState('loginOptions', ['isLogin', 'isLoginDialogShow', 'userInfo'])
  },
  methods: {
    // 点击导航栏的登录按钮，弹出登录对话框
    ...mapMutations('loginOptions', {openLoginDialog: 'OPEN_LOGIN_DIALOG'})
  }
}
</script>

<style scoped lang="less">
// 最外边的包装 div，用来控制导航栏的左右居中整体宽度
.bg-container {
  width: 100%;
  background-color: #fff;

  .nav-container {
    margin: 0 auto;
    width: 100%;
    max-width: 1440px;
    .nav-bar {
      height: 64px;
      align-items: center;
      background-color: white;

      // logo 区域
      .logo-text {
        width: 107px;
        height: 33px;
      }
      .logo-img {
        display: none;
        width: 31px;
      }
      // 媒体查询要写在下面才能生效（不知道为何，实验出来的）
      @media (max-width: 640px) {
        .logo-text {
          display: none;
        }
        .logo-img {
          display: block;
        }
      }
      // 导航栏区域
      .left,
      .right {
        display: flex;
      }
      // 左侧链接部分
      .left {
        justify-content: flex-start;
        align-items: center;
        margin-left: 10px;

        // 导航栏左侧宽度较窄时的下拉导航链接
        .nav-link-dropdown {
          font-size: 16px;
          color: #1e80ff;
          display: none;
          cursor: pointer;
        }

        // 完整的导航链接部分
        .nav-link {
          border-bottom: 1px solid white;
          .el-menu-item {
            width: 50px;
            font-size: 16px;
            padding: 0 10px;
            transition: none !important;
            &.is-active {
              color: #1e80ff;
              border-bottom: none;
            }
            &:hover {
              border-bottom: 1px solid #1e80ff !important;
            }
          }
        }

        @media (max-width: 1150px) {
          .nav-link-dropdown {
            display: block;
          }
          .nav-link {
            display: none;
          }
        }

        // border: 1px solid red;
      }

      // 右侧按钮、用户头像部分
      .right {
        justify-content: flex-end;
        align-items: center;
        margin: 0 5px;

        // 搜索框
        .input {
          max-width: 270px;
          min-width: 140px;
          margin-right: 12px;
          // 搜索框的图标
          .el-icon-search {
            max-width: 250px;
            margin: 3px;
            margin-right: 0;
            width: 30px;
            line-height: 32px;
            background-color: #f2f3f5;
          }
        }
        @media (max-width: 420px) {
          .input {
            display: none;
          }
        }

        // 创作者中心下拉按钮和消息气泡
        .badge {
          right: 72px;
          .originator-drop {
            right: -62px;
            min-width: 140px;
            margin: 0 12px;
            & > * {
              width: 100%;
            }
            &:hover{
              z-index: 0;
            }
          }
        }
        @media (max-width: 799px) {
          .badge {
            display: none;
          }
        }

        // 消息按钮
        .message-btn {
          height: 100%;
          margin: 0 20px;
        }

        // border: 1px solid skyblue;
      }
    }
  }
}
</style>
