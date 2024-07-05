package routes

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/adapters"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChat(server *http.ServeMux, db *mongo.Database) {
	server.Handle("/", middlewares.VerifyAccessToken(adapters.AdaptRoute(controllers.NewCreateChatController())))
}
