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

func setupCreateChatMocks(t *testing.T) (*controllers.CreateChatController, *mocks.MockCreateChat, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockCreateChat := mocks.NewMockCreateChat(ctrl)

	sut := controllers.NewCreateChatController(mockCreateChat)

	return sut, mockCreateChat, ctrl
}

func createCreateChatHttpRequest() protocols.HttpRequest {
	header := http.Header{}
	header.Add("UserId", "fake-user-id")

	return protocols.HttpRequest{
		Body:   nil,
		Header: header,
	}
}

func TestCreateChatController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, createChat, ctrl := setupCreateChatMocks(t)
		defer ctrl.Finish()

		chatData := &usecase.CreateChatOutput{
			Id:        "fake-chat-id",
			OwnerId:   "fake-user-id",
			CreatedAt: "fake-created-at",
		}

		createChat.EXPECT().Create("fake-user-id").Return(chatData, nil)

		res := sut.Handle(createCreateChatHttpRequest())

		require.Equal(t, res.StatusCode, http.StatusOK)
		var responseBody controllers.CreateChatControllerResponse
		err := json.NewDecoder(res.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctSignInControllerResponse := &controllers.CreateChatControllerResponse{
			Id:        "fake-chat-id",
			OwnerId:   "fake-user-id",
			CreatedAt: "fake-created-at",
		}
		require.Equal(t, &responseBody, correctSignInControllerResponse)
	})

	t.Run("CreateChatError", func(t *testing.T) {
		sut, createChat, ctrl := setupCreateChatMocks(t)
		defer ctrl.Finish()

		createChat.EXPECT().Create("fake-user-id").Return(nil, errors.New("fake-error"))

		res := sut.Handle(createCreateChatHttpRequest())

		verifyHttpResponse(t, res, http.StatusInternalServerError, "an error ocurred when creating chat")
	})

}
