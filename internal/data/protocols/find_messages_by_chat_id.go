package protocols

import "github.com/willchat-ofc/api-willchat-golang/internal/domain/models"

type FindMessagesByChatIdRepositoryInput struct {
	ChatId string
	Limit  int
	Offset int
}

type FindMessagesByChatIdRepository interface {
	Find(data *FindMessagesByChatIdRepositoryInput) ([]*models.Message, error)
}
