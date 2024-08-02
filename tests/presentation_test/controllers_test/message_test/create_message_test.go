package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/message"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupCreateMessageMocks(t *testing.T) (*controllers.CreateMessageController, *mocks.MockGetAllChatsByOwnerId, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockGetAllChatsByOwnerId := mocks.NewMockGetAllChatsByOwnerId(ctrl)
	sut := controllers.NewCreateMessageController(mockGetAllChatsByOwnerId)

	return sut, mockGetAllChatsByOwnerId, ctrl
}

func createCreateMessageHttpRequest(t *testing.T) protocols.HttpRequest {
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&controllers.CreateMessageControllerBody{
		ChatId:     "fake-chat-id",
		Message:    "fake-message",
		AuthorName: "fake-author-name",
		AuthorId:   "fake-author-id",
	})
	require.NoError(t, err)

	header := http.Header{}
	header.Add("UserId", "fake-user-id")

	return protocols.HttpRequest{
		Body:   io.NopCloser(&requestBody),
		Header: header,
	}
}

func TestCreateMessageController(t *testing.T) {
	fakeUserId := "fake-user-id"
	t.Run("GetAllChatsByOwnerId Error", func(t *testing.T) {
		sut, getAllChatsByOwnerId, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerId.EXPECT().Get(fakeUserId).Return(nil, errors.New("fake-error"))

		res := sut.Handle(createCreateMessageHttpRequest(t))

		verifyHttpResponse(t, res, http.StatusInternalServerError, "an error ocurred while getting chats")
	})

	t.Run("InvalidBodyRequest", func(t *testing.T) {
		signUpController, _, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		httpRequest := &protocols.HttpRequest{
			Body:   io.NopCloser(strings.NewReader("{invalid json")),
			Header: nil,
		}

		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid body request")
	})

	t.Run("OwnerDoNotHaveThisChat", func(t *testing.T) {
		sut, getAllChatsByOwnerId, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		getAllChatsByOwnerId.EXPECT().Get(fakeUserId).Return([]*usecase.GetAllChatsByOwnerIdOutput{{
			Id:        "other-49e44949-cdd4-4a80-960b-297905c5d514",
			CreatedAt: "fake-date",
		}}, nil)

		res := sut.Handle(createCreateMessageHttpRequest(t))

		verifyHttpResponse(t, res, http.StatusForbidden, "you do not have this chat")
	})
}
