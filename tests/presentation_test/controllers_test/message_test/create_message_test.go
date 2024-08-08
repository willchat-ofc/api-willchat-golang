package controllers_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
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

// func createCreateMessageHttpRequest(t *testing.T) protocols.HttpRequest {
// 	var requestBody bytes.Buffer
// 	err := json.NewEncoder(&requestBody).Encode(&controllers.CreateMessageControllerBody{
// 		ChatId:     "fake-chat-id",
// 		Message:    "fake-message",
// 		AuthorName: "fake-author-name",
// 		AuthorId:   "fake-author-id",
// 	})
// 	require.NoError(t, err)

// 	header := http.Header{}
// 	header.Add("UserId", "fake-user-id")

// 	return protocols.HttpRequest{
// 		Body:   io.NopCloser(&requestBody),
// 		Header: header,
// 	}
// }

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
}
