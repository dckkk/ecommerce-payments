package interfaces

import (
	"context"
	"ecommerce_payments/internal/models"

	"github.com/labstack/echo/v4"
)

type IPaymentRepo interface {
	InsertNewPaymentMethod(ctx context.Context, req *models.PaymentMethod) error
	DeletePaymentMethod(ctx context.Context, userID int, sourceID int, sourceName string) error
}

type IPaymentService interface {
	PaymentMethodLink(ctx context.Context, req models.PaymentMethodLinkRequest) error
	PaymentMethodLinkConfirm(ctx context.Context, userID int, req models.PaymentMethodOTPRequest) error
	PaymentMethodUnlink(ctx context.Context, userID int, req models.PaymentMethodLinkRequest) error
}

type IPaymentAPI interface {
	PaymentMethodLink(e echo.Context) error
	PaymentMethodOTP(e echo.Context) error
	PaymentMethodUnlink(e echo.Context) error
}
