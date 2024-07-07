package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/protocols"
	"github.com/willchat-ofc/api-willchat-golang/internal/data/usecase"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupDbGetAllChatsByOwnerIdMocks(t *testing.T) (*usecase.DbGetAllChatsByOwnerId, *mocks.MockGetAllChatsByOwnerIdRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockGetAllChatsByOwnerIdRepository := mocks.NewMockGetAllChatsByOwnerIdRepository(ctrl)
	sut := usecase.NewDbGetAllChatsByOwnerId(mockGetAllChatsByOwnerIdRepository)

	return sut, mockGetAllChatsByOwnerIdRepository, ctrl
}

func TestDbGetAllChatsById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, getAllChatsByOwnerIdRepository, ctrl := setupDbGetAllChatsByOwnerIdMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdResponse := &protocols.GetAllChatsByOwnerIdRepositoryOutput{
			Id:        "fake-id",
			CreatedAt: "fake-date",
			OwnerId:   "fake-owner-id",
		}

		var getAllChatsByOwnerIdResponseSlice []*protocols.GetAllChatsByOwnerIdRepositoryOutput
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)

		getAllChatsByOwnerIdRepository.EXPECT().Get("fake-owner-id").Return(getAllChatsByOwnerIdResponseSlice, nil)

		chatData, err := sut.Get("fake-owner-id")
		require.NoError(t, err)
		require.Equal(t, chatData[0].Id, "fake-id")
		require.Equal(t, chatData[0].CreatedAt, "fake-date")
		require.Equal(t, chatData[0].OwnerId, "fake-owner-id")
	})

	t.Run("GetAllChatsByOwnerIdRepositoryError", func(t *testing.T) {
		sut, getAllChatsByOwnerIdRepository, ctrl := setupDbGetAllChatsByOwnerIdMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdRepository.EXPECT().Get("fake-owner-id").Return(nil, errors.New("fake-error"))

		_, err := sut.Get("fake-owner-id")
		require.Error(t, err)
	})
}
