'use strict';

var request = require('supertest');
request = request('http://localhost:3000');

var times = 1000,
  startTime = Date.now(),
  count = 0;

for (var x = 0; x < times; x++) {
  request
    .get('/view')
    .expect(200)
    .end(function(err, res) {
      if (err) {
        console.log(err);
      } else {
        count++;
        if (count === times) {
          console.log(res.text);
          console.log('due (s): ', (Date.now() - startTime) / 1000);
        }
      }
    });
}