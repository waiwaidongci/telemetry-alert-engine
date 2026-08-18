# Bug
查询缺失设备时错误身份丢失，服务返回 500 并把请求当成可重试错误。

# 触发
在项目根目录运行 `go test ./internal/devicequery -run TestMissingDeviceKeepsNotFoundContract`。

# 错误信息
`not-found identity was lost`
