# Bug
批量导出会把资源释放推迟到函数结束，业务错误可能被提交结果覆盖，错误分支还会遗留句柄。

# 触发
在项目根目录运行 `go test ./internal/exportbatch -run TestExportResourcesAndErrors`。

# 错误信息
`peak=8 open=0`、`business error was lost`、`open=-1`
