package handler

import (
	"context"
	"go-demo/internal/final/account/define"
	acountpb "go-demo/internal/final/account/proto/acount"
	"go-demo/internal/final/account/server"
)

type PBServer struct {
	s *server.AcountServer
}

func NewPBServer(s *server.AcountServer) *PBServer {
	return &PBServer{
		s: s,
	}
}

func (ps *PBServer) DeleteAcount(ctx context.Context, req *acountpb.DeleteReq) (*acountpb.DeleteResp, error) {
	err := ps.s.DeleteAcount(req.ID)
	if err != nil {
		return &acountpb.DeleteResp{}, err
	}

	return &acountpb.DeleteResp{}, nil
}

func (ps *PBServer) AddAcount(ctx context.Context, req *acountpb.AddReq) (*acountpb.AddResp, error) {
	return &acountpb.AddResp{}, ps.s.AddAcount(&define.AcountData{
		AcountID:  req.Acount.AcountID,
		Name:      req.Acount.Name,
		Age:       int(req.Acount.Age),
		EmailAddr: req.Acount.EmailAddr,
	})
}

func (ps *PBServer) QueryAcount(ctx context.Context, req *acountpb.Queryreq) (*acountpb.QueryResp, error) {
	datas, err := ps.s.QueryAcount(req.AcountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return &acountpb.QueryResp{}, err
	}

	resp := &acountpb.QueryResp{
		Acounts: make([]*acountpb.Acount, 0, len(datas)),
	}
	for _, data := range datas {
		acount := &acountpb.Acount{
			AcountID:  data.AcountID,
			Name:      data.Name,
			EmailAddr: data.EmailAddr,
			Age:       int64(data.Age),
		}

		resp.Acounts = append(resp.Acounts, acount)
	}

	return resp, nil
}
