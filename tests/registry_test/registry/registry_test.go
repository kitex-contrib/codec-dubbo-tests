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

package registry

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/kitex-contrib/codec-dubbo-tests/code/dubbo-go/api"
	"github.com/kitex-contrib/codec-dubbo-tests/util"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/cloudwego/kitex/client"
	registry2 "github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo/testservice"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries"
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/registry"
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/resolver"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	zookeeperAddress1 := "127.0.0.1:2181"
	kitexAddress := ":20000"
	javaInterfaceName := "org.apache.dubbo.tests.api.UserProvider"
	tests := []struct {
		regOpts      []registry.Option
		codecOpts    []dubbo.Option
		tags         map[string]string
		cliResOpts   []resolver.Option
		cliCodecOpts []dubbo.Option
		cliOpts      []client.Option
		cliCfgPath   string
	}{
		{
			regOpts: []registry.Option{
				registry.WithServers(zookeeperAddress1),
				registry.WithRegistryGroup("myGroup"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			tags: map[string]string{
				registries.DubboServiceInterfaceKey:   javaInterfaceName,
				registries.DubboServiceApplicationKey: "dubbo",
			},
			cliResOpts: []resolver.Option{
				resolver.WithServers(zookeeperAddress1),
				resolver.WithRegistryGroup("myGroup"),
			},
			cliCodecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			cliOpts: []client.Option{
				client.WithTag(registries.DubboServiceInterfaceKey, javaInterfaceName),
			},
			cliCfgPath: "./conf/dubbogo.yaml",
		},
	}

	for _, test := range tests {
		startChan := make(chan struct{})
		exitChan := make(chan error)
		exitFinish := make(chan struct{})
		go startKitexServerWithRegistry(t, startChan, exitChan, exitFinish, kitexAddress, test.regOpts, test.codecOpts, test.tags)
		<-startChan
		// wait for registering
		time.Sleep(3 * time.Second)
		testKitexClient(t, test.cliResOpts, test.cliCodecOpts, test.cliOpts)
		testDubboJavaClient(t)
		testDubboGoClient(t, test.cliCfgPath)
		exitChan <- nil
		<-exitFinish
	}
}

func startKitexServerWithRegistry(t *testing.T, startCh chan struct{}, exitCh chan error, exitFinish chan struct{},
	addr string, regOpts []registry.Option, codecOpts []dubbo.Option, tags map[string]string,
) {
	netAddr, _ := net.ResolveTCPAddr("tcp", addr)
	reg, err := registry.NewZookeeperRegistry(regOpts...)
	assert.Nil(t, err)
	svr := testservice.NewServer(
		new(handler.TestServiceImpl),
		server.WithServiceAddr(netAddr),
		server.WithRegistry(reg),
		server.WithRegistryInfo(&registry2.Info{
			Tags: tags,
		}),
		server.WithCodec(dubbo.NewDubboCodec(codecOpts...)),
		server.WithExitSignal(func() <-chan error {
			return exitCh
		}),
	)
	server.RegisterStartHook(func() {
		close(startCh)
	})
	assert.Nil(t, svr.Run())
	exitFinish <- struct{}{}
}

func testKitexClient(t *testing.T, resOpts []resolver.Option, codecOpts []dubbo.Option, cliOpts []client.Option) {
	res, err := resolver.NewZookeeperResolver(resOpts...)
	assert.Nil(t, err)
	opts := []client.Option{
		client.WithResolver(res),
		client.WithCodec(
			dubbo.NewDubboCodec(codecOpts...),
		),
	}
	opts = append(opts, cliOpts...)
	cli, err := testservice.NewClient("testtest", opts...)
	assert.Nil(t, err)
	resp, err := cli.EchoBool(context.Background(), true)
	assert.Nil(t, err)
	assert.True(t, resp)
}

func testDubboGoClient(t *testing.T, configPath string) {
	cli := api.UserProviderClient
	err := config.Load(config.WithPath(configPath))
	assert.Nil(t, err)
	resp, err := cli.EchoBool(context.Background(), true)
	assert.Nil(t, err)
	assert.True(t, resp)
}

func testDubboJavaClient(t *testing.T) {
	util.RunAndTestDubboJavaClient(t, "../../../code/dubbo-java", "org.apache.dubbo.tests.client.Application",
		[]string{"withRegistry"},
		[]string{
			"EchoBool",
		},
	)
}
