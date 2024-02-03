package main

import (
	"edresco/handler"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Println(err)
	code := http.StatusInternalServerError
	c.String(code, "error")
}

func main() {
	app := echo.New()
	app.HTTPErrorHandler = customHTTPErrorHandler
	app.Static("/public", "public")
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.GET("/dashboard", userHandler.Dashboard)
	app.POST("/onboard/info", userHandler.OnboardInfo)
	app.POST("/onboard/otp", userHandler.VerifyOnboardOtp)
	app.Logger.Fatal(app.Start(":1323"))
}
