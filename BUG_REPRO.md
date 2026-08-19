# Bug 是什么

通知记录缺失错误在仓储、分类、HTTP 映射和重试策略之间丢失身份，最终成为 500 并被重试。

# 如何触发

执行 `go test ./internal/notificationerrors -run '^TestAbsentDeliveryPreservesSentinelAndStopsQueue$' -count=1`。

# 错误信息

测试报告 sentinel chain broken、classifier returned system、handler returned 500、missing attempt was scheduled for retry。
