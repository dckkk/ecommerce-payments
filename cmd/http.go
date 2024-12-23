package cmd

import (
	"ecommerce_payments/external"
	"ecommerce_payments/helpers"
	"ecommerce_payments/internal/api"
	"ecommerce_payments/internal/interfaces"
	"ecommerce_payments/internal/repository"
	"ecommerce_payments/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()
	healthcheckAPI := &api.HealthcheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.Healthcheck)

	paymentV1 := e.Group("/payment/v1")
	paymentV1.POST("/link", d.PaymentAPI.PaymentMethodLink, d.MiddlewareValidateAuth)
	paymentV1.POST("/link/confirm", d.PaymentAPI.PaymentMethodOTP, d.MiddlewareValidateAuth)
	paymentV1.POST("/unlink", d.PaymentAPI.PaymentMethodUnlink, d.MiddlewareValidateAuth)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthcheckAPI

	PaymentAPI interfaces.IPaymentAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	paymentRepo := &repository.PaymentRepo{
		DB: helpers.DB,
	}
	paymentSvc := &services.PaymentService{
		PaymentRepo: paymentRepo,
		External:    external,
	}
	paymentAPI := &api.PaymentAPI{
		PaymentService: paymentSvc,
	}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthcheckAPI{},

		PaymentAPI: paymentAPI,
	}
}
