package protocols

import "github.com/willchat-ofc/api-willchat-golang/internal/domain/models"

type FindMessagesByChatIdRepository interface {
	Find(chatId string) ([]*models.Message, error)
}
