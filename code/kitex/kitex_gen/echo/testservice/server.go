// Code generated by Kitex v0.8.0. DO NOT EDIT.
package testservice

import (
	server "github.com/cloudwego/kitex/server"
	echo "github.com/kitex-contrib/codec-dubbo-tests/code/kitex/kitex_gen/echo"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo.TestService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}