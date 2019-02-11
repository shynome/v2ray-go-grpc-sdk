### 简要描述

这是 `v2ray.com/core` 的一个子集, 只包含了调用 `grpc` 中必要的代码, 
但目前不是所有 `grpc` 接口代码都会放进来, 只会优先放置一些我会用的接口,
如果没有你所需要的可以 `fork` 一份自己修改, 如果能贡献到本仓库就更好了

### 使用方法:
```shell
go mod edit -replace=v2ray.com@v4.14.0=github.com/shynome/v2ray-go-grpc-sdk@master
```
