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

package kitex2dubbo

import (
	"context"
	"testing"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
)

func TestEchoJavaEnum_Java(t *testing.T) {
	t.Run("raw enumeration test", func(t *testing.T) {
		resp1 := "ONE"
		req, err1 := cli2Java.EchoJavaEnum(context.Background(), echo.KitexEnum_ONE)
		assertEcho(t, err1, req.String(), resp1)
		resp2 := "TWO"
		req, err2 := cli2Java.EchoJavaEnum(context.Background(), echo.KitexEnum_TWO)
		assertEcho(t, err2, req.String(), resp2)
		resp3 := "THREE"
		req, err3 := cli2Java.EchoJavaEnum(context.Background(), echo.KitexEnum_THREE)
		assertEcho(t, err3, req.String(), resp3)
	})

	t.Run("with args enumeration test", func(t *testing.T) {
		resp1 := &echo.EchoEnumResponse{
			Echo:      "hello1",
			KitexEnum: echo.KitexEnum_ONE,
		}
		request1 := &echo.EchoEnumRequest{
			Echo:      "hello1",
			KitexEnum: echo.KitexEnum_ONE,
		}
		req1, err := cli2Java.EchoJavaEnumWithArg(context.Background(), request1)
		assertEcho(t, err, req1, resp1)

		resp2 := &echo.EchoEnumResponse{
			Echo:      "hello2",
			KitexEnum: echo.KitexEnum_TWO,
		}
		request2 := &echo.EchoEnumRequest{
			Echo:      "hello2",
			KitexEnum: echo.KitexEnum_TWO,
		}
		req2, err := cli2Java.EchoJavaEnumWithArg(context.Background(), request2)
		assertEcho(t, err, req2, resp2)

		resp3 := &echo.EchoEnumResponse{
			Echo:      "hello3",
			KitexEnum: echo.KitexEnum_THREE,
		}
		request3 := &echo.EchoEnumRequest{
			Echo:      "hello3",
			KitexEnum: echo.KitexEnum_THREE,
		}
		req3, err := cli2Java.EchoJavaEnumWithArg(context.Background(), request3)
		assertEcho(t, err, req3, resp3)
	})
}
