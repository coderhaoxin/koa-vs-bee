var parse = require('co-body')

exports.get = function* () {
	this.body = {
		id: 10086,
		name: 'tea',
		quantity: 100,
		price: 1000,
		createTime: Date.now()
	}
}


exports.post = function* () {
	var body = yield parse(this)
	this.body = body
}
