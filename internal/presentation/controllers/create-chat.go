package controllers

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type CreateChatController struct{}

func NewCreateChatController() *CreateChatController {
	return &CreateChatController{}
}

type CreateChatControllerResponse struct{}

func (c *CreateChatController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	return helpers.CreateResponse(&CreateChatControllerResponse{}, http.StatusOK)
}
