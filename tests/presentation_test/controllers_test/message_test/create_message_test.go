package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/models"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/message"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupCreateMessageMocks(t *testing.T) (*controllers.CreateMessageController, *mocks.MockFindChatById, *mocks.MockCreateMessage, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockFindChatById := mocks.NewMockFindChatById(ctrl)
	mockCreateMessage := mocks.NewMockCreateMessage(ctrl)

	sut := controllers.NewCreateMessageController(mockFindChatById, mockCreateMessage)

	return sut, mockFindChatById, mockCreateMessage, ctrl
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
	t.Run("InvalidBodyRequest", func(t *testing.T) {
		signUpController, _, _, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		httpRequest := &protocols.HttpRequest{
			Body:   io.NopCloser(strings.NewReader("{invalid json")),
			Header: nil,
		}

		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid body request")
	})

	t.Run("ChatNotFound", func(t *testing.T) {
		signUpController, mockFindChatById, _, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		mockFindChatById.EXPECT().Find("fake-chat-id").Return(nil, errors.New("chat not found"))

		httpRequest := createCreateMessageHttpRequest(t)
		httpResponse := signUpController.Handle(httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusNotFound, "chat not found")
	})

	t.Run("ErrorWhileCreatingMessage", func(t *testing.T) {
		signUpController, mockFindChatById, mockCreateMessage, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		fakeChat := &models.Chat{
			Id:        "fake-chat-id",
			CreatedAt: time.Now(),
			OwnerId:   "fake-owner-id",
		}
		mockFindChatById.EXPECT().Find("fake-chat-id").Return(fakeChat, nil)

		fakeCreateMessageInput := &usecase.CreateMessageInput{
			ChatId:     "fake-chat-id",
			Message:    "fake-message",
			AuthorName: "fake-author-name",
			AuthorId:   "fake-author-id",
		}

		mockCreateMessage.EXPECT().Create(fakeCreateMessageInput).Return(nil, errors.New("fake-error"))

		httpRequest := createCreateMessageHttpRequest(t)
		httpResponse := signUpController.Handle(httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusInternalServerError, "an error occurred while creating message")
	})

	t.Run("Success", func(t *testing.T) {
		signUpController, mockFindChatById, mockCreateMessage, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		fakeChat := &models.Chat{
			Id:        "fake-chat-id",
			CreatedAt: time.Now(),
			OwnerId:   "fake-owner-id",
		}
		mockFindChatById.EXPECT().Find("fake-chat-id").Return(fakeChat, nil)

		fakeCreateMessageInput := &usecase.CreateMessageInput{
			ChatId:     "fake-chat-id",
			Message:    "fake-message",
			AuthorName: "fake-author-name",
			AuthorId:   "fake-author-id",
		}
		fakeMessage := &models.Message{
			Id:         "fake-message-id",
			ChatId:     "fake-chat-id",
			Message:    "fake-message",
			AuthorName: "fake-author",
			AuthorId:   "fake-author-id",
		}

		mockCreateMessage.EXPECT().Create(fakeCreateMessageInput).Return(fakeMessage, nil)

		httpRequest := createCreateMessageHttpRequest(t)
		res := signUpController.Handle(httpRequest)

		var responseBody controllers.CreateMessageControllerResponse
		err := json.NewDecoder(res.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctCreateMessageResponse := controllers.CreateMessageControllerResponse{
			Id:         "fake-message-id",
			ChatId:     "fake-chat-id",
			Message:    "fake-message",
			AuthorName: "fake-author",
			AuthorId:   "fake-author-id",
		}

		require.Equal(t, correctCreateMessageResponse, responseBody)
		require.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
