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

package kitex2kitex

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo/testservice"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
)

var cli testservice.Client

// initKitexClient inits Kitex client with specified destService and hostPort
func initKitexClient(destService, hostPort string) {
	var err error
	cli, err = testservice.NewClient(destService,
		client.WithHostPorts(hostPort),
		client.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
		)),
	)
	if err != nil {
		panic(fmt.Sprintf("Kitex client initialized failed, err :%s", err))
	}
}

// runKitexServer starts Kitex server for testing based on specified address.
// use startCh to tell outer layer that Kitex server has already started.
// use exitCh to receive exiting signal.
func runKitexServer(startCh chan struct{}, exitCh chan error, addr string) {
	netAddr, _ := net.ResolveTCPAddr("tcp", addr)
	svr := testservice.NewServer(
		new(handler.TestServiceImpl),
		server.WithServiceAddr(netAddr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
		)),
		server.WithExitSignal(func() <-chan error {
			return exitCh
		}),
	)
	server.RegisterStartHook(func() {
		close(startCh)
	})

	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	startCh := make(chan struct{})
	exitCh := make(chan error)
	go runKitexServer(startCh, exitCh, ":20000")
	<-startCh
	initKitexClient("test", "127.0.0.1:20000")
	m.Run()
	exitCh <- nil
	os.Exit(0)
}

func TestEchoRetBool(t *testing.T) {
	resp, err := cli.EchoRetBool(context.Background())
	assertEcho(t, err, false, resp)
}

func TestEchoRetByte(t *testing.T) {
	resp, err := cli.EchoRetByte(context.Background())
	assertEcho(t, err, int8(0), resp)
}

func TestEchoRetInt16(t *testing.T) {
	resp, err := cli.EchoRetInt16(context.Background())
	assertEcho(t, err, int16(0), resp)
}

func TestEchoRetInt32(t *testing.T) {
	resp, err := cli.EchoRetInt32(context.Background())
	assertEcho(t, err, int32(0), resp)
}

func TestEchoRetInt64(t *testing.T) {
	resp, err := cli.EchoRetInt64(context.Background())
	assertEcho(t, err, int64(0), resp)
}

func TestEchoRetFloat(t *testing.T) {
	resp, err := cli.EchoRetFloat(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetDouble(t *testing.T) {
	resp, err := cli.EchoRetDouble(context.Background())
	assertEcho(t, err, float64(0), resp)
}

func TestEchoRetString(t *testing.T) {
	resp, err := cli.EchoRetString(context.Background())
	assertEcho(t, err, "", resp)
}

func TestEchoBool(t *testing.T) {
	req := true
	resp, err := cli.EchoBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoByte(t *testing.T) {
	var req int8 = 12
	resp, err := cli.EchoByte(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt16(t *testing.T) {
	var req int16 = 12
	resp, err := cli.EchoInt16(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt32(t *testing.T) {
	var req int32 = 12
	resp, err := cli.EchoInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64(t *testing.T) {
	var req int64 = 12
	resp, err := cli.EchoInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoFloat(t *testing.T) {
	var req float64 = 1.0
	resp, err := cli.EchoFloat(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDouble(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli.EchoDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoString(t *testing.T) {
	req := "12"
	resp, err := cli.EchoString(context.Background(), req)
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
