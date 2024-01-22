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
	"github.com/stretchr/testify/assert"
)

func TestEchoGeneric_Java(t *testing.T) {
	list := []interface{}{
		&echo.EchoGenericEmbedded{
			InternalStringReq:  "EchoGeneric1",
			InternalIntegerReq: 1,
		},
		&echo.EchoGenericEmbedded{
			InternalStringReq:  "EchoGeneric2",
			InternalIntegerReq: 2,
		},
	}
	req := &echo.EchoGenericRequest{
		ReqField: 1,
		List:     list,
	}
	resp, err := cli2Java.EchoGeneric(context.Background(), req)
	assert.Equal(t, req.ReqField, resp.RespField)
	assert.Nil(t, err)
	for i, itemRaw := range resp.List {
		item, ok := itemRaw.(*echo.EchoGenericEmbedded)
		assert.True(t, ok)
		assert.Equal(t, list[i], item)
	}
}
