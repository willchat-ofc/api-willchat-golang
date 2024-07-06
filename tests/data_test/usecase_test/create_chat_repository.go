package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"

	usecaseDomain "github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
)

func setupCreateChatRepositoryMocks(t *testing.T) (*usecase.DbCreateChat, *mocks.MockCreateChatRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockCreateChatRepository := mocks.NewMockCreateChatRepository(ctrl)
	sut := usecase.NewDbCreateChat(mockCreateChatRepository)

	return sut, mockCreateChatRepository, ctrl
}

func TestDbCreateChat(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, createChatRepository, ctrl := setupCreateChatRepositoryMocks(t)
		defer ctrl.Finish()

		chatData := &usecaseDomain.CreateChatOutput{
			Id:        "fake-chat-id",
			OwnerId:   "fake-user-id",
			CreatedAt: "fake-created-at",
		}

		createChatRepository.EXPECT().Create("fake-user-id").Return(chatData, nil)
		res, err := sut.Create("fake-user-id")

		require.NoError(t, err)
		require.Equal(t, res, chatData)
	})
}
