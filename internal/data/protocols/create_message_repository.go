package protocols

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type CreateMessageRepository interface {
	Create(data *usecase.CreateMessageInput) (*models.Message, error)
}
