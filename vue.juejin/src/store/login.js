export default {
  namespaced: true,
  state:{
    // 是否已经登录
    isLogin: false,
    // 登录对话框是否打开
    isLoginDialogShow: false,
    // 保存用户登录状态
    userInfo: {

    }
  },
  actions: {
    // 输入登录信息，并发送 ajax 请求判断能否登录
    login(context, value){
      
      fetch()
    }
  },
  mutations: {
    // 打开登录弹窗
    OPEN_LOGIN_DIALOG(state){
      state.isLoginDialogShow = true
    },
    // 设置状态为登录, 并将获取的 token 保存到 localStorage
    LOGIN(state, token){
      state.isLogin = true
      
    }
  },

}