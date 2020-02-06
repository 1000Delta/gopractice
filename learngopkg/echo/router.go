package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetName 获取名称
func GetName(c echo.Context) error {
	if c.Param("name") == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, "Hello World, "+c.Param("name"))
}

// LoginParams 登录参数对象
type LoginParams struct {
	User string `json:"user" query:"user" form:"user"`
	Pass string `json:"pass" query:"pass" form:"pass"`
}

type jwtClaims struct {
	User string `json:"user"`
	Exp int `json:"exp"`
	jwt.StandardClaims
}

// Login 登录
func Login(c echo.Context) error {
	params := &LoginParams{}
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusPaymentRequired, err.Error())
	}
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// 添加声明（claim）
	claims["user"] = params.User
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	// 编码
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func UserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtClaims)
	return c.String(http.StatusOK, fmt.Sprintf("user: %s", claims.User))
}
