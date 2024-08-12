package message_repository

import (
	"context"

	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
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

func (c *FindMessagesByChatIdMongoRepository) Find(data *protocols.FindMessagesByChatIdRepositoryInput) ([]*models.Message, error) {
	collection := c.Db.Collection("message")

	filter := bson.D{{Key: "chat_id", Value: data.ChatId}}

	options := new(options.FindOptions)
	options.SetSkip(int64(data.Offset))
	options.SetLimit(int64(data.Limit))
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
