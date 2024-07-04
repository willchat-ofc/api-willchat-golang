package controllers

import presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"

type CreateChatController struct{}

func NewCreateChatController() *CreateChatController {
	return &CreateChatController{}
}

func (c *CreateChatController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	// var
}
