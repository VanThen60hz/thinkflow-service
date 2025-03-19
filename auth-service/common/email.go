package common

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	smtpHost  string
	smtpPort  int
	emailUser string
	emailPass string
}

func NewEmailService(emailUser, emailPass string) *EmailService {
	return &EmailService{
		smtpHost:  "smtp.gmail.com",
		smtpPort:  587,
		emailUser: emailUser,
		emailPass: emailPass,
	}
}

func (e *EmailService) SendOTP(toEmail, otp string) error {
	auth := smtp.PlainAuth("", e.emailUser, e.emailPass, e.smtpHost)

	subject := "ThinkFlow - Password Reset OTP"

	// Create HTML email content
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background-color: #ffffff; border-radius: 5px; padding: 20px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <h2 style="color: #2c3e50; margin-bottom: 20px;">Password Reset Request</h2>
        <p>Hello,</p>
        <p>We received a request to reset your password for your ThinkFlow account. To proceed with the password reset, please use the following One-Time Password (OTP):</p>
        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 4px; text-align: center; margin: 20px 0;">
            <h3 style="color: #2c3e50; font-size: 24px; margin: 0;">%s</h3>
        </div>
        <p><strong>Important:</strong></p>
        <ul style="margin-bottom: 20px;">
            <li>This OTP will expire in 10 minutes</li>
            <li>If you didn't request this password reset, please ignore this email</li>
            <li>Do not share this OTP with anyone</li>
        </ul>
        <p>For security reasons, please complete the password reset process within the next 10 minutes.</p>
        <hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
        <p style="color: #666; font-size: 12px;">This is an automated message, please do not reply to this email. If you need assistance, please contact our support team.</p>
    </div>
</body>
</html>`, otp)

	// Email headers
	headers := make(map[string]string)
	headers["From"] = e.emailUser
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["List-Unsubscribe"] = fmt.Sprintf("<%s>", e.emailUser)

	// Build email message
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + htmlBody

	// Send email
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", e.smtpHost, e.smtpPort),
		auth,
		e.emailUser,
		[]string{toEmail},
		[]byte(message),
	)
}

func (e *EmailService) SendVerificationOTP(toEmail, otp string) error {
	auth := smtp.PlainAuth("", e.emailUser, e.emailPass, e.smtpHost)

	subject := "ThinkFlow - Email Verification OTP"

	// Create HTML email content
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background-color: #ffffff; border-radius: 5px; padding: 20px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <h2 style="color: #2c3e50; margin-bottom: 20px;">Email Verification</h2>
        <p>Hello,</p>
        <p>Thank you for registering with ThinkFlow! To complete your registration, please use the following One-Time Password (OTP):</p>
        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 4px; text-align: center; margin: 20px 0;">
            <h3 style="color: #2c3e50; font-size: 24px; margin: 0;">%s</h3>
        </div>
        <p><strong>Important:</strong></p>
        <ul style="margin-bottom: 20px;">
            <li>This OTP will expire in 10 minutes</li>
            <li>If you didn't register for a ThinkFlow account, please ignore this email</li>
            <li>Do not share this OTP with anyone</li>
        </ul>
        <p>For security reasons, please complete the verification process within the next 10 minutes.</p>
        <hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
        <p style="color: #666; font-size: 12px;">This is an automated message, please do not reply to this email. If you need assistance, please contact our support team.</p>
    </div>
</body>
</html>`, otp)

	// Email headers
	headers := make(map[string]string)
	headers["From"] = e.emailUser
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["List-Unsubscribe"] = fmt.Sprintf("<%s>", e.emailUser)

	// Build email message
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + htmlBody

	// Send email
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", e.smtpHost, e.smtpPort),
		auth,
		e.emailUser,
		[]string{toEmail},
		[]byte(message),
	)
}
