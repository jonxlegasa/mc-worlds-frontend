package routes

import (
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/repos/msg"

	"github.com/labstack/echo/v4"
)

type logout struct {
	controller.Controller
}

func (l *logout) Get(c echo.Context) error {
	if err := l.Container.Auth.Logout(c); err == nil {

	} else {
		msg.Danger(c, "An error occurred. Please try again.")
	}
	return l.Redirect(c, routeNameLandingPage)
}
