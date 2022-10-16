var id = function(id){
  return document.getElementById(id);
}

// 购物车
sh =id("c-shopping");
shicon=id("shoppingicon");
g=id("gou");
d=id("shoping-list");
sh.onmouseover=d.onmouseover=
function(){
  sh.style.background="#fff";
  shicon.src="./image/index/购物车-橙.png";
  g.style.color="#FF7600";
  d.style.display="block";
}
sh.onmouseout=d.onmouseout=
function(){
  sh.style.background="#474747";
  shicon.src="./image/index/购物车-灰.png";
  g.style.color="#AEAEAE";
  d.style.display="none";
}



var a1 = id("a-1");
var i1 = id("a-1-i");
a1.onmouseover=
function(){
  a1.style.color="#FFFFFF";
  i1.src="./image/index/i-1-白.png"
}
a1.onmouseout=
function(){
  a1.style.color="#AEAEAE";
  i1.src="./image/index/i-1-灰.png"
}

var a2 = id("a-2");
var i2 = id("a-2-i");
a2.onmouseover=
function(){
  a2.style.color="#FFFFFF";
  i2.src="./image/index/i-2-白.png"
}
a2.onmouseout=
function(){
  a2.style.color="#AEAEAE";
  i2.src="./image/index/i-2-灰.png"
}

var a3 = id("a-3");
var i3 = id("a-3-i");
a3.onmouseover=
function(){
  a3.style.color="#FFFFFF";
  i3.src="./image/index/i-3-白.png"
}
a3.onmouseout=
function(){
  a3.style.color="#AEAEAE";
  i3.src="./image/index/i-3-灰.png"
}

var a4 = id("a-4");
var i4 = id("a-4-i");
a4.onmouseover=
function(){
  a4.style.color="#FFFFFF";
  i4.src="./image/index/i-4-白.png"
}
a4.onmouseout=
function(){
  a4.style.color="#AEAEAE";
  i4.src="./image/index/i-4-灰.png"
}

var a5 = id("a-5");
var i5 = id("a-5-i");
a5.onmouseover=
function(){
  a5.style.color="#FFFFFF";
  i5.src="./image/index/i-5-白.png"
}
a5.onmouseout=
function(){
  a5.style.color="#AEAEAE";
  i5.src="./image/index/i-5-灰.png"
}

var a6 = id("a-6");
var i6 = id("a-6-i");
a6.onmouseover=
function(){
  a6.style.color="#FFFFFF";
  i6.src="./image/index/i-6-白.png"
}
a6.onmouseout=
function(){
  a6.style.color="#AEAEAE";
  i6.src="./image/index/i-6-灰.png"
}