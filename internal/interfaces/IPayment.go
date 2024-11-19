package interfaces

import (
	"context"
	"ecommerce_payments/internal/models"

	"github.com/labstack/echo/v4"
)

type IPaymentRepo interface {
	InsertNewPaymentMethod(ctx context.Context, req *models.PaymentMethod) error
	DeletePaymentMethod(ctx context.Context, userID int, sourceID int, sourceName string) error
	GetPaymentMethod(ctx context.Context, userID int, sourceName string) (models.PaymentMethod, error)
	GetPaymentMethodByID(ctx context.Context, paymentMethodID int) (models.PaymentMethod, error)

	InsertNewPaymentTransaction(ctx context.Context, req *models.PaymentTransaction) error
	InsertNewPaymentRefund(ctx context.Context, req *models.PaymentRefund) error
	GetPaymentByOrderID(ctx context.Context, orderID int) (models.PaymentTransaction, error)
}

type IPaymentService interface {
	PaymentMethodLink(ctx context.Context, req models.PaymentMethodLinkRequest) error
	PaymentMethodLinkConfirm(ctx context.Context, userID int, req models.PaymentMethodOTPRequest) error
	PaymentMethodUnlink(ctx context.Context, userID int, req models.PaymentMethodLinkRequest) error
	InitiatePayment(ctx context.Context, req models.PaymentInitiatePayload) error
	RefundPayment(ctx context.Context, req models.RefundPayload) error
}

type IPaymentAPI interface {
	PaymentMethodLink(e echo.Context) error
	PaymentMethodOTP(e echo.Context) error
	PaymentMethodUnlink(e echo.Context) error

	InitiatePayment(kafkaPayload []byte) error
	RefundPayment(kafkaPayload []byte) error
}
