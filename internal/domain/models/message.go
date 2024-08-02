package models

type Message struct {
	Id         string `json:"id"`
	ChatId     string `json:"chatId"`
	Message    string `json:"message"`
	AuthorName string `json:"authorName"`
	AuthorId   string `json:"authorId"`
}
