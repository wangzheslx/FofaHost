package main

import (
	"encoding/json"
	"fmt"
	"myfofa/internal/biz"
	"myfofa/internal/conf"
	"myfofa/internal/service"
	mlog "myfofa/log"

	"github.com/fofapro/fofa-go/fofa"
)

var (
	//从配置当中读取
	config = &conf.UserInfo{
		Email: "shilixiang@baimaohui.net",
		Key:   "90a961e5ab880b697e352b32d75eb883",
	}
)


func wireService() *service.FofaUsecase {
	logger := mlog.NewLogger()
	fofaRepo := biz.NewFofaRepo(config, logger)
	fofaUsecase := service.NewFofaUsecase(fofaRepo, logger)
	return fofaUsecase
}

func main() {
	// app := wireService()
	// //返回结构体
	// resp, err := app.HostDtl(context.Background(), &service.HostDtlReq{
	// 	Ip: "159.75.231.54",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //fmt.Println(resp)
	// //格式化输出
	// bs, _ := json.Marshal(resp)
	// var out bytes.Buffer
	// json.Indent(&out, bs, "", "\t")
	// fmt.Printf("student=%v\n", out.String())

	email := config.Email
	key := config.Key

	clt := fofa.NewFofaClient([]byte(email), []byte(key))
	if clt == nil {
		fmt.Printf("create fofa client\n")
		return
	}
	ret, err := clt.QueryAsJSON(1, []byte(`body="小米"&&cert="true"&&status_code="200"`))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Printf("%s\n", ret)
	arr, err := clt.QueryAsArray(1, []byte(`domain="163.com"`), []byte("ip,host,title"))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Printf("count: %d\n", len(arr))
	encodeArr, _ := json.Marshal(arr)
	fmt.Printf("\n%s\n", encodeArr)

}
