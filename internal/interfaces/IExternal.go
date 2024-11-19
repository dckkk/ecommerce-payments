package interfaces

import (
	"context"
	"ecommerce_payments/external"
	"ecommerce_payments/internal/models"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (external.Profile, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error

	PaymentLink(ctx context.Context, req models.PaymentMethodLinkRequest) (external.PaymentLinkResponse, error)
	PaymentUnlink(ctx context.Context, req models.PaymentMethodLinkRequest) (external.PaymentLinkResponse, error)
	PaymentLinkConfirmation(ctx context.Context, walletID int, otp string) (external.PaymentLinkResponse, error)
	WalletTransaction(ctx context.Context, req external.PaymentTransactionRequest) (external.PaymentTransactionResponse, error)

	OrderCallback(ctx context.Context, orderID int, status string) (external.OrderResponse, error)
}
