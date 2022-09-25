<template>
<div class="personal-data">
            <div>
               <div class="itemscope">
                <img src="../assets/info/head protrait.jpg" alt="" width="100px" class="head-photo">
                    <h1 class="user-name">
                        {{username}}
                        <span>
                            <img src="../assets/info/l-grade.jpg" alt="" width="40px">
                        </span>
                        <span>
                            <img src="../assets/info/j-grade.jpg" alt="" class="j-level" width="40px">
                        </span>
                    </h1>
                    <div class="introduce">
                        <a href=""  class="text1" >+ 你从事什么职业？</a>
                        <br>
                        <a href="" class="text1" >+ 你的人生格言是什么？</a>
                    </div>
                    <div class="link">
                        <a href="" class="click">
                           <img src="../assets/info/weibo.jpg" alt="" width="30px">
                        </a>
                        <a href="" class="click">
                           <img src="../assets/info/github.jpg" alt="" width="26px">
                        </a>
                        <a href="" class="click">
                           <img src="../assets/info/address.jpg" alt="" width="23px">
                        </a>
                    </div>
                    <hr>
                    <div class="edit">
                        <button style="width: 120px;height:30px">编辑个人资料</button>
                    </div>

                    <div class="father">
                <a href="" class="text">
                <div class="badge">
                    <img src="../assets/info/badge.png" alt="" width="20px">
                </div>
                <div class="getbadge">获得徽章0</div>
                <div>
                    <ul class="badge-list">
                        <img src="../assets/info/unfold.jpg" alt="">
                    </ul>
                </div>
                </a>
                </div>
                </div>
        </div>
       
        <div class="navigation">
            <el-menu default-active="1" class="nav-link" mode="horizontal">
            <el-menu-item index="1"><el-link >文章</el-link></el-menu-item>
            <el-menu-item index="2"><el-link >专栏</el-link></el-menu-item>
            <el-menu-item index="3"><el-link >沸点</el-link></el-menu-item>
            <el-menu-item index="4"><el-link >收藏集</el-link></el-menu-item>
            <el-menu-item index="5"><el-link >关注</el-link></el-menu-item>
            <el-menu-item index="6"><el-link >赞</el-link></el-menu-item>
          </el-menu> 
           <div class="search">
           <img src="../assets/info/search.jpg" alt="" width="26px">
           </div>
        </div>
</div>
</template>

<script>
import { mapState } from 'vuex';
export default {
    name:"HomepageCenter",
    data(){
        return {
            username: localStorage.getItem('username') || 'username'
        }
    }
}
</script>

<style scoped>
        .personal-data
        {
            display:absolute;

        }
        .head-photo
        {
            margin-top: 39px;
            margin-left: 29px;
            display: inline-block;
        }
        .itemscope
        {
            width: 700px;
            height: 180px;
            background-color: #fff;
            display: inline-block;
            position: absolute;
            top: 80px;
            left: 200px;
        }

        .user-info
        {
            height: 50px;
            font-size: 20px;
            font-weight: 600;
            align-items: center;
            margin-block-start:10px;
        }
        .user-name
        {
            display:inline-block;
            float:right;
            margin-top:30px;
            margin-right:330px;
        }
        .introduce
        {
            position: absolute;
            top:70px;
            left:150px;
        }

        .text1
        {
            text-decoration: none;
            color: rgb(10, 10, 186);
        }
        .link
        {
            position: absolute;
            top: 40px;
            left: 570px;
        }
        .edit{
            position:absolute;
            top: 80px;
            left: 550px;
        }
        .badge{
            width: 700px;
            position: absolute;
            top: 17px;
            left: 10px;
        }

        .getbadge
        {
            position: absolute;
            top:15px;
            left:35px;
            width: 120px;
            display: inline-block;
        }
        .badge-list
        {
            padding: 0 0 0 0 ;
            position: absolute;
            top: 15px;
            left: 690px;
            display: inline-block;
        }
        .navigation{
            position: absolute;
            top:200px;
            width: 700px;
            height: 80px;
            background-color: #fff;
            font-size: 18px;
            margin-top: 70px;
            margin-left: 200px;
        }
        .text
        {
            text-decoration: none;
            color: black;
            margin-right: 40px;
            margin-left: 20px;
            font-size: 18px;
            font-weight: lighter;
        }
        .father
        {
            position: absolute;
        }
        .search
        {
            position: absolute;
            top:18px;
            left: 650px;
        }
        .navi{
            position: absolute;
            top:20px;
            left: 0px;
        }
        .nav-link-dropdown {
          font-size:20px;
          color: #1e80ff;
          cursor: pointer;
        }
</style>