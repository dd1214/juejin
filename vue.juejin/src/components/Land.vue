<template>
  <div class="login-section">
    <!-- :rules="rules" -->
    <el-form
      label-position="top"
      label-width="100px" class="demo-ruleForm"
      :rules="rules"
      :model="rulesForm"
      status-icon
      ref="ruleForm"
    >
      <el-form-item label="用户名" prop="name">
        <el-input type="text" v-model="rulesForm.name"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="rulesForm.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
        <el-button>重置</el-button>
      </el-form-item>
    </el-form>
  </div>
  
</template>
<script>

export default {
  data() {
    return {
      //存储数据的对象
      rulesForm:{
        name:'',
        password:''
      },
      rules:{
        name:[
          {required:true,message:'名字',trigger:'blur'},
          {min:1,max:5,message:'长度在3到5',trigger:'blur'}
        ],
        password:[
          {required:true,message:'密码',trigger:'blur'},
          {min:3,max:5,message:'长度在3到5',trigger:'blur'}
        ]
      }
    };
  },
  methods: {
    submitForm(formName){
      this.$refs[formName].validate( (valid) =>{
        if(valid){
          //如果校检通过，在这里向后端发送用户名和密码
          login({
            name:this.rulesForm.name,
            password:this.rulesForm.password
          }).then((data) => {
            console.log(data)
            if(data.code === 0 ){
              localStorage.setItem('token',data.data.token)
              window.location.href = '/';
            }
            if(data.code === 1 ){
              this.$message.error(data.mes)
            }
          })
        }else{
          console.log('error submit!!');
          return false
        }
      });
    }
  }
}
</script>
<style scoped>
  .login-section{
    position:absolute; /*绝对定位*/
    top:50%; /*距顶部50%*/
    left:50%;
    margin:-100px 0 0 -150px; /*设定这个div的margin-top的负值为自身的高度的一半,margin-left的值也是自身的宽度的一半的负值.(感觉在绕口令)*/
    width:300px; /*宽为400,那么margin-top为-200px*/
    height:200px; /*高为200那么margin-left为-100px;*/
     z-index:99; /*浮动在最上层 */
    }
    

</style>
