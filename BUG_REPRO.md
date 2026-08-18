# Bug
通知从 `retrying` 重试成功后无法进入 `sent`，正在重试的记录也不会出现在活动列表。

# 触发
在项目根目录运行 `go test ./internal/deliveryflow -run TestRetrySuccessReachesSentAndActiveQueryIncludesRetrying`。

# 错误信息
`retrying delivery was treated as terminal`
