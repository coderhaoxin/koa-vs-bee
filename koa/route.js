'use strict';

var r = require('koa-route'),
  item = require('./service/item'),
  view = require('./service/view');

module.exports = function(app) {
  app.use(r.get('/item', item.get));
  app.use(r.post('/item', item.post));
  app.use(r.get('/view', view.index));
};