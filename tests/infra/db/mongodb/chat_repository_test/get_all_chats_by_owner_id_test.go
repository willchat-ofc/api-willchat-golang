package chat_repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/tests/infra/db/mongodb/chat_repository_test/helper_test"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupGetAllChatsByOwnerIdRepositoryMocks(t *testing.T) (*chat_repository.GetAllChatsByOwnerIdMongoRepository, func()) {
	db, teardown := helper_test.SetupMongoContainer(t)

	insertFakeChatsToGet(db)
	sut := chat_repository.NewGetAllChatsByOwnerIdMongoRepository(db)

	return sut, teardown
}

func insertFakeChatsToGet(db *mongo.Database) {
	collection := db.Collection("chat")
	fakeChat := &models.Chat{
		Id:        "fake-id",
		CreatedAt: time.Now(),
		OwnerId:   "fake-owner-id",
	}
	collection.InsertOne(context.TODO(), fakeChat)
}

func TestGetAllChatsByOwnerIdRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, teardown := setupGetAllChatsByOwnerIdRepositoryMocks(t)
		defer teardown()

		chats, err := sut.Get("fake-owner-id")
		require.NoError(t, err)

		require.Len(t, chats, 1)
		require.Equal(t, chats[0].OwnerId, "fake-owner-id")
	})

	t.Run("Error", func(t *testing.T) {
		sut, teardown := setupGetAllChatsByOwnerIdRepositoryMocks(t)
		teardown()

		chatData, err := sut.Get("fake-owner-id")
		require.Error(t, err)
		require.Nil(t, chatData)
	})
}
