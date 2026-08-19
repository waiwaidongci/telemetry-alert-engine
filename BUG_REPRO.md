# Bug 是什么

遥测采集入口丢失请求 context，网关与重试 worker 都不会响应取消，下一次调用也无法保留自己的请求值。

# 如何触发

执行 `go test ./internal/ingestcontext -run '^TestRequestCancellationStopsTelemetryRetries$' -count=1`。

# 错误信息

测试同时报告 request lost its cancellation、gateway ignored cancellation、worker retried after cancellation、service replaced request context。
