# Bug 是什么

窗口存储、过滤、缓存和聚合服务共享调用方切片底层数组，空缓存读取还会越界 panic。

# 如何触发

执行 `go test ./internal/querysnapshot -run '^TestWindowCopiesRemainIndependent$' -count=1`，再执行 `go test ./internal/querysnapshot -run '^TestVacantWindowReadIsSafe$' -count=1`。

# 错误信息

历史快照断言全部失败；空读报 `panic: runtime error: slice bounds out of range [:1] with capacity 0`。
