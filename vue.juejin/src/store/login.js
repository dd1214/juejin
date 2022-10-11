const LOGIN_URL = 'http://diandian.free.idcfengye.com/users/login'

export default {
  namespaced: true,
  state:{
    // 用户登录时输入的信息
    userInput: {
      username: '',
      password: ''
    },
    // 是否已经登录
    isLogin: false,
    // 登录对话框是否打开
    isLoginDialogShow: false,
    // 登录后的状态
    userLoginState: {
      token: ''
    }
  },
  actions: {
    // 输入登录信息，并发送 ajax 请求判断能否登录
    login(context, value){
      console.log('login:actions:',context)
      let username = context.state.userInput.username
      let password = context.state.userInput.password
      if(!username.trim().length){
        alert('用户名不能为空！')
        return
      }
      if(!password.trim().length){
        alert('密码不能为空！')
        return
      }
      // 向服务端发送请求，进行登录
      fetch(LOGIN_URL, {
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
        // data: {
        //   token: 登录成功时会返回 token， 失败时没有
        //   flag: true | false, 表示是否登录成功
        // }
        if(data.flag){
          context.commit('LOGIN', data.token)
          context.commit('CLOSE_LOGIN_DIALOG')
        }else{
          alert('登录失败，请检查输入！')
        }
        
      })
    }
  },
  mutations: {
    // 打开登录弹窗
    OPEN_LOGIN_DIALOG(state){
      state.isLoginDialogShow = true
    },
    // 关闭登录弹窗
    CLOSE_LOGIN_DIALOG(state){
      state.isLoginDialogShow = false
      // this.commit('loginOptions/CLEAR_USER_INPUT')
    },
    // 设置状态为登录, 并将获取的 token 保存到 localStorage
    LOGIN(state, token){
      state.isLogin = true
      // this.commit('loginOptions/CLEAR_USER_INPUT')
      localStorage.setItem('token', token)
      localStorage.setItem('username', state.userInput.username)
    },
    // 登出，清空 localStorage，刷新页面
    LOGOUT(state){
      state.isLogin = false
      this.commit('loginOptions/CLEAR_USER_INPUT')
      localStorage.clear()
      // location.reload()
    },
    // 清空用户输入
    CLEAR_USER_INPUT(state){
      state.userInput.username = ''
      state.userInput.password = ''
    }
  },

}