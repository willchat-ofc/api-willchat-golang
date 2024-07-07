package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupDeleteChatControllerMocks(t *testing.T) (*controllers.DeleteChatController, *mocks.MockGetAllChatsByOwnerId, *mocks.MockDeleteChatById, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockGetAllChatsByOwnerId := mocks.NewMockGetAllChatsByOwnerId(ctrl)
	mockDeleteChatById := mocks.NewMockDeleteChatById(ctrl)

	sut := controllers.NewDeleteChatController(mockGetAllChatsByOwnerId, mockDeleteChatById)

	return sut, mockGetAllChatsByOwnerId, mockDeleteChatById, ctrl
}

func createDeleteChatHttpRequest() protocols.HttpRequest {
	header := http.Header{}
	header.Add("UserId", "fake-owner-id")

	urlParams := url.Values{}
	urlParams.Set("id", "fake-chat-id")

	return protocols.HttpRequest{
		Body:      nil,
		Header:    header,
		UrlParams: urlParams,
	}
}

func TestDeleteChatController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, getAllChatsByOwnerId, deleteChatById, ctrl := setupDeleteChatControllerMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdResponse := &usecase.GetAllChatsByOwnerIdOutput{
			Id:        "fake-chat-id",
			CreatedAt: "fake-date",
			OwnerId:   "fake-owner-id",
		}

		var getAllChatsByOwnerIdResponseSlice []*usecase.GetAllChatsByOwnerIdOutput
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)

		getAllChatsByOwnerId.EXPECT().Get("fake-owner-id").Return(getAllChatsByOwnerIdResponseSlice, nil)
		deleteChatById.EXPECT().Delete("fake-chat-id").Return(nil)

		res := sut.Handle(createDeleteChatHttpRequest())

		var responseBody controllers.CreateChatControllerResponse
		err := json.NewDecoder(res.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctDeleteChatControllerResponse := &controllers.CreateChatControllerResponse{}
		require.Equal(t, &responseBody, correctDeleteChatControllerResponse)
		require.Equal(t, res.StatusCode, http.StatusNoContent)
	})

	t.Run("GetAllChatsByOwnerIdError", func(t *testing.T) {
		sut, getAllChatsByOwnerId, _, ctrl := setupDeleteChatControllerMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerId.EXPECT().Get("fake-owner-id").Return(nil, errors.New("fake-error"))

		res := sut.Handle(createDeleteChatHttpRequest())

		verifyHttpResponse(t, res, http.StatusInternalServerError, "an error ocurred while getting chats")
	})

	t.Run("IsNotCorrectChatIdError", func(t *testing.T) {
		sut, getAllChatsByOwnerId, _, ctrl := setupDeleteChatControllerMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdResponse := &usecase.GetAllChatsByOwnerIdOutput{
			Id:        "other-fake-chat-id",
			CreatedAt: "fake-date",
			OwnerId:   "fake-id",
		}

		var getAllChatsByOwnerIdResponseSlice []*usecase.GetAllChatsByOwnerIdOutput
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)

		getAllChatsByOwnerId.EXPECT().Get("fake-owner-id").Return(getAllChatsByOwnerIdResponseSlice, nil)

		res := sut.Handle(createDeleteChatHttpRequest())

		verifyHttpResponse(t, res, http.StatusForbidden, "you do not have this chat")
	})

	t.Run("Success", func(t *testing.T) {
		sut, getAllChatsByOwnerId, deleteChatById, ctrl := setupDeleteChatControllerMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerIdResponse := &usecase.GetAllChatsByOwnerIdOutput{
			Id:        "fake-chat-id",
			CreatedAt: "fake-date",
			OwnerId:   "fake-owner-id",
		}

		var getAllChatsByOwnerIdResponseSlice []*usecase.GetAllChatsByOwnerIdOutput
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)
		getAllChatsByOwnerIdResponseSlice = append(getAllChatsByOwnerIdResponseSlice, getAllChatsByOwnerIdResponse)

		getAllChatsByOwnerId.EXPECT().Get("fake-owner-id").Return(getAllChatsByOwnerIdResponseSlice, nil)
		deleteChatById.EXPECT().Delete("fake-chat-id").Return(errors.New("fake-error"))

		res := sut.Handle(createDeleteChatHttpRequest())

		verifyHttpResponse(t, res, http.StatusInternalServerError, "an error ocurred while deleting chat")
	})
}
