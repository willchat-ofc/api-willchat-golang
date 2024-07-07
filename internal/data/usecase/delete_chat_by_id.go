package usecase

import "github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"

type DbDeleteChatById struct {
	DeleteChatByIdRepository protocols.DeleteChatByIdRepository
}

func NewDbDeleteChatById(deleteChatByIdRepository protocols.DeleteChatByIdRepository) *DbDeleteChatById {
	return &DbDeleteChatById{
		DeleteChatByIdRepository: deleteChatByIdRepository,
	}
}

func (c *DbDeleteChatById) Delete(chatId string) error {
	return c.DeleteChatByIdRepository.Delete(chatId)
}
