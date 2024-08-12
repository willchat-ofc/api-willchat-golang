package message_repository

import (
	"context"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindMessagesByChatIdMongoRepository struct {
	Db *mongo.Database
}

func NewFindMessagesByChatIdMongoRepository(db *mongo.Database) *FindMessagesByChatIdMongoRepository {
	return &FindMessagesByChatIdMongoRepository{
		Db: db,
	}
}

func (c *FindMessagesByChatIdMongoRepository) Find(chatId string) ([]*models.Message, error) {
	collection := c.Db.Collection("message")

	filter := bson.D{{Key: "chat_id", Value: chatId}}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	if err = cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
