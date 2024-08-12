package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type DbFindMessagesByChatId struct {
	FindMessagesByChatIdRepository protocols.FindMessagesByChatIdRepository
}

func NewDbFindMessagesByChatId(findMessagesByChatIdRepository protocols.FindMessagesByChatIdRepository) *DbFindMessagesByChatId {
	return &DbFindMessagesByChatId{
		FindMessagesByChatIdRepository: findMessagesByChatIdRepository,
	}
}

func (c *DbFindMessagesByChatId) Find(data *usecase.FindMessagesByChatIdInput) ([]*models.Message, error) {
	return c.FindMessagesByChatIdRepository.Find(&protocols.FindMessagesByChatIdRepositoryInput{
		ChatId: data.ChatId,
		Limit:  data.Limit,
		Offset: data.Offset,
	})
}
