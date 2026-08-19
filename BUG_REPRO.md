# Bug

派发超时和父请求取消没有传到子任务，取消原因丢失，服务还会复用首次请求的上下文。

# 触发方式

运行 `go test ./tj20w -run '^TestJulietDispatch020$' -count=1`，覆盖慢目标超时、随后快速派发和父请求预取消。

# 错误信息

```text
timeout cause was lost: <nil>
first dispatch did not time out: <nil>
parent cancellation was lost: <nil>
```
