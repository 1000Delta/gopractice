package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// 直接访问
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// 路径参数
	e.GET("/name/:name", GetName)

	// 结构体绑定参数测试
	e.POST("/login", Login)

	// JWT 中间件测试
	userGroup := e.Group("/user", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtClaims{},
		SigningKey: []byte("secret"),
	}))

	// JWT 访问测试
	userGroup.GET("/info", UserInfo)
	e.Debug = true
	e.Logger.Fatal(e.Start(":8081"))
}
