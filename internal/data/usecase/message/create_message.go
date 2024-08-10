package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type DbCreateMessage struct{}

func NewDbCreateMessage() *DbCreateMessage {
	return &DbCreateMessage{}
}

func (c *DbCreateMessage) Create(data *usecase.CreateMessageInput) (*models.Message, error) {
	return nil, nil
}
