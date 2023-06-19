package s_go_project

type PostEvent struct {
	AuthorId 	   int    `json:"authorId" binding:"required"`
	PostId         int    `json:"postId" binding:"required"`
	PostTitle      string `json:"postTitle" binding:"required"`
	Action         string `json:"action" binding:"required"`
}