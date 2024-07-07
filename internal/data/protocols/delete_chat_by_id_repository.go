package protocols

type DeleteChatByIdRepository interface {
	Delete(chatId string) error
}
