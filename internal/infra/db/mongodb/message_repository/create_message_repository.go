package message_repository

import (
	"github.com/google/uuid"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateMessageMongoRepository struct {
	Db *mongo.Database
}

func NewCreateMessageMongoRepository(db *mongo.Database) *CreateMessageMongoRepository {
	return &CreateMessageMongoRepository{
		Db: db,
	}
}

func (c *CreateMessageMongoRepository) Create(data *usecase.CreateMessageInput) (*models.Message, error) {
	collection := c.Db.Collection("message")

	id := uuid.New().String()
	_, err := collection.InsertOne(helpers.Ctx, &models.Message{
		Id:         id,
		ChatId:     data.ChatId,
		Message:    data.Message,
		AuthorName: data.AuthorName,
		AuthorId:   data.AuthorId,
	})

	if err != nil {
		return nil, err
	}

	return &models.Message{
		Id:         id,
		ChatId:     data.ChatId,
		Message:    data.Message,
		AuthorName: data.AuthorName,
		AuthorId:   data.AuthorId,
	}, nil
}
