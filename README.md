# 1、运行项目的注意事项
## 1.1 juejin_test.sql
这个文件是导出的测试数据库文件，里面内容不多，有一个 juejin_test 数据库，里面重点是 users 表

**使用方法**：
  1. 本机安装 MySQL
  2. 注册本机 root 账户，并记住密码
  3. 进入 MySQL workbench，上方 Server - Data Import 导入 .sql 文件

## 1.2 server 文件夹
本文件夹内保存了 Node.js 写的本地服务器，想要实现注册登录注销功能的话必须开启这台服务器。

**使用方法**：
  1. 本机安装 Node.js 并添加环境变量（可参考网上教程）
  2. 安装完成之后命令行进入 server 文件夹，执行 `npm i` 来安装依赖
  3. **注意在本文件中的 `mysql.createConnection` 函数中修改 password 为 1.1 中 root 账户的密码**
  3. 执行 `node server-zzq.js` 来运行服务


