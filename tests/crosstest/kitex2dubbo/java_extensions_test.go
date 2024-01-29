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
	"time"

	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/java"
)

func TestEchoJavaDate_Java(t *testing.T) {
	req := time.Unix(1000, 0)
	resp, err := cli2Java.EchoJavaDate(context.Background(), &req)
	assertEcho(t, err, &req, resp)
}

func TestEchoJavaDateList_Java(t *testing.T) {
	date1 := time.Unix(1001, 0)
	date2 := time.Unix(1002, 0)
	req := []*java.Date{&date1, &date2}
	resp, err := cli2Java.EchoJavaDateList(context.Background(), req)
	assertEcho(t, err, req, resp)
}
