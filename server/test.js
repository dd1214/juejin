const express = require('express')
const app = express()
const port=3030
app.get('/', (req, res) => {
  const { name = 'World' } = req.query;
  return res.send(`Hello ${name}!`);
})
app.listen(port, () => {
  console.log(`listening on ${port}`)
})