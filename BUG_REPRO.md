# Bug

扫描请求的 deadline 没有传播到网关，仓储还复用首次请求上下文，取消后的重试循环不能及时退出。

# 触发方式

运行 `go test ./oe15r -run '^TestEchoSweep015$' -count=1`，先发短超时扫描，再发正常扫描并检查取消后的重试计数。

# 错误信息

```text
expected deadline error, got <nil>
expected cancellation, got <nil>
```
