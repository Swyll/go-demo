package comment

import (
	"context"
	commentpb "go-demo/internal/final/bff/comment/proto"

	"github.com/pkg/errors"
)

type CommentModule struct {
	client commentpb.CommentSerClient
}

func NewCommentModule(client commentpb.CommentSerClient) *CommentModule {
	return &CommentModule{
		client: client,
	}
}

func (am *CommentModule) AddComment(req *commentpb.AddReq) error {
	_, err := am.client.AddComment(context.Background(), req)
	if err != nil {
		return errors.Wrapf(err, "req:%+v", req)
	}

	return nil
}

func (am *CommentModule) DeleteComment(req *commentpb.DeleteReq) error {
	_, err := am.client.DeleteComment(context.Background(), req)
	if err != nil {
		return errors.Wrapf(err, "req:%+v", req)
	}

	return nil
}

func (am *CommentModule) QueryComment(req *commentpb.Queryreq) (*commentpb.QueryResp, error) {
	resp, err := am.client.QueryComment(context.Background(), req)
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", req)
	}

	return resp, nil
}
