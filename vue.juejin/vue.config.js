const path = require('path')
const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  // devServer: {
  //   proxy: {
  //     '/users': {
  //       target: 'http://diandian210.top',
	// 			pathRewrite:{'^/users':''},
  //       ws: true, //用于支持websocket
  //       changeOrigin: true //用于控制请求头中的host值
  //     }
  //   }
  // }
})

