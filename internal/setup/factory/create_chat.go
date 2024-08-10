package factory

import (
	usecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/chat"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/chat"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeCreateChatController(db *mongo.Database) *controllers.CreateChatController {
	createChatRepository := chat_repository.NewCreateChatMongoRepository(db)
	dbCreateChat := usecase.NewDbCreateChat(createChatRepository)

	return controllers.NewCreateChatController(dbCreateChat)
}
