package handler

import (
	"edresco/model"
	"edresco/view/user"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@test.com",
	}
	return render(c, user.OnboardInfo(u))
}

func (h *UserHandler) Dashboard(c echo.Context) error {
	return render(c, user.Dashboard())
}

func (h *UserHandler) OnboardInfo(c echo.Context) error {
	time.Sleep(time.Second * 2)
	u := model.User{
		Email: "created@test.com",
	}
	return render(c, user.OnboardOtp(u))
}

func (h *UserHandler) VerifyOnboardOtp(c echo.Context) error {
	time.Sleep(time.Second * 2)

	return c.Redirect(301, "/dashboard")
}
