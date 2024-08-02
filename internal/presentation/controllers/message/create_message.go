package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type CreateMessageController struct {
	GetAllChatsByOwnerId usecase.GetAllChatsByOwnerId
}

func NewCreateMessageController(getAllChatsByOwnerId usecase.GetAllChatsByOwnerId) *CreateMessageController {
	return &CreateMessageController{
		GetAllChatsByOwnerId: getAllChatsByOwnerId,
	}
}

type CreateMessageControllerResponse struct {
	Id         string `json:"id"`
	ChatId     string `json:"chatId"`
	Message    string `json:"message"`
	AuthorName string `json:"authorName"`
	AuthorId   string `json:"authorId"`
}

type CreateMessageControllerBody struct {
	ChatId     string `json:"chatId"`
	Message    string `json:"message"`
	AuthorName string `json:"authorName"`
	AuthorId   string `json:"authorId"`
}

func (c *CreateMessageController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	var body CreateMessageControllerBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid body request",
		}, http.StatusBadRequest)
	}

	ownerId := r.Header.Get("UserId")

	chats, err := c.GetAllChatsByOwnerId.Get(ownerId)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while getting chats",
		}, http.StatusInternalServerError)
	}

	if !isThereOwnerChat(chats, body.ChatId) {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "you do not have this chat",
		}, http.StatusForbidden)
	}

	return nil
}

func isThereOwnerChat(chats []*usecase.GetAllChatsByOwnerIdOutput, chatId string) bool {
	for _, chat := range chats {
		if chat.Id == chatId {
			return true
		}
	}

	return false
}
