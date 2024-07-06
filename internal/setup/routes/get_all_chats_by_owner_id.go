package routes

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/setup/adapters"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/factory"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllChatsByOwner(server *http.ServeMux, db *mongo.Database) {
	server.Handle("GET /chat", middlewares.VerifyAccessToken(adapters.AdaptRoute(factory.MakeGetAllChatsByOwnerIdController(db))))
}
