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
	"github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEchoException_Java(t *testing.T) {
	_, err := cli2Java.EchoException(context.Background(), true)
	exception, ok := hessian2_exception.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, "java.lang.Exception", exception.JavaClassName())
	assert.Equal(t, "EchoException", exception.Error())
}

func TestEchoCustomizedException_Java(t *testing.T) {
	_, err := cli2Java.EchoCustomizedException(context.Background(), true)
	exception, ok := hessian2_exception.FromError(err)
	assert.True(t, ok)
	customizedException, ok := exception.(*echo.EchoCustomizedException)
	assert.True(t, ok)
	assert.Equal(t, "org.apache.dubbo.tests.api.EchoCustomizedException", customizedException.JavaClassName())
	assert.Equal(t, "EchoCustomizedException", customizedException.CustomizedMessage)
}
