package factory

import (
	usecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/message"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/message_repository"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/message"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeFindMessagesByChatIdController(db *mongo.Database) *controllers.FindMessagesByChatIdController {
	findMessagesByChatIdRepository := message_repository.NewFindMessagesByChatIdMongoRepository(db)
	findMessagesByChatId := usecase.NewDbFindMessagesByChatId(findMessagesByChatIdRepository)

	return controllers.NewFindMessagesByChatIdController(findMessagesByChatId)
}
