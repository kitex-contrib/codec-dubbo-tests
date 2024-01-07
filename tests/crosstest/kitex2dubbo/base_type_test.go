/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kitex2dubbo

import (
	"context"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kitex-contrib/codec-dubbo-tests/code/dubbo-go/api"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo/testservice"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
)

var cli2Go, cli2Java testservice.Client

// runDubboServer starts dubbo-go server.
// use exitChan to receive exit signal.
func runDubboGoServer(exitChan chan struct{}) {
	config.SetProviderService(&api.UserProviderImpl{})
	if err := config.Load(config.WithPath("./conf/dubbogo.yaml")); err != nil {
		panic(err)
	}
	<-exitChan
}

func waitForPort(port string) {
	for {
		conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", port))
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			conn.Close()
			return
		}
	}
}

func TestMain(m *testing.M) {
	exitChan := make(chan struct{})
	go runDubboGoServer(exitChan)
	cancel, finishChan := runDubboJavaServer()
	// wait for dubbo-go and dubbo-java server initialization
	waitForPort("20000")
	waitForPort("20001")
	var err error
	cli2Go, err = testservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20000"),
		client.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
		)),
	)
	if err != nil {
		panic(err)
	}
	cli2Java, err = testservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20001"),
		client.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
			dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi()),
		)),
	)
	if err != nil {
		panic(err)
	}
	m.Run()
	// close dubbo-go server
	close(exitChan)
	// kill dubbo-java server
	cancel()
	// wait for dubbo-java server terminated
	<-finishChan
	os.Exit(0)
}

func TestEchoRetBool(t *testing.T) {
	resp, err := cli2Go.EchoRetBool(context.Background())
	assertEcho(t, err, false, resp)
}

// dubbo-go does not support
//func TestEchoRetByte(t *testing.T) {
//	resp, err := cli2Go.EchoRetByte(context.Background())
//	assertEcho(t, err, 0, resp)
//}

// dubbo-go does not support
//func TestEchoRetInt16(t *testing.T) {
//	resp, err := cli2Go.EchoRetInt16(context.Background())
//	assertEcho(t, err, 0, resp)
//}

func TestEchoRetInt32(t *testing.T) {
	resp, err := cli2Go.EchoRetInt32(context.Background())
	assertEcho(t, err, int32(0), resp)
}

func TestEchoRetInt64(t *testing.T) {
	resp, err := cli2Go.EchoRetInt64(context.Background())
	assertEcho(t, err, int64(0), resp)
}

func TestEchoRetFloat(t *testing.T) {
	resp, err := cli2Go.EchoRetFloat(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetDouble(t *testing.T) {
	resp, err := cli2Go.EchoRetDouble(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetString(t *testing.T) {
	resp, err := cli2Go.EchoRetString(context.Background())
	assertEcho(t, err, "", resp)
}

func TestEchoBool(t *testing.T) {
	req := true
	resp, err := cli2Go.EchoBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-go does not support
//func TestEchoByte(t *testing.T) {
//	var req int8 = 12
//	resp, err := cli2Go.EchoByte(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// dubbo-go does not support
//func TestEchoInt16(t *testing.T) {
//	var req int16 = 12
//	resp, err := cli2Go.EchoInt16(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

func TestEchoInt32(t *testing.T) {
	var req int32 = 12
	resp, err := cli2Go.EchoInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64(t *testing.T) {
	var req int64 = 12
	resp, err := cli2Go.EchoInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoFloat(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli2Go.EchoFloat(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDouble(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli2Go.EchoDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoString(t *testing.T) {
	req := "12"
	resp, err := cli2Go.EchoString(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// ----------kitex -> dubbo-java----------

func TestEchoRetBool_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetBool(context.Background())
	assertEcho(t, err, false, resp)
}

func TestEchoRetByte_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetByte(context.Background())
	assertEcho(t, err, int8(0), resp)
}

func TestEchoRetInt16_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetInt16(context.Background())
	assertEcho(t, err, int16(0), resp)
}

func TestEchoRetInt32_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetInt32(context.Background())
	assertEcho(t, err, int32(0), resp)
}

func TestEchoRetInt64_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetInt64(context.Background())
	assertEcho(t, err, int64(0), resp)
}

func TestEchoRetFloat_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetFloat(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetDouble_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetDouble(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetString_Java(t *testing.T) {
	resp, err := cli2Java.EchoRetString(context.Background())
	assertEcho(t, err, "", resp)
}

func TestEchoBool_Java(t *testing.T) {
	req := true
	resp, err := cli2Java.EchoBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoByte_Java(t *testing.T) {
	var req int8 = 12
	resp, err := cli2Java.EchoByte(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt16_Java(t *testing.T) {
	var req int16 = 12
	resp, err := cli2Java.EchoInt16(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt32_Java(t *testing.T) {
	var req int32 = 12
	resp, err := cli2Java.EchoInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64_Java(t *testing.T) {
	var req int64 = 12
	resp, err := cli2Java.EchoInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoFloat_Java(t *testing.T) {
	var req float64 = 1.0
	resp, err := cli2Java.EchoFloat(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDouble_Java(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli2Java.EchoDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoString_Java(t *testing.T) {
	req := "12"
	resp, err := cli2Java.EchoString(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func assertEcho(t *testing.T, err error, req, resp interface{}) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(req, resp) {
		t.Fatalf("req is not equal to resp, req: %v, resp: %v", req, resp)
	}
}
