package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type CreateMessageController struct {
	FindChatById  usecase.FindChatById
	CreateMessage usecase.CreateMessage
}

func NewCreateMessageController(findChatById usecase.FindChatById, createMessage usecase.CreateMessage) *CreateMessageController {
	return &CreateMessageController{
		FindChatById:  findChatById,
		CreateMessage: createMessage,
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

	_, err = c.FindChatById.Find(body.ChatId)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "chat not found",
		}, http.StatusNotFound)
	}

	_, err = c.CreateMessage.Create(&usecase.CreateMessageInput{
		ChatId:     body.ChatId,
		Message:    body.Message,
		AuthorName: body.AuthorName,
		AuthorId:   body.AuthorId,
	})
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error occurred while creating message",
		}, http.StatusInternalServerError)
	}

	return nil
}
