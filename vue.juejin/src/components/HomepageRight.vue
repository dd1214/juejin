<template>
<div class="all">
    <!--  个人成就 -->
    <div class="achievement">
        <p class="achive">个人成就</p>
        <hr class="margin">
        <ul>
            <li>
                <img src="../assets/info/like.jpg" alt="" class="ico">
                <p class="word">文章点赞数 3</p>
            </li>
            <br>
            <li>
                <img src="../assets/info/read.jpg" alt="" class="ico">
                <p class="word">文章被阅读 217</p>
            </li>
            <br>
            <li>
                <img src="../assets/info/value.jpg" alt="" class="ico">
                <p class="word">掘力值</p>
            </li>
            <br>
        </ul>
    </div>
    <!-- 关注状态 -->
    <div class="state">
            <div class="center">
                <p class="state-p">关注了</p>
                <p class="state-p">关注者</p>
                <br>
                <p class="state-p" id="num">2</p>
                <p class="state-p" id="num">0</p>
            </div> 
    </div>
    <hr class="margin">
    <!-- 个人信息 -->
    <div class="per-info">
        <div>
            <div id="info-btn" >收藏集</div>
            <div id="info-btn" class="info-data">1</div>
        </div>
        <hr>
        <div>
            <div id="info-btn">关注标签</div>
            <div id="info-btn" class="info-data">9</div>
        </div>
        <hr>
        <div>
            <div id="info-btn">加入于</div>
            <div id="info-btn" class="info-data">2021-5-25</div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    name:"homepage-right"
}
</script>

<style scoped>
        body{
            background-color: #f5f5f5;
        }
        .all
        {
            position: absolute;
            top:80px;
            left:930px;
        }
        .achive{
            font-size: 20px;
            margin-left: 10px;
        }
        .achievement
        {
            width: 200px;
            height: 230px;
            background-color: #fff;
            border-width: 2px;
            border-color: rgb(112, 107, 107);
        }
        .state
        {
            width: 300px;
            height: 60px;
        }
        li{
            list-style-type: none;
            display: inline-block;
            float: left;
        }
        
        .ico
        {
            margin-top: 20px;
            margin-left: 10px;
            width: 40px;
        }
        .word
        {
            font-size: 13px;
            display: inline-block;
            font-weight: lighter;
            float: right;
            margin-left: 35px;
            margin-top: 33px;
        }

        .center
        {
            height: 50px;
            margin-top: 10px;
            width: 200px;
            text-align: center;
            background-color: #fff;
        }
        .state-p
        {
            font-size: 15px;
            font-weight: lighter;
            margin-right:30px;
            display: inline-block;
        }
        #num
        {
            margin-left: 20px;
            margin-right: 50px;
        }
        .per-info
        {
            width: 200px;
        }
        #info-btn {
            margin-top: 10px;
            margin-bottom: 10px;
            display: inline-block;
            font-size: 15px;
            font-weight: lighter;
        }
        .info-data
        {
            float: right;
        }

        .margin
        {
            width: 200px;
            margin-top: 10px;
            /* margin-bottom: 10px; */
        }
</style>