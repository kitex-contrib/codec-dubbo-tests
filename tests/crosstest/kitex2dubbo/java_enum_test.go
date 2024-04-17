package kitex2dubbo

import (
	"context"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
	"testing"
)

func TestEchoJavaEnum_Java(t *testing.T) {
	req := "1"
	resp, err := cli2Java.EchoJavaEnum(context.Background(), echo.KitexEnum_ONE)
	assertEcho(t, err, req, resp)
}
