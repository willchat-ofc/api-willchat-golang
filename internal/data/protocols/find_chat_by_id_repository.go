package protocols

import "github.com/willchat-ofc/api-willchat-golang/internal/domain/models"

type FindChatByIdRepository interface {
	Find(chatId string) (*models.Chat, error)
}
