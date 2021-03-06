package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	proto "github.com/micro/examples/service/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

