package usecase

type GetAllChatsByOwnerIdOutput struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	OwnerId   string `json:"owner_id"`
}

type GetAllChatsByOwnerId interface {
	Get(ownerId string) ([]*GetAllChatsByOwnerIdOutput, error)
}
