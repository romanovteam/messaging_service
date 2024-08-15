package repositories

import (
	"context"
	"messaging_service/internal/db"
	"messaging_service/internal/models"
	"time"
)

type MessageRepository interface {
	SaveMessage(message *models.Message) error
	MarkMessageProcessed(id int) error
	GetProcessedMessagesCount() (int, error)
	GetTotalMessagesCount() (int, error)
	GetUnprocessedMessagesCount() (int, error)
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) SaveMessage(message *models.Message) error {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO messages (content, processed, created_at, updated_at, description) VALUES ($1, $2, $3, $4, $5)",
		message.Content, message.Processed, message.CreatedAt, message.UpdatedAt, message.Description)
	return err
}

func (r *messageRepository) MarkMessageProcessed(id int) error {
	_, err := db.Pool.Exec(context.Background(),
		"UPDATE messages SET processed = TRUE, updated_at = $1 WHERE id = $2",
		time.Now(), id)
	return err
}

func (r *messageRepository) GetProcessedMessagesCount() (int, error) {
	var count int
	err := db.Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM messages WHERE processed = TRUE").Scan(&count)
	return count, err
}

func (r *messageRepository) GetTotalMessagesCount() (int, error) {
	var count int
	err := db.Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM messages").Scan(&count)
	return count, err
}

func (r *messageRepository) GetUnprocessedMessagesCount() (int, error) {
	var count int
	err := db.Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM messages WHERE processed = FALSE").Scan(&count)
	return count, err
}
