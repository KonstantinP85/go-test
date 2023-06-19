package service

import (
	"math/rand"
    "time"    

	s_go_project "github.com/KonstantinP85/s-go-service"
)

type PostService struct {

}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) PostChangeEvent(postEvent s_go_project.PostEvent) (int, error) {

	var interval int = 3 + rand.Intn(12)

	time.Sleep(time.Duration(interval) * time.Second)

	var postId int = postEvent.PostId
	if postId%3 == 0 {
		return 500, nil
	}

	return 200, nil
}
