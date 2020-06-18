<?php

/** @var \Laravel\Lumen\Routing\Router $router */

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/', function () use ($router) {
    return $router->app->version();
});

$router->get('/', 'Controller@index');
$router->get('/backups/{id}', 'Controller@show_by_id');
$router->get('/files/{id}', 'Controller@show_files_by_id');
$router->post('/create', 'Controller@create');
$router->post('/file', 'Controller@file');
$router->post('/clear', 'Controller@clear');