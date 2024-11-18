package services

import (
	"context"
	"ecommerce_payments/helpers"
	"ecommerce_payments/internal/interfaces"
	"ecommerce_payments/internal/models"

	"github.com/pkg/errors"
)

type PaymentService struct {
	PaymentRepo interfaces.IPaymentRepo
	External    interfaces.IExternal
}

func (s *PaymentService) PaymentMethodLink(ctx context.Context, req models.PaymentMethodLinkRequest) error {
	resp, err := s.External.PaymentLink(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to request payment link to e-wallet")
	}
	helpers.Logger.WithField("OTP", resp.Data.OTP).Info("Link response is success. need otp confirm")
	return nil
}

func (s *PaymentService) PaymentMethodLinkConfirm(ctx context.Context, userID int, req models.PaymentMethodOTPRequest) error {
	_, err := s.External.PaymentLinkConfirmation(ctx, req.SourceID, req.OTP)
	if err != nil {
		return errors.Wrap(err, "failed to request payment link confirm to e-wallet")
	}

	paymentMethod := models.PaymentMethod{
		UserID:     userID,
		SourceID:   req.SourceID,
		SourceName: "fastcampus_ewallet",
	}
	return s.PaymentRepo.InsertNewPaymentMethod(ctx, &paymentMethod)
}

func (s *PaymentService) PaymentMethodUnlink(ctx context.Context, userID int, req models.PaymentMethodLinkRequest) error {
	_, err := s.External.PaymentUnlink(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to request payment unlink to e-wallet")
	}

	return s.PaymentRepo.DeletePaymentMethod(ctx, userID, req.SourceID, "fastcampus_ewallet")
}
