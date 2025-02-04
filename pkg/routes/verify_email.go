package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/repos/msg"
)

type verifyEmail struct {
	controller.Controller
}

func (c *verifyEmail) Get(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := c.Container.Auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return c.Redirect(ctx, routeNameLandingPage)
	}

	// Check if it matches the authenticated user
	u := ctx.Get(context.AuthenticatedUserKey)
	if u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed
	if usr == nil {
		usr, err = c.Container.ORM.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return c.Fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed
	if !usr.Verified {
		_, err = usr.
			Update().
			SetVerified(true).
			Save(ctx.Request().Context())

		if err != nil {
			return c.Fail(err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")

	// If we have a user, they are already logged in and just redirect them to their home feed
	if u != nil {
		return c.Redirect(ctx, routeNamePreferences)

	}
	return c.Redirect(ctx, routeNameLogin)
}
