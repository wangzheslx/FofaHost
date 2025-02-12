package biz

import (
	"context"
	"encoding/json"
	"myfofa/internal/conf"

	"github.com/buger/jsonparser"
	"github.com/fofapro/fofa-go/fofa"
	"github.com/go-kratos/kratos/v2/log"
)

type FofaRepo struct {
	key    string
	Client *fofa.Fofa
	log    *log.Helper
}

const (
	HostAPIUrl = "https://fofa.info/api/v1/host"
)

func NewFofaRepo(ui *conf.UserInfo, logger log.Logger) *FofaRepo {
	return &FofaRepo{

		key:    ui.Key,
		Client: fofa.NewFofaClient([]byte(ui.Email), []byte(ui.Key)),
		log:    log.NewHelper(logger),
	}
}
func (f *FofaRepo) Host(ctx context.Context, req *HostReq) (resp *HostResp, err error) {
	f.log.WithContext(ctx).Info("Repo key : %+v", f.key)
	q := HostAPIUrl + "/" + req.Ip + "?key=" + f.key
	content, err := f.Client.Get(q)
	if err != nil {
		f.log.WithContext(ctx).Error("RepoGet: %+v", err)
	}
	errmsg, err := jsonparser.GetString(content, "errmsg")
	if err == nil {
		f.log.WithContext(ctx).Error("Repomsg: %+v", errmsg)
		return nil, err
	} else {
		err = nil
	}
	resp = &HostResp{}
	err = json.Unmarshal(content, resp)
	if err != nil {
		f.log.WithContext(ctx).Error("Repo: %+v", err)
	}
	return resp, nil
}
