package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type DeleteChatController struct {
	GetAllChatsByOwnerId usecase.GetAllChatsByOwnerId
	DeleteChatById       usecase.DeleteChatById
}

func NewDeleteChatController(getAllChatsByOwnerId usecase.GetAllChatsByOwnerId, deleteChatById usecase.DeleteChatById) *DeleteChatController {
	return &DeleteChatController{
		GetAllChatsByOwnerId: getAllChatsByOwnerId,
		DeleteChatById:       deleteChatById,
	}
}

func (c *DeleteChatController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	ownerId := r.Header.Get("UserId")

	chats, err := c.GetAllChatsByOwnerId.Get(ownerId)

	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while getting chats",
		}, http.StatusInternalServerError)
	}

	chatId := r.UrlParams.Get("id")
	if _, err := uuid.Parse(chatId); err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid chat id format",
		}, http.StatusBadRequest)
	}

	isCorrectChat := false
	for _, chat := range chats {
		if chat.Id == chatId {
			isCorrectChat = true
			break
		}
	}

	if !isCorrectChat {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "you do not have this chat",
		}, http.StatusForbidden)
	}

	err = c.DeleteChatById.Delete(chatId)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while deleting chat",
		}, http.StatusInternalServerError)
	}

	return helpers.CreateResponse(nil, http.StatusNoContent)
}
