package chat_repository

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/tests/infra/db/mongodb/chat_repository_test/helper_test"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupFindChatByIdRepositoryMocks(t *testing.T) (*chat_repository.FindChatByIdMongoRepository, *mongo.Database, func()) {
	db, teardown := helper_test.SetupMongoContainer(t)
	insertFakeChatsToGet(db)

	sut := chat_repository.NewFindChatByIdMongoRepository(db)

	return sut, db, teardown
}

func TestFindChatByIdRepositoryMocks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, _, teardown := setupFindChatByIdRepositoryMocks(t)
		defer teardown()

		chat, err := sut.Find("fake-chat-id")
		require.NoError(t, err)
		require.Equal(t, "fake-chat-id", chat.Id)
		require.Equal(t, chat.OwnerId, "fake-owner-id")
	})
}
