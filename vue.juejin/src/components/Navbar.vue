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

          <!-- 用户menu -->
          <div>
            <!-- 登录后 -->

             <el-popover
              placement="bottom"
              trigger="click">
              
              <!-- 信息框 -->
            <div class="infobox">
              <div class="user-card" >
                <div class="user-info">
                  <div class="avatar">
                    <img src="@/assets/info/head protrait.jpg" alt="" class="lazy avatar" width="40px">
                  </div>
                  <div class="user-detail">
                    <p class="name">user-name</p>
                    <div class="ore">
                      矿石:
                      <span>2.8w</span>
                      <img src="@/assets/info/unfold.jpg" alt=""  class="path">
                    </div>
                  </div>
                </div>
                <a href="" class="progress-list">
                  <div class="jscore" style="background-color:#fff">
                    <div class="js-level">
                      掘友等级
                      <span>JY.4</span>
                    </div>
                    <div class="progress">
                      <span>
                        308.4/500
                        <img src="@/assets/info/unfold.jpg" alt="" class="progress-path">
                      </span>
                    </div>
                    <div class="bar">
                      <div class="current-bar" style="width:128px;"></div>
                    </div>
                  </div>
                </a>
                <ul class="status-bar">
                  <li class="item">
                    <div class="item-count">2</div>
                    <div class="item-name">关注</div>
                  </li>
                  <li class="item">
                    <div class="item-count">27</div>
                    <div class="item-name">赞过</div>
                  </li>
                  <li class="item">
                    <div class="item-count">1</div>
                    <div class="item-name">收藏</div>
                  </li>
                </ul>
              </div>
              <ul class="drop-down list box">
                <li class="drop-down" >
                  <div class="open-menu">
                    <img src="@/assets/info/homepage.jpg" alt="" ip="image">
                    <span @click.native="jumpTOPersonalPage(gerenzhuye)">我的主页</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/welfare.jpg" alt="" ip="image">
                    <span>成长福利</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/member.jpg" alt="" ip="image">
                    <span>会员中心</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/course.jpg" alt="" ip="image">
                    <span>我的课程</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/discounts.jpg" alt="" ip="image">
                    <span>我的优惠</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/apply.jpg" alt="" ip="image">
                    <span>我的报名</span>
                  </div>
                </li>
                <li class="drop-down">
                  <div class="open-menu">
                    <img src="@/assets/info/footprint.jpg" alt="" ip="image">
                    <span>我的足迹</span>
                  </div>
                </li>
              </ul>
              <div class="operation menu">
                <div class="setup">我的设置</div>
                  <a href="loginout" class="login-out" style="color:rgb(188, 187, 187);">退出登录</a>
              </div>
            </div>
              
              <div class="avatar-div"  style="width:40px; height:40px" slot="reference">
                <img src="@/assets/info/head protrait.jpg" style="width: 100%; height: 100%;" >
            </div>   
            </el-popover>

            
              
          </div>
        </el-col>
      </el-row>      
    </div>
  </div>
</template>

<script>

export default {
  name: 'Navbar',
  components: {
    LinkPop,
  },
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
