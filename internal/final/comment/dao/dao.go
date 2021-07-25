package dao

import "go-demo/internal/final/comment/define"

type CommentEngine interface {
	DeleteComment(id string) error
	AddComment(*define.CommentData) error
	QueryComment(acountID string, limit, offset int) ([]*define.CommentData, error)
}
