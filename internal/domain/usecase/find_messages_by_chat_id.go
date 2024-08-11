package usecase

import "github.com/willchat-ofc/api-willchat-golang/internal/domain/models"

type FindMessagesByChatId interface {
	Find(chatId string) ([]*models.Message, error)
}
