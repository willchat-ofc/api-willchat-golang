package factory

import (
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeCreateChatController(db *mongo.Database) *controllers.CreateChatController {
	createChatRepository := chat_repository.NewCreateChatMongoRepository(db)
	dbCreateChat := usecase.NewDbCreateChat(createChatRepository)

	return controllers.NewCreateChatController(dbCreateChat)
}
