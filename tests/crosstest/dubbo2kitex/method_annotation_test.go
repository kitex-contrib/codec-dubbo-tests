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

package dubbo2kitex

import (
	"context"
	"fmt"
	"testing"
)

func TestEchoMethodA(t *testing.T) {
	var req bool = true
	resp, err := cli.EchoMethodA(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("A:%v", req), resp)
}

func TestEchoMethodB(t *testing.T) {
	var req int32 = 1
	resp, err := cli.EchoMethodB(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("B:%v", req), resp)
}

func TestEchoMethodC(t *testing.T) {
	var req int32 = 1
	resp, err := cli.EchoMethodC(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("C:%v", req), resp)
}

func TestEchoMethodD(t *testing.T) {
	var req1 bool = true
	var req2 int32 = 1
	resp, err := cli.EchoMethodD(context.Background(), req1, req2)
	assertEcho(t, err, fmt.Sprintf("D:%v,%v", req1, req2), resp)
}
