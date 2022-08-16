<template>
  <el-row type="flex" class="nav-bar">
    <img src="@/assets/logo-text.svg" alt="" class="logo-text" />
    <img src="@/assets/logo.svg" alt="" class="logo-img" />
    <!-- 主导航栏左半部分，包含链接 -->
    <el-col :span="12" class="left">
      <!-- 屏幕较窄时的下拉链接列表 -->
      <el-dropdown trigger="click" class="nav-link-dropdown">
        <span class="el-dropdown-link">
          首页<i class="el-icon-caret-bottom"></i>
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
        <el-menu-item index="1">首页</el-menu-item>
        <el-menu-item index="2">沸点</el-menu-item>
        <el-menu-item index="3">直播</el-menu-item>
        <el-menu-item index="4">活动</el-menu-item>
        <el-menu-item index="5">商城</el-menu-item>
        <el-menu-item index="6">APP</el-menu-item>
        <el-menu-item index="7">插件</el-menu-item>
      </el-menu>
    </el-col>

    <!-- 主导航栏右半部分，包含搜索框、创作者中心按钮、消息提示和用户头像 -->
    <el-col :span="12" class="right">
      <!-- 输入框 --> 
      <el-input placeholder="探索掘金社区" class="input">
        <i slot="suffix" class="el-icon-search"></i>
      </el-input>   

      <!-- 创作者中心按钮 -->
      <el-badge :value="1" class="badge" type="danger">
        <el-dropdown split-button type="primary" size="small" class="originator-drop">
          创作者中心
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item><img src="@/assets/write-article.svg" /> 写文章</el-dropdown-item>
            <el-dropdown-item><img src="@/assets/boiling.svg" /> 发沸点</el-dropdown-item>
            <el-dropdown-item><img src="@/assets/coding.svg" /> 写代码</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-badge>

      <!-- vip 图标 -->
      <img src="@/assets/logo-vip.svg">

      <!-- 消息图标 -->
      <el-button type="text" icon="el-icon-bell" class="message-btn"></el-button>

      <!-- 用户头像 -->
      <div>
        <el-avatar> user </el-avatar>
      </div>
    </el-col>
  </el-row>


  
</template>

<script>
export default {
  name: "MainNav",
};
</script>

<style scoped lang="less">
.nav-bar {
  height: 63px;
  align-items: center;

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
      margin-left: 20px;
      font-size: 16px;
      color: #1e80ff;
      display: none;
      cursor: pointer;
    }

    // 完整的导航链接部分
    .nav-link {
      
      display: block;

      .el-menu-item{
        width: 50px;
        font-size: 16px;
        padding: 0 10px;
        &.is-active{
          color: #1e80ff; 
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
    .input{
      max-width: 250px;
      // 搜索框的图标
      .el-icon-search{
        max-width: 250px;
        margin: 3px;
        margin-right: 0;
        width: 30px;
        line-height: 32px;
        background-color: #f2f3f5;
      }
    }
    

    // 创作者中心下拉按钮和消息气泡
    .badge{
      right: 72px;
      .originator-drop {
        right: -62px;
        min-width: 140px;
        margin: 0 12px;

        & > * {
          width: 100%;
        }
      }
    }
    
    // 消息按钮
    .message-btn{
        height: 100%;
        margin: 0 20px;
    }

    // border: 1px solid skyblue;
  }
}
</style>
