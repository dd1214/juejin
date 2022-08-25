<template>
  <div class="login-section">
    <!-- :rules="rules" -->
    <el-form label-position="top" label-width="100px" class="demo-ruleForm" :rules="rules" :model="userInput" status-icon ref="ruleForm">
      <h2>用户登录</h2>
      <el-form-item label="用户名" prop="username">
        <el-input type="text" v-model="userInput.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="userInput.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item class="login-buttons">
        <el-button type="primary" @click="login(userInput)">登录</el-button>
        <el-button @click="closeLoginDialog">取消</el-button>
      </el-form-item>
      <!-- 注册链接，点击打开注册页面 -->
      <a href="#" @click="openRegisterDialog">注册</a>
    </el-form>
  </div>
</template>
<script>
import { mapState, mapActions, mapMutations } from 'vuex'
export default {
  data() {
    return {
      //存储数据的对象
      rules: {
        username: [
          { required: true, message: '名字', trigger: 'blur' },
          { min: 1, max: 9, message: '长度在1到9', trigger: 'blur' },
        ],
        password: [
          { required: true, message: '密码', trigger: 'blur' },
          { min: 3, max: 16, message: '长度在3到16', trigger: 'blur' },
        ],
      },
    }
  },
  computed: {
    ...mapState('loginOptions', ['userInput'])
  },
  methods: {
    // 点击“登录”时的逻辑
    ...mapActions('loginOptions', ['login']),
    // 点击“注册”时关闭登录弹窗，打开注册弹窗
    openRegisterDialog(){
      this.$store.commit('loginOptions/CLOSE_LOGIN_DIALOG')
      this.$store.commit('registerOptions/OPEN_REGISTER_DIALOG')
    },
    // 点击“取消”时，关闭登录弹窗
    ...mapMutations('loginOptions', {closeLoginDialog: 'CLOSE_LOGIN_DIALOG'})
  },
}
</script>
<style scoped>
h2{
  text-align: center;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
.login-section {
  /* 用来 mask */
  position: fixed; /*固定定位*/
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: hsla(120, 0%, 30%, 0.3);
  z-index: 98;
}
.el-form {
  position: fixed; /*固定定位*/
  top: 50%; /*距顶部50%*/
  left: 50%;
  /* margin: -100px 0 0 -150px; 设定这个div的margin-top的负值为自身的高度的一半,margin-left的值也是自身的宽度的一半的负值.(感觉在绕口令) */
  transform: translateX(-50%) translateY(-50%); /* 用平移实现居中更好，自适应 */
  width: 300px; /*宽为400,那么margin-top为-200px*/
  padding: 20px;
  border-radius: 8px;
  background-color: #f3f5f6;
  z-index: 99; /*浮动在最上层 */
}
.login-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
}
/* 注册 */
a{
  text-decoration: none;


}
a:link{
  color: black;
}
a:visited{
  color: black;
}
a:hover{
  color: #1e80ff;
}
</style>
