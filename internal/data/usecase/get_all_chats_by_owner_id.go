package usecase

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

type DbGetAllChatsByOwnerId struct {
	GetAllChatsByOwnerIdRepository protocols.GetAllChatsByOwnerIdRepository
}

func NewDbGetAllChatsByOwnerId() *DbGetAllChatsByOwnerId {
	return &DbGetAllChatsByOwnerId{}
}

func (c *DbGetAllChatsByOwnerId) Get(ownerId string) ([]*usecase.GetAllChatsByOwnerIdOutput, error) {
	res, err := c.GetAllChatsByOwnerIdRepository.Get(ownerId)

	if err != nil {
		return nil, err
	}

	return convertToUsecaseSlice(res), nil
}

func convertToUsecaseSlice(input []*protocols.GetAllChatsByOwnerIdRepositoryOutput) []*usecase.GetAllChatsByOwnerIdOutput {
	output := make([]*usecase.GetAllChatsByOwnerIdOutput, len(input))
	for i, item := range input {
		output[i] = &usecase.GetAllChatsByOwnerIdOutput{
			Id:        item.Id,
			CreatedAt: item.CreatedAt,
			OwnerId:   item.OwnerId,
		}
	}
	return output
}
