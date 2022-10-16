var time = 0;//2秒后自动换图
var timeId = -1;//自动函数的ID值，默认-1
var banner = document.querySelector(".banner");//获取banner
var bannerImg = document.querySelectorAll(".bannerImg li");//获取所有的li(图片)
var dian = document.querySelectorAll(".dian li");//获取所有的小按钮
var right = document.getElementsByClassName("right")[0];//获取点击下一张
var left = document.getElementsByClassName("left")[0];//获取点击上一张
index = 0;//图片索引

// 点击下一张
right.onclick=function(){
    changeImg(true);
}

// 点击上一张
left.onclick=function(){
    changeImg(false);
}

// 自动轮播
function autoPlay(){
    timeId  = setInterval(function(){
        time++;
        if(time == 20){
            time = 0;
            changeImg(true);
        }
    },100);
}

autoPlay();

// 当鼠标移上banner空间时，停止自动轮播
banner.onmouseover=function(){
    clearTimeout(timeId);
}
// 当鼠标移开banner空间时，继续自动轮播
banner.onmouseout=function(){
    autoPlay();
}


// 清除样式
function clearStyle(index){
    // 清除上一个样式
    bannerImg[index].className = "";
    dian[index].className = "items";
}

// 添加样式
function addStyle(index){
    // 添加样式
    bannerImg[index].className = "on";
    dian[index].className = "items on";
}


/* 
	更改图片和小按钮样式
	flag:
		true:下一张
		false:上一张
*/
function changeImg(flag){
    time = 0;//用来阻止2秒后自动换图的bug
    clearStyle(index);
    // 判断是上一张还是下一张
    if(flag){
        // 下一张
        index++;
        index %= 5;
    }else{
        // 上一张
        index--;
        if(index<0) index = 0;
    }
    addStyle(index);
}

/* 
	点击小按钮切换到对应图片
*/
function buttonClick(num){
    time = 0;//用来阻止2秒后自动换图所产生的小bug
    clearStyle(index);
    index = num;
    addStyle(index);
}
