//  页面放大隐藏图片
var a = document.getElementById("left");
var r = document.getElementById("right");

function h() {
  // console.log("页面宽度变化！");
  // console.log(window.innerWidth);

  if (window.innerWidth<900){
    a.style.display='none';
    r.style.width='100%';
  }
  else if (window.innerWidth>=900){
    a.style.display='inline';
    r.style.width='78%';
  }

}

window.addEventListener('load',h);

window.addEventListener('resize',h);