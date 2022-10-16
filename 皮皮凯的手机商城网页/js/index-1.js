var id = function(id){
  return document.getElementById(id);
}

//下载APP弹出
var a = id("downloadapp");
var b = id("download-image");

a.onmouseover=b.onmouseover=
function(){
  b.style.display="block";
}
a.onmouseout=b.onmouseout=
function(){
  b.style.display="none";
}

//导航栏Xiaomi手机弹出
var c = id("Xiaomi手机");
var d = id("Xiaomi手机-nav");
c.onmouseover=d.onmouseover=
function(){
  d.style.display="block";
}
c.onmouseout=d.onmouseout=
function(){
  d.style.display="none";
}

//左侧手机弹出
e = id("l")
f = id("shouji")
e.onmouseover=f.onmouseover=
function(){
  f.style.display="block";
}
e.onmouseout=f.onmouseout=
function(){
  f.style.display="none";
}

//中间图片变换
var array = new Array(5);
array[0] = "url('../image/index/change/0.jpg')"
array[1] = "url('../image/index/change/1.jpg')"
array[2] = "url('../image/index/change/2.jpg')"
array[3] = "url('../image/index/change/3.jpg')"
array[4] = "url('../image/index/change/4.jpg')"
var ray=["a-00","a-01","a-02","a-03","a-04"];
var i=0;

id("button-prev").onclick =function(){
  i--;
  change();
}
id("button-next").onclick =function(){
  i++;
  change();
}

e = id("c-img1");
function change() {
  i=(i+5)%5;
  // console.log("now:"+i);
  // console.log("减一:"+(i+4)%5);
  e.style.background=array[i];
  e.style.backgroundSize = "contain";

  // 右下角圆圈变换
  id(ray[(i+4)%5]).style.background="#666";
  id(ray[(i+6)%5]).style.background="#666";
  id(ray[i]).style.background="#FFF";
}