package services

import (
	"messaging_service/internal/kafka"
	"messaging_service/internal/models"
	"messaging_service/internal/repositories"
)

type MessageService interface {
	ProcessMessage(message *models.Message) error
	GetProcessedMessagesCount() (int, error)
	GetTotalMessagesCount() (int, error)
	GetUnprocessedMessagesCount() (int, error)
}

type messageService struct {
	repo  repositories.MessageRepository
	kafka kafka.KafkaProducer
}

func NewMessageService(repo repositories.MessageRepository, kafka kafka.KafkaProducer) MessageService {
	return &messageService{repo: repo, kafka: kafka}
}

func (s *messageService) ProcessMessage(message *models.Message) error {
	err := s.repo.SaveMessage(message)
	if err != nil {
		return err
	}
	err = s.kafka.SendMessage(message.Content)
	return err
}

func (s *messageService) GetProcessedMessagesCount() (int, error) {
	return s.repo.GetProcessedMessagesCount()
}

func (s *messageService) GetTotalMessagesCount() (int, error) {
	return s.repo.GetTotalMessagesCount()
}

func (s *messageService) GetUnprocessedMessagesCount() (int, error) {
	return s.repo.GetUnprocessedMessagesCount()
}
