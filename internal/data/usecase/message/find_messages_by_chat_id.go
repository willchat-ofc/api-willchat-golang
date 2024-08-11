package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
)

type DbFindMessagesByChatId struct {
	FindMessagesByChatIdRepository protocols.FindMessagesByChatIdRepository
}

func NewDbFindMessagesByChatId(findMessagesByChatIdRepository protocols.FindMessagesByChatIdRepository) *DbFindMessagesByChatId {
	return &DbFindMessagesByChatId{
		FindMessagesByChatIdRepository: findMessagesByChatIdRepository,
	}
}

func (c *DbFindMessagesByChatId) Find(chatId string) ([]*models.Message, error) {
	return c.FindMessagesByChatIdRepository.Find(chatId)
}
