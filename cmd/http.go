package cmd

import (
	"ecommerce_payments/external"
	"ecommerce_payments/helpers"
	"ecommerce_payments/internal/api"
	"ecommerce_payments/internal/interfaces"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	// d := dependencyInject()
	healthcheckAPI := &api.HealthcheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.Healthcheck)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthcheckAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthcheckAPI{},
	}
}
