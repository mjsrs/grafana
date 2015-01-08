package api

import (
	"github.com/torkelo/grafana-pro/pkg/bus"
	"github.com/torkelo/grafana-pro/pkg/middleware"
	m "github.com/torkelo/grafana-pro/pkg/models"
	"github.com/torkelo/grafana-pro/pkg/util"
)

func CreateAccount(c *middleware.Context) {
	var cmd m.CreateAccountCommand

	if !c.JsonBody(&cmd) {
		c.JsonApiErr(400, "Validation error", nil)
		return
	}

	cmd.Login = cmd.Email
	cmd.Salt = util.GetRandomString(10)
	cmd.Password = util.EncodePassword(cmd.Password, cmd.Salt)

	if err := bus.Dispatch(&cmd); err != nil {
		c.JsonApiErr(500, "failed to create account", err)
		return
	}

	c.JsonOK("Account created")
}
