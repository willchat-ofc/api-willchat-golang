package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

type DeleteChatById struct{}

func setupDeleteChatByIdRequest(t *testing.T) (*usecase.DbDeleteChatById, *mocks.MockDeleteChatByIdRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockDeleteChatByIdRequest := mocks.NewMockDeleteChatByIdRepository(ctrl)

	sut := usecase.NewDbDeleteChatById(mockDeleteChatByIdRequest)

	return sut, mockDeleteChatByIdRequest, ctrl
}

func TestDbDeleteChatById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, deleteChatByIdRepository, ctrl := setupDeleteChatByIdRequest(t)
		defer ctrl.Finish()

		deleteChatByIdRepository.EXPECT().Delete("fake-chat-id").Return(nil)

		err := sut.Delete("fake-chat-id")
		require.NoError(t, err)
	})
}
