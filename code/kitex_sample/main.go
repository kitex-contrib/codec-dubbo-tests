/*
 * Copyright 2024 CloudWeGo Authors
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

package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/detection"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex_sample/handler"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex_sample/kitex_gen/hello/greetservice"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":9102")
	if err != nil {
		klog.Fatal(err)
	}

	svr := greetservice.NewServer(new(handler.GreetServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "hello",
		}),
		server.WithServiceAddr(addr),
		server.WithTransHandlerFactory(detection.NewSvrTransHandlerFactory(netpoll.NewSvrTransHandlerFactory(),
			dubbo.NewSvrTransHandlerFactory(
				// set the Java interface name corresponding to the Kitex service
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider")),
			nphttp2.NewSvrTransHandlerFactory(),
		)),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
