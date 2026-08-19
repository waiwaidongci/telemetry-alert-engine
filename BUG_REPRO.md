# Bug

告警升级重试成功后仍停在 retrying，历史被覆盖，进行中和已完成查询对中间态的分类相反。

# 触发方式

运行 `go test ./qg17t -run '^TestGolfState017$' -count=1`，驱动 failed 到 retrying 再到成功，并检查两个查询集合。

# 错误信息

```text
unexpected final state "retrying"
retrying incident disappeared
```
