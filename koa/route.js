var r    = require('koa-route')
var item = require('./service/item')
var view = require('./service/view')

module.exports = function (app) {
	app.use(r.get('/item', item.get))
	app.use(r.post('/item', item.post))
	app.use(r.get('/view'), view.index)
}
