package server

import (
	"go-demo/internal/final/comment/dao"
	"go-demo/internal/final/comment/define"
)

type CommentServer struct {
	dataengine dao.CommentEngine
}

func NewCommentServer(e dao.CommentEngine) *CommentServer {
	return &CommentServer{
		dataengine: e,
	}
}

func (s *CommentServer) DeleteComment(id string) error {
	return s.dataengine.DeleteComment(id)
}

func (s *CommentServer) AddComment(data *define.CommentData) error {
	return s.dataengine.AddComment(data)
}

func (s *CommentServer) QueryComment(acountID string, limit, offset int) ([]*define.CommentData, error) {
	return s.dataengine.QueryComment(acountID, limit, offset)
}
