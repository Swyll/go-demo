package acount

import (
	"context"
	acountpb "go-demo/internal/final/bff/acount/proto"

	"github.com/pkg/errors"
)

type AcountModule struct {
	client acountpb.AcountSerClient
}

func NewAcountModule(client acountpb.AcountSerClient) *AcountModule {
	return &AcountModule{
		client: client,
	}
}

func (am *AcountModule) AddAcount(req *acountpb.AddReq) error {
	_, err := am.client.AddAcount(context.Background(), req)
	if err != nil {
		return errors.Wrapf(err, "req:%+v", req)
	}

	return nil
}

func (am *AcountModule) DeleteAcount(req *acountpb.DeleteReq) error {
	_, err := am.client.DeleteAcount(context.Background(), req)
	if err != nil {
		return errors.Wrapf(err, "req:%+v", req)
	}

	return nil
}

func (am *AcountModule) QueryAcount(req *acountpb.Queryreq) (*acountpb.QueryResp, error) {
	resp, err := am.client.QueryAcount(context.Background(), req)
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", req)
	}

	return resp, nil
}
