package service

import (
	"context"
	"myfofa/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// 接口层可扩展
type FofaRepo interface {
	Host(ctx context.Context, req *HostReq) (*HostResp, error)
}

type FofaUsecase struct {
	log  *log.Helper
	repo *biz.FofaRepo
}

func NewFofaUsecase(repo *biz.FofaRepo, logger log.Logger) *FofaUsecase {
	return &FofaUsecase{
		log:  log.NewHelper(log.With(logger, "module", "usecase/fofa")),
		repo: repo,
	}
}

func (uc *FofaUsecase) Host(ctx context.Context, req *HostReq) (*HostResp, error) {
	resp, err := uc.repo.Host(ctx, &biz.HostReq{
		Ip: req.Ip,
	})
	if err != nil {
		uc.log.WithContext(ctx).Error("Repo: %+v", err)
		return nil, err
	}
	return &HostResp{
		Host:        resp.Host,
		IP:          resp.IP,
		Asn:         resp.Asn,
		Org:         resp.Org,
		CountryName: resp.CountryName,
		CountryCode: resp.CountryCode,
		Protocol:    resp.Protocol,
		Port:        resp.Port,
		Category:    resp.Category,
		Product:     resp.Product,
		UpdateTime:  resp.UpdateTime,
	}, nil
}
