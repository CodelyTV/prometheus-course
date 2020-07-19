const express = require('express');
const client = require('prom-client');
const server = express();

const counter = require('prom-client').Counter;
const gauge = require('prom-client').Gauge;

const c = new counter({
    name: 'test_counter',
    help: 'Example of a counter',
    labelNames: ['code']
});
const g = new gauge({
    name: 'test_gauge',
    help: 'Example of a gauge',
    labelNames: ['method', 'code']
});

setInterval(() => {
    c.inc(({ code: 200 }))
}, 1000);

setInterval(() => {
    c.inc(({ code: 400 }))
}, 4000);

server.get('/send', function (req, res) {
    g.set({ method: 'get', code: 200 }, Math.random());
    g.set(Math.random());
    g.labels('post', '300').inc();
    res.end();
});

server.get('/metrics', function (req, res) {
    res.end(client.register.metrics());
});

server.listen(8081, '0.0.0.0');