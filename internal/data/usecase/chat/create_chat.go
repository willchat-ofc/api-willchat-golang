package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type DbCreateChat struct {
	CreateChatRepository protocols.CreateChatRepository
}

func NewDbCreateChat(createChatRepository protocols.CreateChatRepository) *DbCreateChat {
	return &DbCreateChat{
		CreateChatRepository: createChatRepository,
	}
}

func (c *DbCreateChat) Create(ownerId string) (*usecase.CreateChatOutput, error) {
	chat, err := c.CreateChatRepository.Create(ownerId)

	if err != nil {
		return nil, err
	}

	return &usecase.CreateChatOutput{
		Id:        chat.Id,
		OwnerId:   chat.OwnerId,
		CreatedAt: chat.CreatedAt,
	}, nil
}
