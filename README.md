# 1、运行项目的注意事项
## 1.1 juejin_test.sql
这个文件是导出的测试数据库文件，里面内容不多，有一个 juejin_test 数据库，里面重点是 users 表

**使用方法**：
  1. 本机安装 MySQL （建议安装 MySQL 5.x 后的版本，记得会自带 MySQL workbench）
  2. 注册本机 root 账户，并记住密码 （*）
  3. 进入 MySQL workbench，打开之后看左侧窗口下方的 schemas，右键 -- create schema 创建一个新的数据库，记住数据库名（*）
  4. 左上角 File -- open SQL script，选择 `juejin_test.sql`，打开后中间窗口左上角有黄色小闪电，点击执行。
  5. 在左侧窗口空白处 右键 -- Refresh All 刷新，可以看到多了刚才创建的数据库，里面有三张表，主要是 users 表用来实现登录注册注销。

## 1.2 server 文件夹
本文件夹内保存了 Node.js 写的本地服务器，想要实现注册登录注销功能的话必须开启这台服务器。

**使用方法**：
  1. 本机安装 Node.js 并添加环境变量（可参考网上教程）
  2. 安装完成之后命令行进入 server 文件夹，执行 `npm i` 来安装依赖
  3. **注意在本文件中的 `mysql.createConnection` 函数中修改 password 和 database 为 1.1 中 root 账户的密码 和 新建的数据库名**
  4. 当前目录下命令行执行 `node server-zzq.js` 来运行服务

  *ps：中间可能出现端口被占用的情况，此时可以
  1. 关闭本机占用 5000 端口的服务（不推荐）
  2. 修改 `server-zzq.js` 中的 `port` 为一个合理的端口号（一般大于 1024 即可）；之后在 vue.juejin 的项目代码里，找到 `/src/store/login.js`和`/src/store/register.js`，将最上方的 url 中的端口号修改为刚才指定的端口号。



