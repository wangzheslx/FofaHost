package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"myfofa/internal/biz"
	"myfofa/internal/conf"
	"myfofa/internal/service"
	mlog "myfofa/log"
)

var (
	//从配置当中读取
	config = &conf.UserInfo{
		Email: "shilixiang@baimaohui.net",
		Key:   "90a961e5ab880b697e352b32d75eb883",
	}
)

// 依赖注入
func wireService() *service.FofaUsecase {
	logger := mlog.NewLogger()
	fofaRepo := biz.NewFofaRepo(config, logger)
	fofaUsecase := service.NewFofaUsecase(fofaRepo, logger)
	return fofaUsecase
}

func main() {
	app := wireService()
	//返回结构体
	resp, err := app.Host(context.Background(), &service.HostReq{
		Ip: "159.75.231.54",
	})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp)
	//格式化输出
	bs, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("student=%v\n", out.String())

}
