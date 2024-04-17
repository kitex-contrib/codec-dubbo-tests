module github.com/kitex-contrib/codec-dubbo-tests

go 1.13

require (
	dubbo.apache.org/dubbo-go/v3 v3.1.0
	github.com/apache/dubbo-go-hessian2 v1.12.4
	github.com/cloudwego/kitex v0.9.1
	github.com/dubbogo/gost v1.14.0
	github.com/dubbogo/tools v1.0.9
	github.com/kitex-contrib/codec-dubbo v0.2.5
	github.com/kitex-contrib/codec-dubbo-tests/code/kitex v0.0.0
	github.com/kitex-contrib/codec-dubbo/registries/zookeeper v0.0.0-20231220024013-016c9d2b688f
	github.com/stretchr/testify v1.8.2
)

replace (
	github.com/kitex-contrib/codec-dubbo => github.com/s5364733/codec-dubbo v1.0.1-0.20240416122451-8e50d9080418
	github.com/kitex-contrib/codec-dubbo-tests/code/kitex v0.0.0 => ./code/kitex
	github.com/kitex-contrib/codec-dubbo-tests/code/kitex/extensions v0.0.0 => ./code/kitex/extensions
)
