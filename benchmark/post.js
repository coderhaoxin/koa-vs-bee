'use strict';

var request = require('supertest');

var times = 100,
  startTime = Date.now(),
  count = 0;

for (var x = 0; x < times; x++) {
  request('http://localhost:3000')
    .post('/item')
    .send({
      id: 12345,
      name: 'coffee',
      quantity: 123,
      price: 2345,
      createTime: 1387617079074
    })
    .expect(200)
    .end(function(err, res) {
      if (err) {
        console.log(err);
      } else {
        count++;
        if (count === times) {
          console.log('due (ms): ', (Date.now() - startTime));
          // console.log(res.body);
        }
      }
    });
}