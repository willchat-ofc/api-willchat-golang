package chat_repository_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/infra/db/mongodb/chat_repository"
	"github.com/willchat-ofc/api-willchat-golang/tests/infra/db/mongodb/chat_repository_test/helper_test"
)

func setupCreateChatMongoRepositoryMocks(t *testing.T) (*chat_repository.CreateChatMongoRepository, func()) {
	db, teardown := helper_test.SetupMongoContainer(t)

	sut := chat_repository.NewCreateChatMongoRepository(db)

	return sut, teardown
}

func TestCreateChatMongoRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, teardown := setupCreateChatMongoRepositoryMocks(t)
		defer teardown()

		chatData, err := sut.Create("fake-user-id")
		require.NoError(t, err)

		require.NotEmpty(t, chatData.Id)
		require.Equal(t, chatData.OwnerId, "fake-user-id")
		require.NotEmpty(t, chatData.CreatedAt)
	})
}
