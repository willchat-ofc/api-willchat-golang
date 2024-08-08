package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
)

type DbFindChatById struct {
	FindChatByIdRepository protocols.FindChatByIdRepository
}

func NewDbFindChatById(findChatByIdRepository protocols.FindChatByIdRepository) *DbFindChatById {
	return &DbFindChatById{
		FindChatByIdRepository: findChatByIdRepository,
	}
}

func (c *DbFindChatById) Find(chatId string) (*models.Chat, error) {
	return c.FindChatByIdRepository.Find(chatId)
}
