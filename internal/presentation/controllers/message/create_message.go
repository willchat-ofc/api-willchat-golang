package controllers

import presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"

type CreateMessageController struct{}

func NewCreateMessageController() *CreateMessageController {
	return &CreateMessageController{}
}

func (c *CreateMessageController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	return nil
}
