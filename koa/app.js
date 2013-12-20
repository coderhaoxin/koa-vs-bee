var staticCache = require('koa-static-cache')
var koa         = require('koa')
var app         = koa()

app.use(staticCache('./static', { maxAge: 1000 }))

var route = require('./route')
route(app)

app.listen(3000)
console.log('server start on port: 3000')
