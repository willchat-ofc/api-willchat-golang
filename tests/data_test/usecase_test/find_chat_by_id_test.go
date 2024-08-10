package usecase_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	usecase "github.com/willchat-ofc/api-willchat-golang/internal/data/usecase/chat"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupDbFindChatByIdTest(t *testing.T) (*usecase.DbFindChatById, *mocks.MockFindChatByIdRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockFindChatByIdRepository := mocks.NewMockFindChatByIdRepository(ctrl)
	sut := usecase.NewDbFindChatById(mockFindChatByIdRepository)

	return sut, mockFindChatByIdRepository, ctrl
}

func TestDbFindChatById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, findChatByIdRepository, ctrl := setupDbFindChatByIdTest(t)
		defer ctrl.Finish()

		nowTime := time.Now()
		chatData := &models.Chat{
			Id:        "fake-id",
			OwnerId:   "fake-owner-id",
			CreatedAt: nowTime,
		}

		findChatByIdRepository.EXPECT().Find("fake-id").Return(chatData, nil)

		chat, err := sut.Find("fake-id")
		require.NoError(t, err)
		require.Equal(t, chat.Id, "fake-id")
		require.Equal(t, chat.OwnerId, "fake-owner-id")
		require.Equal(t, chat.CreatedAt, nowTime)
	})
}
