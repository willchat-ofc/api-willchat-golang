package routes

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/setup/adapters"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/factory"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMessage(server *http.ServeMux, db *mongo.Database) {
	server.Handle("POST /message", adapters.AdaptRoute(factory.MakeCreateMessageController(db)))
}
