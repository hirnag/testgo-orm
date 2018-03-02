package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "../dao"
    modBoil "../models"
    "fmt"
)

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World")
    }
}
func Clac() echo.HandlerFunc {
    return func(c echo.Context) error {
        calcValue := c.Param("calcValue")
        return c.String(http.StatusOK, "result:" + calcValue)
    }
}
func SelectDB() echo.HandlerFunc {
    return func(c echo.Context) error {
        users, err := modBoil.UsersG().All()
        if err != nil {
            return err
        }
        res := modBoil.GetUsers(users)
        return c.String(http.StatusOK, "result on database:\n"+res)
    }
}
func Model() echo.HandlerFunc {
    return func(c echo.Context) error {
        res := fmt.Sprintf("%#v\n", modBoil.UserSlice{})
        return c.String(http.StatusOK, "result on models:\n"+res)
    }
}
func RawQuery() echo.HandlerFunc {
    return func(c echo.Context) error {
    	dao.Raw()
        res := fmt.Sprintf("%#v\n", modBoil.UserSlice{})
        return c.String(http.StatusOK, "result on models:\n"+res)
    }
}
