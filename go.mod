module github.com/kitex-contrib/codec-dubbo-tests

go 1.13

require (
	dubbo.apache.org/dubbo-go/v3 v3.1.0
	github.com/apache/dubbo-go-hessian2 v1.12.4
	github.com/cloudwego/kitex v0.8.0
	github.com/dubbogo/gost v1.14.0
	github.com/dubbogo/tools v1.0.9
	github.com/kitex-contrib/codec-dubbo v0.2.2
	github.com/kitex-contrib/codec-dubbo-tests/code/kitex v0.0.0
	github.com/kitex-contrib/codec-dubbo/registries/zookeeper v0.0.0-20231220024013-016c9d2b688f
	github.com/stretchr/testify v1.8.2
)

replace github.com/kitex-contrib/codec-dubbo-tests/code/kitex v0.0.0 => ./code/kitex

replace github.com/kitex-contrib/codec-dubbo v0.2.2 => github.com/DMwangnima/codec-dubbo v0.0.0-20231230090821-a110bb2e0a0b
