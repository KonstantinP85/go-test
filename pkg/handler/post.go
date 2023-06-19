package handler

import (
	"encoding/json"
	"net/http"
	"time" 
	"log"

	"github.com/gin-gonic/gin"

	s_go_project "github.com/KonstantinP85/s-go-service"
)


func (h *Handler) postEvent(c *gin.Context) {
	var input s_go_project.PostEvent

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"status": "error", "message":"invalid request body"})
		return
	}

	start := time.Now()
	code, err := h.services.Post.PostChangeEvent(input)
	log.Println("Response time:", time.Since(start))
	log.Println("Response code:", code)

	res, _ := json.Marshal(s_go_project.PostEvent{
		AuthorId  : input.AuthorId,
		PostId    : input.PostId,
		PostTitle : input.PostTitle,
		Action    : input.Action,
	})

	log.Println("Request body:", string(res))

	if code == 500 || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}