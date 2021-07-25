package handler

import (
	"context"
	"go-demo/internal/final/comment/define"
	commentpb "go-demo/internal/final/comment/proto/comment"
	"go-demo/internal/final/comment/server"
)

type PBServer struct {
	s *server.CommentServer
}

func NewPBServer(s *server.CommentServer) *PBServer {
	return &PBServer{
		s: s,
	}
}

func (ps *PBServer) DeleteComment(ctx context.Context, req *commentpb.DeleteReq) (*commentpb.DeleteResp, error) {
	err := ps.s.DeleteComment(req.ID)
	if err != nil {
		return &commentpb.DeleteResp{}, err
	}

	return &commentpb.DeleteResp{}, nil
}

func (ps *PBServer) AddComment(ctx context.Context, req *commentpb.AddReq) (*commentpb.AddResp, error) {
	return &commentpb.AddResp{}, ps.s.AddComment(&define.CommentData{
		ID:       req.Comment.CommentID,
		AcountID: req.Comment.AcountID,
		Comment:  req.Comment.CommentData,
	})
}

func (ps *PBServer) QueryComment(ctx context.Context, req *commentpb.Queryreq) (*commentpb.QueryResp, error) {
	datas, err := ps.s.QueryComment(req.CommentID, int(req.Limit), int(req.Offset))
	if err != nil {
		return &commentpb.QueryResp{}, err
	}

	resp := &commentpb.QueryResp{
		Comments: make([]*commentpb.Comment, 0, len(datas)),
	}
	for _, data := range datas {
		comment := &commentpb.Comment{
			CommentID:   data.ID,
			AcountID:    data.AcountID,
			CommentData: data.Comment,
		}

		resp.Comments = append(resp.Comments, comment)
	}

	return resp, nil
}
