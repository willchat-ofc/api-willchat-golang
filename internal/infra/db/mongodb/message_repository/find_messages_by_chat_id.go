package message_repository

import (
	"context"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	options := new(options.FindOptions)
	options.SetSkip(10)
	options.SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, options)

	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	if err = cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
