package factory

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeGetAllChatsByOwnerIdController(db *mongo.Database) *controllers.GetAllChatsByOwnerIdController {
	getAllChatsByOwnerIdRepository := chat_repository.NewGetAllChatsByOwnerIdMongoRepository(db)
	dbGetAllChatsByOwnerId := usecase.NewDbGetAllChatsByOwnerId(getAllChatsByOwnerIdRepository)

	return controllers.NewGetAllChatsByOwnerIdController(dbGetAllChatsByOwnerId)
}
