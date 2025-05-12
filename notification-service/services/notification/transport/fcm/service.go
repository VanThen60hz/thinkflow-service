package fcm

import (
	"context"
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

func (s *Service) SendNotification(ctx context.Context, userID, title, body string, data map[string]string) error {
	tokens, err := s.tokenRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		return nil
	}

	var fcmTokens []string
	for _, token := range tokens {
		fcmTokens = append(fcmTokens, token.Token)
	}

	message := &messaging.MulticastMessage{
		Tokens: fcmTokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}

	_, err = s.client.SendMulticast(ctx, message)
	if err != nil {
		log.Printf("Error sending FCM multicast: %v", err)
		return err
	}

	return nil
}
