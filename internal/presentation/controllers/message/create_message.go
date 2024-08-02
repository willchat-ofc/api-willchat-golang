package controllers

import (
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
	ownerId := r.Header.Get("UserId")

	_, err := c.GetAllChatsByOwnerId.Get(ownerId)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while getting chats",
		}, http.StatusInternalServerError)
	}

	return nil
}
