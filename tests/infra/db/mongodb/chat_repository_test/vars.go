package chat_repository

import (
	"context"
	"time"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertFakeChatsToGet(db *mongo.Database) {
	collection := db.Collection("chat")
	fakeChat := &models.Chat{
		Id:        "fake-chat-id",
		CreatedAt: time.Now(),
		OwnerId:   "fake-owner-id",
	}
	collection.InsertOne(context.TODO(), fakeChat)
}
