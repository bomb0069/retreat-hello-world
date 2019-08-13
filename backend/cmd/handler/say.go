package handler

import (
	"say/internal/platform/repository"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

type Resource struct {
	Repository repository.SayRepository
}

func (resource Resource) SayHandler(context *gin.Context) {
	say := resource.Repository.GetFirstMessage()
	message := Message{}
	message.translate(say.Message)
	context.JSON(200, message)
}

func (message *Message) translate(text string) {
	message.Message = "สวัสดีชาวโลก"
}
