const REGISTER_URL_TEST = 'http://diandian210.top/users/register'

export default {
  namespaced: true,
  state: {
    isRegisterDialogShow: false,
  },
  actions: {
    // 注册操作
    register(context, value){
      console.log('register-actions:',context)
      let username = value.name
      let password = value.password
      if(!username.trim().length){
        alert('用户名不能为空！')
        return
      }
      if(!password.trim().length){
        alert('密码不能为空！')
        return
      }
      // 向服务端发送请求，进行登录
      fetch(REGISTER_URL_TEST, {
        method: 'post',
        body: JSON.stringify({
          username,
          password
        }),
        headers: {
          'Content-type': 'application/json'
        }
      }).then(res => res.json())
      .then(data => {
        console.log(`response: ${data}`) 
        // data(本地测试): {
        //   token: 登录成功时会返回 token， 失败时没有
        //   flag: true | false, 表示是否登录成功
        // }


      //   data(远程)/**
      //   * 状态码，0-成功，其他值-失败
      //   */
      //  status_code: number;
      //  /**
      //   * 返回状态描述
      //   */
      //  status_msg: string;
      //  /**
      //   * 用户鉴权token
      //   */
      //  token: string;
      //  /**
      //   * 返回的用户id
      //   */
      //  user_id: number;
        if(!data.status_code){
          context.commit('REGISTERED', data.token)
          context.commit('CLOSE_REGISTER_DIALOG')
        }else{
          alert('注册失败，请检查输入！')
        }
      })
    }
  },
  mutations: {
    OPEN_REGISTER_DIALOG(state){
      state.isRegisterDialogShow = true
    },
    CLOSE_REGISTER_DIALOG(state){
      state.isRegisterDialogShow = false
      // this.commit('registerOptions/CLEAR_USER_INPUT')
    },
    // 注册成功后，关闭注册弹窗，打开登录弹窗
    REGISTERED(state){
      this.commit('registerOptions/CLOSE_REGISTER_DIALOG')
      this.commit('loginOptions/OPEN_LOGIN_DIALOG')
    }
    // // 清空用户输入
    // CLEAR_USER_INPUT(state){
    //   state.userInput.username = ''
    //   state.userInput.password = ''
    // }
  }
}