<template>
<div class="personaltop">
            <div class="information">
               <div class="itemscope">
                <div class="head-photo" >
                <img src="../assets/info/head protrait.jpg" alt="" width="105px"  >
                    </div>
                    </div>
                    <div class="user-name">
                       <span style="color:black;font-size:xx-large; font-weight: bold;">
                        {{username}}
                       </span> 
                        <span>
                            <img src="../assets/info/l-grade.jpg" alt="" width="40px">
                        </span>
                        <span>
                            <img src="../assets/info/j-grade.jpg" alt="" class="j-level" width="40px">
                        </span>
                    <br /><br />
                    <br />
                 
                        <el-link :underline="false">+你从事什么职业？</el-link>
                        <br />
                        <el-link :underline="false">+你的爱好是？</el-link>
                    </div>
                   
                    <div class="link">
                        <a href="" class="link">
                           <img src="../assets/info/weibo.jpg" width="20" height="20">
                        </a>
                        <a href="" class="link">
                           <img src="../assets/info/github.jpg" width="20" height="20">
                        </a>
                        <a href="" class="link">
                           <img src="../assets/info/address.jpg" width="20" height="20">
                        </a>
                        </div>
                    <div class="click">
                        <el-button  plain style="border-color:#1e80ff;color: #1e80ff;" >编辑个人资料</el-button> 
                    </div>
                     <el-divider content-position="cnter"></el-divider>

                    <div class="father">
                       
                
                <div class="badge">
                    
                    <img src="../assets/info/badge.png" alt="" width="20px">
                </div>
                <div class="getbadge" style="font-size:14px;">获得徽章  0</div>
                <div>
                    <ul class="badge-list">
                        <img src="../assets/info/unfold.jpg" alt="" width="8px">
                    </ul>
                </div>
                
                </div>
                </div>
                
      
    
        <div class="navigation">
            <el-menu  mode="horizontal" >
            <el-menu-item index="1"><el-link >文章</el-link></el-menu-item>
            <el-menu-item index="2"><el-link >专栏</el-link></el-menu-item>
            <el-menu-item index="3"><el-link >沸点</el-link></el-menu-item>
            <el-menu-item index="4"><el-link >收藏集</el-link></el-menu-item>
            <el-menu-item index="5"><el-link >关注</el-link></el-menu-item>
            <el-menu-item index="6"><el-link >赞</el-link></el-menu-item>
          </el-menu> 
           
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
        .information
        {
            display: flex;
            flex-direction: column;
            justify-content:flex-star;
            align-items: flex-star;
        
        }
        .head-photo
        {
            margin-top: 38px;
            margin-left: 29px;
            display: inline-block;

        }
        .itemscope
        {
            width: 700px;
            height: 220px;
            background-color: #fff;
            margin-top: 10px;
            margin-left: 3px;
            margin-bottom: -10px;
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
            margin-left:170px;
            margin-top: -175px;
            justify-content:flex-star;
            align-items: flex-star;
        }
        

        .text1
        {
            text-decoration: none;
            color: rgb(10, 10, 186);
        }
        .link{
            margin-right: 21px;
            margin-top: -50px;
            display: flex;
            justify-content: end;
        }
        .click{
            
            margin-right: 30px;
            margin-top: 14px;
            display: flex;
            justify-content: end;
         
        }
        .badge{
           
            margin-top:-12px;
            margin-left: 43px;
            
        }

        .getbadge
        {
           
            margin-top:-11px;
            margin-left: 10px;
           
        }
        .badge-list
        {
            margin-top:-9px;
            margin-left: 525px;
           
        
        }
        .navigation{
            margin-top: 20px;
            margin-left: 3px;
            background-color: #fff;
           
        }
       
        .father
        {
            display: flex;
            justify-content:row;
        }
      
      

        
        
</style>