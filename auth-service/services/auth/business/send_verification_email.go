package business

import (
	"context"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) SendVerificationEmail(ctx context.Context, email string) error {
	otp := core.GenerateOTP()
	return biz.SendOTPEmail(
		ctx,
		email,
		otp,
		common.EmailVerifyOTPSubject,
		"Email Verification",
		"Thanks for signing up! Please use the OTP below to verify your email:",
		"email verification",
	)
}
