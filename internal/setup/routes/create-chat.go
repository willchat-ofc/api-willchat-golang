package routes

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChat(server *http.ServeMux, db *mongo.Database) {
	// server.Handle("/", middlewares.VerifyAccessToken(adapters.AdaptRoute(controllers.NewCreateChatController())))
}
