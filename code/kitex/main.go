package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/remote/trans/detection"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	"github.com/cloudwego/kitex/server"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo/testservice"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := testservice.NewServer(new(handler.TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithTransHandlerFactory(detection.NewSvrTransHandlerFactory(netpoll.NewSvrTransHandlerFactory(),
			dubbo.NewSvrTransHandlerFactory(
				dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
				dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi())),
			nphttp2.NewSvrTransHandlerFactory(),
		)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
