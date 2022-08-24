<template>

<div class="personal-data">
            <div>
               <div class="itemscope">
                <img src="../assets/info/head protrait.jpg" alt="" width="100px" class="head-photo">
                    <h1 class="user-name">
                        user-name
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
                           <img src="../assets/info/weibo.jpg" alt="">
                        </a>
                        <a href="" class="click">
                           <img src="../assets/info/github.jpg" alt="">
                        </a>
                        <a href="" class="click">
                           <img src="../assets/info/address.jpg" alt="">
                        </a>
                    </div>

                    <div class="edit">
                        <button style="width: 150px;height:40px">编辑个人资料</button>
                    </div>
                </div>
             <hr>
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
    
        <div class="navigation">
            <div class="navi">
            <a href="" class="text">文章</a>
            <a href="" class="text">专栏</a> 
            <a href="" class="text">沸点</a>
            <a href="" class="text">收藏集</a>
            <a href="" class="text">关注</a>
            <a href="" class="text">赞</a>
            </div>
           <div class="search">
           <img src="../assets/info/search.jpg" alt="">
           </div>
        </div>
</div>
</template>

<script>
export default {
    name:"HomepageCenter",
}
</script>

<style scoped>
 .personal-data
        {
            width: 800px;
        }
        .head-photo
        {
            margin-top: 40px;
            display: inline-block;
        }
        .itemscope
        {
            width: 800px;
            background-color: #fff;
            display: inline-block;
            position: relative;
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
            margin-top:40px;
            margin-right:380px;
        }
        .introduce
        {
            position: absolute;
            top:90px;
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
            top: 50px;
            left: 600px;
        }
        .edit{
            position:absolute;
            top: 100px;
            left: 590px;
        }
        .badge{
            width: 800px;
            position: absolute;
            top: 17px;
            left: 10px;
        }

        .getbadge
        {
            position: absolute;
            top:15px;
            left:35px;
            display: inline-block;
        }
        .badge-list
        {
            padding: 0 0 0 0 ;
            position: absolute;
            top: 2px;
            left: 790px;
            display: inline-block;
        }
        .navigation{
            position: relative;
            width: 800px;
            height: 100px;
            margin-top: 70px;
        }
        .text
        {
            text-decoration: none;
            color: black;
            margin-right: 40px;
            font-size: 22px;
        }
        .father
        {
            position: relative;
        }
        .search
        {
            position: absolute;
            top:20px;
            left: 790px;
        }
        .navi{
            position: absolute;
            top:20px;
            left: 0px;
        }
</style>