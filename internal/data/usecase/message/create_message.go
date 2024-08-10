package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type DbCreateMessage struct {
	CreateMessageRepository protocols.CreateMessageRepository
}

func NewDbCreateMessage(createMessageRepository protocols.CreateMessageRepository) *DbCreateMessage {
	return &DbCreateMessage{
		CreateMessageRepository: createMessageRepository,
	}
}

func (c *DbCreateMessage) Create(data *usecase.CreateMessageInput) (*models.Message, error) {
	return c.CreateMessageRepository.Create(data)
}
