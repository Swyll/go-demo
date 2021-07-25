package server

import (
	"go-demo/internal/final/bff/acount"
	acountpb "go-demo/internal/final/bff/acount/proto"
	"go-demo/internal/final/bff/comment"
	commentpb "go-demo/internal/final/bff/comment/proto"
)

type BffServer struct {
	acount  *acount.AcountModule
	comment *comment.CommentModule
}

type Info struct {
	AcountInfo  *acountpb.QueryResp  `json:"acount_info"`
	CommentInfo *commentpb.QueryResp `json:"comment_info"`
}

func NewBffServer(aclient acountpb.AcountSerClient, cclient commentpb.CommentSerClient) *BffServer {
	return &BffServer{
		acount:  acount.NewAcountModule(aclient),
		comment: comment.NewCommentModule(cclient),
	}
}

func (bs *BffServer) QueryAcountInfo(acountID string) (*Info, error) {
	resp, err := bs.acount.QueryAcount(&acountpb.Queryreq{
		AcountID: acountID,
		Limit:    0,
		Offset:   1,
	})
	if err != nil {
		return nil, err
	}

	resp2, err := bs.comment.QueryComment(&commentpb.Queryreq{
		AcountID: acountID,
		Limit:    0,
		Offset:   1,
	})
	if err != nil {
		return nil, err
	}

	info := new(Info)
	info.AcountInfo = resp
	info.CommentInfo = resp2

	return info, nil
}
