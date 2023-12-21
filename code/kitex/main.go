package main

import (
	"log"
	"net"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"

	"github.com/cloudwego/kitex/server"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo/testservice"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := testservice.NewServer(new(handler.TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
			dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi()),
		)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
