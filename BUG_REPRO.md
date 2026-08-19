# Bug 是什么

告警扇出只生产首条结果，协调器在 worker 完成前关闭 done，consumer 吞掉取消，错误通道没有容量。

# 如何触发

执行 `go test -race ./internal/fanoutpipeline -run '^TestNotificationFanoutWaitsForAllTargets$' -count=1`。

# 错误信息

测试报告 len=1、coordinator completed before workers started、consumer ignored cancellation、error channel cannot accept one error per worker。
