<template>
<div class="all">
    <!--  个人成就 -->
    <div class="achievement">
        <h1>个人成就</h1>
        <hr>
        <ul>
            <li>
                <img src="../assets/info/like.jpg" alt="" class="ico">
                <p class="word">文章点赞数 3</p>
            </li>
            <br>
            <li>
                <img src="../assets/info/read.jpg" alt="" c lass="ico">
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
            width: 300px;
            height: 700px;
            position: absolute;
            top: 100px;
            right: 100px;
        }
        .achievement
        {
            width: 300px;
            height: 250px;
            background-color: #fff;
            border-width: 2px;
            border-color: rgb(112, 107, 107);
        }
        .state
        {
            width: 300px;
            height: 100px;
        }
        li{
            list-style-type: none;
            display: inline-block;
            float: left;
        }
        .word
        {
            display: inline-block;
            float: right;
            margin-left: 50px;
        }

        .center
        {
            margin-top: 20px;
            text-align: center;
            background-color: #fff;
        }
        .state-p
        {
            margin-right:30px;
            display: inline-block;
        }
        #num
        {
            margin-left: 20px;
            margin-right: 50px;
        }
        #info-btn
        {
            display: inline-block;
            font-size: 18px;
        }
        .info-data
        {
            float: right;
        }

        .margin
        {
            margin-top: 20px;
        }
</style>