package controllers

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type GetAllChatsByOwnerIdController struct {
	GetAllChatsByOwnerId usecase.GetAllChatsByOwnerId
}

type GetAllChatsByOwnerIdControllerResponse struct {
	Chats []*usecase.GetAllChatsByOwnerIdOutput `json:"chats"`
}

func NewGetAllChatsByOwnerIdController(getAllChatsByOwnerId usecase.GetAllChatsByOwnerId) *GetAllChatsByOwnerIdController {
	return &GetAllChatsByOwnerIdController{
		GetAllChatsByOwnerId: getAllChatsByOwnerId,
	}
}

func (c *GetAllChatsByOwnerIdController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	ownerId := r.Header.Get("UserId")

	chats, err := c.GetAllChatsByOwnerId.Get(ownerId)

	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while finding chats",
		}, http.StatusInternalServerError)
	}

	return helpers.CreateResponse(&GetAllChatsByOwnerIdControllerResponse{
		Chats: chats,
	}, http.StatusOK)
}
