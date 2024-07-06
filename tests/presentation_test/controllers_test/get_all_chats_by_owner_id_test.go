package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupGetAllChatsByOwnerIdMocks(t *testing.T) (*controllers.GetAllChatsByOwnerIdController, *mocks.MockGetAllChatsByOwnerId, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockGetAllChatsByOwnerId := mocks.NewMockGetAllChatsByOwnerId(ctrl)

	sut := controllers.NewGetAllChatsByOwnerIdController(mockGetAllChatsByOwnerId)

	return sut, mockGetAllChatsByOwnerId, ctrl
}

func createGetAllChatsByOwnerIdHttpRequest() protocols.HttpRequest {
	header := http.Header{}
	header.Add("UserId", "fake-user-id")

	return protocols.HttpRequest{
		Body:   nil,
		Header: header,
	}
}

func TestGetAllChatsByOwnerId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, getAllChatsByOwnerId, ctrl := setupGetAllChatsByOwnerIdMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdResponse := &usecase.GetAllChatsByOwnerIdOutput{
			Id:        "fake-id",
			CreatedAt: "fake-date",
			OwnerId:   "fake-owner-id",
		}
		var getAllChatsByOwnerIdResponseSlice []*usecase.GetAllChatsByOwnerIdOutput
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)

		getAllChatsByOwnerId.EXPECT().Get("fake-user-id").Return(getAllChatsByOwnerIdResponseSlice, nil)

		res := sut.Handle(createGetAllChatsByOwnerIdHttpRequest())

		var responseBody controllers.GetAllChatsByOwnerIdControllerResponse
		err := json.NewDecoder(res.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctGetAllChatsByOwnerIdResponse := &controllers.GetAllChatsByOwnerIdControllerResponse{
			Chats: getAllChatsByOwnerIdResponseSlice,
		}
		require.Equal(t, &responseBody, correctGetAllChatsByOwnerIdResponse)
	})

	t.Run("GetAllChatsByOwnerIdError", func(t *testing.T) {
		sut, getAllChatsByOwnerId, ctrl := setupGetAllChatsByOwnerIdMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerId.EXPECT().Get("fake-user-id").Return(nil, errors.New("fake-error"))

		res := sut.Handle(createGetAllChatsByOwnerIdHttpRequest())

		verifyHttpResponse(t, res, http.StatusInternalServerError, "an error ocurred while finding chats")
	})
}
