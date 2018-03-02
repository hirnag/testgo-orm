package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
    "./dao"
)

func main() {
    // Echoのインスタンス作る
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // DB接続
    dao.Connect()

    // ルーティング
    e.GET("/", handler.Hello())
    e.GET("/hello", handler.Hello())
    e.GET("/sel", handler.SelectDB())
    e.GET("/mod", handler.Model())
    e.GET("/raw", handler.RawQuery())

    // サーバー起動
    e.Start(":8080")
}

