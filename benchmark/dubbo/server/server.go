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

package main

import (
	"flag"
	"strconv"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func main() {
	var srvPort int
	flag.IntVar(&srvPort, "p", 20001, "")
	flag.Parse()

	config.SetProviderService(&UserProvider{})
	protCfg := config.NewProtocolConfigBuilder().
		SetName("dubbo").
		SetPort(strconv.Itoa(srvPort)).
		Build()
	svcCfg := config.NewServiceConfigBuilder().
		SetProtocolIDs("dubbo").
		SetInterface("org.apache.dubbo.UserProvider").
		Build()
	proCfg := config.NewProviderConfigBuilder().
		AddService("UserProvider", svcCfg).
		Build()
	logCfg := config.NewLoggerConfigBuilder().
		SetLevel("error").
		Build()
	rootCfg := config.NewRootConfigBuilder().
		AddProtocol("dubbo", protCfg).
		SetProvider(proCfg).
		SetLogger(logCfg).
		Build()
	if err := rootCfg.Init(); err != nil {
		panic(err)
	}
	select {}
}
