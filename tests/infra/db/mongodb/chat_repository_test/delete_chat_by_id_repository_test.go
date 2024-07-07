package chat_repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/tests/infra/db/mongodb/chat_repository_test/helper_test"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupDeleteChatByIdRepositoryMocks(t *testing.T) (*chat_repository.DeleteChatByIdMongoRepository, *mongo.Database, func()) {
	db, teardown := helper_test.SetupMongoContainer(t)
	insertFakeChatsToGet(db)

	sut := chat_repository.NewDeleteChatByIdMongoRepository(db)

	return sut, db, teardown
}

func insertFakeChatsToGet(db *mongo.Database) {
	collection := db.Collection("chat")
	fakeChat := &models.Chat{
		Id:        "fake-chat-id",
		CreatedAt: time.Now(),
		OwnerId:   "fake-owner-id",
	}
	collection.InsertOne(context.TODO(), fakeChat)
}

func TestDeleteChatByIdRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, db, teardown := setupDeleteChatByIdRepositoryMocks(t)
		defer teardown()

		err := sut.Delete("fake-chat-id")
		require.NoError(t, err)

		chatCollection := db.Collection("chat")

		filter := bson.D{{Key: "_id", Value: "fake-chat-id"}}

		var chat models.Chat
		chatCollection.FindOne(context.TODO(), filter).Decode(&chat)

		require.Equal(t, models.Chat{}, chat)
	})

	t.Run("Error", func(t *testing.T) {
		sut, _, teardown := setupDeleteChatByIdRepositoryMocks(t)
		teardown()

		err := sut.Delete("fake-error")
		require.Error(t, err)
	})
}
