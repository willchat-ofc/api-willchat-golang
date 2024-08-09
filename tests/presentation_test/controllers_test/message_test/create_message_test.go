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
	controllers "github.com/willchat-ofc/api-willchat-golang/internal/presentation/controllers/message"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
	"github.com/willchat-ofc/api-willchat-golang/tests/mocks"
)

func setupCreateMessageMocks(t *testing.T) (*controllers.CreateMessageController, *mocks.MockFindChatById, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mockFindChatById := mocks.NewMockFindChatById(ctrl)

	sut := controllers.NewCreateMessageController(mockFindChatById)

	return sut, mockFindChatById, ctrl
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
		signUpController, _, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		httpRequest := &protocols.HttpRequest{
			Body:   io.NopCloser(strings.NewReader("{invalid json")),
			Header: nil,
		}

		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid body request")
	})

	t.Run("ChatNotFound", func(t *testing.T) {
		signUpController, mockFindChatById, ctrl := setupCreateMessageMocks(t)
		defer ctrl.Finish()

		mockFindChatById.EXPECT().Find("fake-chat-id").Return(nil, errors.New("chat not found"))

		httpRequest := createCreateMessageHttpRequest(t)
		httpResponse := signUpController.Handle(httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusNotFound, "chat not found")
	})
}
