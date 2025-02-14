package service

import (
	"context"
	"myfofa/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// 接口层可扩展
type FofaRepo interface {
	Host(ctx context.Context, req *HostReq) (*HostResp, error)
	HostDtl(ctx context.Context, req *HostDtlReq) (*HostDtlResp, error)
	//Query(ctx context.Context, req *QueryReq) (*QueryResp, error)
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

func (uc *FofaUsecase) HostDtl(ctx context.Context, req *HostDtlReq) (*HostDtlResp, error) {
	resp, err := uc.repo.HostDtl(ctx, &biz.HostDtlReq{
		Ip: req.Ip,
	})
	if err != nil {
		uc.log.WithContext(ctx).Error("Repo: %+v", err)
	}

	var mport []Ports
	for _, port := range resp.Ports {
		var products []Products
		for _, product := range port.Products {
			products = append(products, Products{
				Product:      product.Product,
				Category:     product.Category,
				Level:        product.Level,
				SortHardCode: product.SortHardCode,
				Company:      product.Company,
			})
		}
		mport = append(mport, Ports{
			Port:       port.Port,
			UpdateTime: port.UpdateTime,
			Protocol:   port.Protocol,
			Products:   products,
		})
	}

	return &HostDtlResp{
		Error:       resp.Error,
		Host:        resp.Host,
		IP:          resp.IP,
		Asn:         resp.Asn,
		Org:         resp.Org,
		CountryName: resp.CountryName,
		CountryCode: resp.CountryCode,
		Domain:      resp.Domain,
		Ports:       mport,
		UpdateTime:  resp.UpdateTime,
	}, nil
}
