package protocols

type CreateChatRepositoryOutput struct {
	Id        string
	OwnerId   string
	CreatedAt string
}

type CreateChatRepository interface {
	Create(ownerId string) (*CreateChatRepositoryOutput, error)
}
