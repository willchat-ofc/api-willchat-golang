package protocols

type GetAllChatsByOwnerIdRepositoryOutput struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	OwnerId   string `json:"owner_id"`
}

type GetAllChatsByOwnerIdRepository interface {
	Get(ownerId string) ([]*GetAllChatsByOwnerIdRepositoryOutput, error)
}
