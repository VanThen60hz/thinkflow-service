package fcm

import (
	"context"
	"fmt"
	"log"

	"thinkflow-service/services/notification/model"
	"thinkflow-service/services/notification/repository"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type Service struct {
	client    *messaging.Client
	tokenRepo *repository.FCMTokenRepository
}

func NewService(credentialsPath string, tokenRepo *repository.FCMTokenRepository) (*Service, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &Service{
		client:    client,
		tokenRepo: tokenRepo,
	}, nil
}

func (s *Service) RegisterToken(ctx context.Context, userID, token, deviceID, platform string) error {
	fcmToken := &model.FCMToken{
		UserID:   userID,
		Token:    token,
		DeviceID: deviceID,
		Platform: platform,
	}
	return s.tokenRepo.Save(ctx, fcmToken)
}

func (s *Service) UnregisterToken(ctx context.Context, token string) error {
	return s.tokenRepo.Delete(ctx, token)
}

func (s *Service) SendNotification(ctx context.Context, userID, title, body string, data map[string]string) error {
	tokens, err := s.tokenRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		return nil
	}

	// Send notifications one by one to avoid batch API issues
	successCount := 0
	for _, token := range tokens {
		message := &messaging.Message{
			Token: token.Token,
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Data: data,
		}

		// Try to send the message
		_, err := s.client.Send(ctx, message)
		if err != nil {
			log.Printf("Error sending FCM to token %s: %v", token.Token, err)
			continue
		}

		successCount++
	}

	if successCount == 0 && len(tokens) > 0 {
		return fmt.Errorf("failed to send notification to any of %d tokens", len(tokens))
	}

	return nil
}
