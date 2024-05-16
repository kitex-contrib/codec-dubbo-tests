package kitex2kitex

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"
	"testing"

	"github.com/cloudwego/kitex/client"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"

	"github.com/cloudwego/kitex/pkg/remote/trans/detection"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	hello "github.com/kitex-contrib/codec-dubbo-tests/code/kitex_sample"
	khello "github.com/kitex-contrib/codec-dubbo-tests/code/kitex_sample/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex_sample/kitex_gen/hello/greetservice"
)

var nativeCli greetservice.Client
var dubboCli greetservice.Client

// initHelloClient inits Kitex client with specified destService and hostPort
func initHelloClient(destService, hostPort string) {
	var err error
	dubboCli, err = greetservice.NewClient(destService,
		client.WithHostPorts(hostPort),
		client.WithCodec(
			dubbo.NewDubboCodec(
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
			),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("Kitex client initialized failed, err :%s", err))
	}

	nativeCli, err = greetservice.NewClient(destService,
		client.WithHostPorts(hostPort),
	)
	if err != nil {
		panic(fmt.Sprintf("Kitex client initialized failed, err :%s", err))
	}
}

// use exitCh to receive exiting signal.
func runHelloKitexServer(startCh chan struct{}, exitCh chan error, addr string) {
	netAddr, _ := net.ResolveTCPAddr("tcp", addr)
	svr := greetservice.NewServer(new(hello.GreetServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "hello",
		}),
		server.WithServiceAddr(netAddr),
		server.WithTransHandlerFactory(detection.NewSvrTransHandlerFactory(netpoll.NewSvrTransHandlerFactory(),
			dubbo.NewSvrTransHandlerFactory(
				// set the Java interface name corresponding to the Kitex service
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider")),
			nphttp2.NewSvrTransHandlerFactory(),
		)),
		server.WithExitSignal(func() <-chan error {
			return exitCh
		}),
	)
	once := sync.Once{}
	server.RegisterStartHook(func() {
		once.Do(func() {
			close(startCh)
		})
	})

	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestProtocolProbing(t *testing.T) {
	req := "world"
	ctx := context.Background()
	wantResp := "hello, world"
	resp, err := nativeCli.Greet(ctx, req)
	assertResp(t, err, resp, wantResp)

	resp1, err := nativeCli.GreetWithStruct(ctx, &khello.GreetRequest{Req: req})
	assertResp(t, err, resp1.Resp, wantResp)

	resp, err = dubboCli.Greet(ctx, req)
	assertResp(t, err, resp, wantResp)

	resp1, err = dubboCli.GreetWithStruct(ctx, &khello.GreetRequest{Req: req})
	assertResp(t, err, resp1.Resp, wantResp)
}

func assertResp(t *testing.T, err error, req, resp interface{}) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(req, resp) {
		t.Fatalf("req is not equal to resp, req: %v, resp: %v", req, resp)
	}
}
