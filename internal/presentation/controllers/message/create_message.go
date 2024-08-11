package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type CreateMessageController struct {
	FindChatById  usecase.FindChatById
	CreateMessage usecase.CreateMessage
	Validate      *validator.Validate
}

func NewCreateMessageController(findChatById usecase.FindChatById, createMessage usecase.CreateMessage) *CreateMessageController {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &CreateMessageController{
		FindChatById:  findChatById,
		CreateMessage: createMessage,
		Validate:      validate,
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
	ChatId     string `validate:"uuid,required"`
	Message    string `validate:"min=1,max=400,required"`
	AuthorName string `validate:"min=4,max=50,required"`
	AuthorId   string
}

func (c *CreateMessageController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	var body CreateMessageControllerBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid body request",
		}, http.StatusBadRequest)
	}

	if err := c.Validate.Struct(body); err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: err.Error(),
		}, http.StatusUnprocessableEntity)
	}

	if _, err := c.FindChatById.Find(body.ChatId); err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "chat not found",
		}, http.StatusNotFound)
	}

	message, err := c.CreateMessage.Create(&usecase.CreateMessageInput{
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

	return helpers.CreateResponse(&CreateMessageControllerResponse{
		Id:         message.Id,
		ChatId:     message.ChatId,
		Message:    message.Message,
		AuthorName: message.AuthorName,
		AuthorId:   message.AuthorId,
	}, http.StatusCreated)
}
