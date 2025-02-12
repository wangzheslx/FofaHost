package biz

import (
	"context"
	"fmt"
	"myfofa/internal/conf"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fofapro/fofa-go/fofa"
	"github.com/go-kratos/kratos/v2/log"
)

func TestFofaRepo_Host(t *testing.T) {
	// 模拟配置信息
	ui := &conf.UserInfo{
		Email: "shilixiang@baimaohui.net",
		Key:   "90a961e5ab880b697e352b32d75eb883",
	}

	// 创建 FofaRepo 实例
	repo := NewFofaRepo(ui, log.NewStdLogger(os.Stdout))

	// 模拟请求上下文
	ctx := context.Background()

	// 模拟请求 IP
	req := &HostReq{
		Ip: "159.75.231.54",
	}

	// 模拟成功的响应内容
	//successContent := `{"host": "127.0.0.1", "port": 80}`

	//模拟服务器返回成功响应
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	// 设置 FofaRepo 的客户端为模拟服务器的地址
	repo.Client = fofa.NewFofaClient(nil, nil)

	// 执行 Host 方法
	resp, err := repo.Host(ctx, req)
	if err != nil {
		fmt.Print(err)
		return
	}
	// 断言没有错误发生
	fmt.Printf("%+v", resp)

}
