package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"

	domainUsecase "github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
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

		chatData := &protocols.CreateChatRepositoryOutput{
			Id:        "fake-chat-id",
			OwnerId:   "fake-user-id",
			CreatedAt: "fake-created-at",
		}

		createChatRepository.EXPECT().Create("fake-user-id").Return(chatData, nil)
		res, err := sut.Create("fake-user-id")

		expectedOutput := &domainUsecase.CreateChatOutput{
			Id:        "fake-chat-id",
			OwnerId:   "fake-user-id",
			CreatedAt: "fake-created-at",
		}

		require.NoError(t, err)
		require.Equal(t, res, expectedOutput)
	})

	t.Run("CreateChatRepositoryError", func(t *testing.T) {
		sut, createChatRepository, ctrl := setupCreateChatRepositoryMocks(t)
		defer ctrl.Finish()

		createChatRepository.EXPECT().Create("fake-user-id").Return(nil, errors.New("fake-error"))
		res, err := sut.Create("fake-user-id")

		require.Error(t, err)
		require.Nil(t, res)
	})
}
