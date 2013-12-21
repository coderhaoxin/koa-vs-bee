var request = require('supertest')
request = request('http://localhost:8080')

var times = 1000
var startTime = Date.now()
var count = 0

for (var x = 0; x < times; x++) {
	request
	.post('/item')
	.send({
		id: 12345,
		name: 'coffee',
		quantity: 123,
		price: 2345,
		createTime: 1387617079074
	})
	.expect(200)
	.end(function (err, res) {
		if (err) {
			console.log(err)
		} else {
			count++
			if (count === times) {
				console.log('due (s): ', (Date.now() - startTime) / 1000)
				console.log(res.body)
			}
		}
	})
}
