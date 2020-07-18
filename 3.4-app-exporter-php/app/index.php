<?php

use Laravel\Lumen\Application;
use Prometheus\CollectorRegistry;
use Prometheus\RenderTextFormat;
use Prometheus\Storage\APC;

require __DIR__ . '/vendor/autoload.php';

$app      = new Application();
$registry = new CollectorRegistry(new APC());

$counter = $registry->getOrRegisterCounter('', 'test_counter', 'Example of a counter', ['code']);
$gauge   = $registry->getOrRegisterGauge('', 'test_gauge', 'Example of a gauge', ['method', 'code']);

$counter->inc(['code' => 200]);
$counter->inc(['code' => 400]);

$app->router->get(
    '/send',
    static function () use ($gauge) {
        $gauge->set(mt_rand(0, 1000), ['method' => 'get', 'code' => '200']);
        $gauge->inc(['post' => '300']);
    }
);

$app->router->get(
    '/metrics',
    static function () use ($registry) {
        $renderer = new RenderTextFormat();
        $result = $renderer->render($registry->getMetricFamilySamples());

        header('Content-type: ' . RenderTextFormat::MIME_TYPE);
        echo $result;
    }
);
