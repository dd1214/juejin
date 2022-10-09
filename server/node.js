const express = require('express')
const mysql = require('mysql')
const uuid = require('uuid') // 如果 require 不行就 import
const jwt = require('jsonwebtoken')

// 服务
const app = express()
const port = 3008

const generateToken = function (req) {
  let secret = 'i love u'
  return jwt.sign(req, secret)
}

// mysql
const c = mysql.createConnection({
  host: 'gz-cynosdbmysql-grp-1oqt325v.sql.tencentcdb.com',
  user: 'root',
  password: '011214ljx.',
  database: 'dian',
  port:24276,
})

c.connect((err) => {
  if (err) {
    console.error(`连接错误: ${err}`)
    return
  }
  console.log(`连接成功: ${c.threadId}`)
})

// 服务
app.all("*", function (req, res, next) {
	//设置允许跨域的域名，*代表允许任意域名跨域
	res.header("Access-Control-Allow-Origin", "*");
	//允许的header类型
	res.header("Access-Control-Allow-Headers", "content-type");
	//跨域允许的请求方式 
	res.header("Access-Control-Allow-Methods", "DELETE,PUT,POST,GET,OPTIONS");
	if (req.method == 'OPTIONS')
		res.sendStatus(200); //让options尝试请求快速结束
	else
		next();
});

const bodyParser = require('body-parser')
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

app.get('/users', (req, res) => {
  // res.header('Access-Control-Allow-Origin', '*')
  // res.send()
  c.query('select * from users', (err, result, fields) => {
    if (err) throw err
    console.log(`fields: ${fields}`)
    console.log(`result: ${result}`)
    res.json(result)
  })
  console.log(req.url)
})
// users 注册
app.post('/users/register', (req, res, next) => {
  // res.header('Access-Control-Allow-Origin', '*')
  console.log(req.body)
  let id = uuid.v4()
  let token = generateToken(req.body)
  token = token.slice(token.length - 20, token.length)
  let username = req.body.username
  let pass = req.body.password
  c.query(`select * from users where username = '${username}'`, (err, result, fields) => {
    if (err) throw err
    if (result.length) {
      res.json({ flag: false })
      return
    }
    c.query(`insert into users (user_id, username, password, token) values ('${id}', '${username}', '${pass}', '${token}')`, (err, result) => {
      if (err) throw err
      res.json({...result, flag: true})
    })
  })
})

// users 登录
app.post('/users/login', (req, res) => {
  let username = req.body.username
  let pass = req.body.password
  c.query(`select token from users where username = '${username}' and password = '${pass}'`, (err, result, fields) => {
    if (err) throw err
    if (result.length) {
      res.json({
        ...result[0],
        flag: true,
      })
    } else res.json({ flag: false })
  })
})

app.listen(port, () => {
  console.log(`listening on ${port}`)
})
