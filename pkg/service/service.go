package service

import "github.com/hdget/hdsdk/lib/svc"

type serviceImpl struct {
	svc.DaprService
}

var _ svc.DaprService = (*serviceImpl)(nil)

func New() (svc.Service, error) {
	return &serviceImpl{
		DaprService: svc.NewDaprService(),
	}, nil
}
