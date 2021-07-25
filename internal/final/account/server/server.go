package server

import (
	"go-demo/internal/final/account/dao"
	"go-demo/internal/final/account/define"
)

type AcountServer struct {
	dataengine dao.AcountEngine
}

func NewAcountServer(e dao.AcountEngine) *AcountServer {
	return &AcountServer{
		dataengine: e,
	}
}

func (s *AcountServer) DeleteAcount(id string) error {
	return s.dataengine.DeleteAcount(id)
}

func (s *AcountServer) AddAcount(data *define.AcountData) error {
	return s.dataengine.AddAcount(data)
}

func (s *AcountServer) QueryAcount(acountID string, limit, offset int) ([]*define.AcountData, error) {
	return s.dataengine.QueryAcount(acountID, limit, offset)
}
