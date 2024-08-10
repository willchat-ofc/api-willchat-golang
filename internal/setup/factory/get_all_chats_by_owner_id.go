package factory

import (
	usecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/chat"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/chat"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeGetAllChatsByOwnerIdController(db *mongo.Database) *controllers.GetAllChatsByOwnerIdController {
	getAllChatsByOwnerIdRepository := chat_repository.NewGetAllChatsByOwnerIdMongoRepository(db)
	dbGetAllChatsByOwnerId := usecase.NewDbGetAllChatsByOwnerId(getAllChatsByOwnerIdRepository)

	return controllers.NewGetAllChatsByOwnerIdController(dbGetAllChatsByOwnerId)
}
