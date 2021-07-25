package dao

import "go-demo/internal/final/account/define"

type AcountEngine interface {
	DeleteAcount(id string) error
	AddAcount(*define.AcountData) error
	QueryAcount(acountID string, limit, offset int) ([]*define.AcountData, error)
}
