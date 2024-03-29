<?php

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\API\AuthController;
use App\Http\Controllers\API\AvatarController;
use App\Http\Controllers\API\QuestionController;
use App\Http\Controllers\API\DiamondController;


Route::controller(AuthController::class)->group(function () {
    Route::post('login', 'login');
    Route::post('register', 'register');
    Route::post('logout', 'logout');
    Route::post('refresh', 'refresh');
    });

Route::controller(AvatarController::class)->group(function () {
    Route::get('avatars', 'findAll');
    Route::post('avatars', 'create');
    Route::delete('avatars/{id}', 'delete');
});

Route::controller(QuestionController::class)->group(function () {
    Route::get('questions', 'findAll');
    Route::post('questions', 'create');
    Route::get('questions/{id}', 'findOne');
    Route::delete('questions/{id}', 'delete');
});

Route::controller(DiamondController::class)->group(function () {
    Route::get('diamonds', 'findAll');
    Route::post('diamonds', 'create');
    Route::delete('diamonds/{id}', 'delete');
});