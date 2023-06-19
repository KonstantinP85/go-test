package service

import (
	s_go_project "github.com/KonstantinP85/s-go-service"
)


type Post interface {
	PostChangeEvent(postEvent s_go_project.PostEvent) (int, error)
}

type Service struct {
	Post
}

func NewService() *Service {
	return &Service{
		Post: NewPostService(),
	}
}