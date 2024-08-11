package factory

import (
	usecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/chat"
	messageUsecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/message"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/message_repository"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/message"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeCreateMessageController(db *mongo.Database) *controllers.CreateMessageController {
	findChatByIdRepository := chat_repository.NewFindChatByIdMongoRepository(db)
	dbFindChatById := usecase.NewDbFindChatById(findChatByIdRepository)

	createMessageRepository := message_repository.NewCreateMessageMongoRepository(db)
	dbCreateMessage := messageUsecase.NewDbCreateMessage(createMessageRepository)
	return controllers.NewCreateMessageController(dbFindChatById, dbCreateMessage)
}
