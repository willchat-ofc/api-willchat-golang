package controllers

import (
	"net/http"
	"strconv"
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

	limitStr := r.Header.Get("limit")
	if limitStr == "" {
		limitStr = "0"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "cannot parse limit",
		}, http.StatusBadRequest)
	}

	offsetStr := r.Header.Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "cannot parse offset",
		}, http.StatusBadRequest)
	}

	input := &usecase.FindMessagesByChatIdInput{
		ChatId: id,
		Limit:  limit,
		Offset: offset,
	}
	chats, err := c.FindMessagesByChatId.Find(input)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "chat not found",
		}, http.StatusInternalServerError)
	}

	return helpers.CreateResponse(chats, http.StatusOK)
}
