package controllers

import (
	"net/http"
	"strings"

	"github.com/willchat-ofc/api-willchat-golang/internal/domain/usecase"
	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/helpers"
	presentationProtocols "github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

type FindMessagesByChatIdController struct {
	FindMessagesByChatId usecase.FindMessagesByChatId
}

func NewFindMessagesByChatIdController(findMessagesByChatId usecase.FindMessagesByChatId) *FindMessagesByChatIdController {
	return &FindMessagesByChatIdController{
		FindMessagesByChatId: findMessagesByChatId,
	}
}

func (c *FindMessagesByChatIdController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	id := strings.TrimPrefix(r.UrlPath, "/message/")

	chats, err := c.FindMessagesByChatId.Find(id)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "chat not found",
		}, http.StatusInternalServerError)
	}

	return helpers.CreateResponse(chats, http.StatusOK)
}
