package main

//echo 版本 必须 1.8以上
//

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
    e.GET("/", hello)
	e.GET("/test", hello)

	//静态文件
	e.Static("/test", "test.html")

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}