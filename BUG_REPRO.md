# Bug

策略额度 typed error 在两层包装和聚合中断链，HTTP 状态退化为 500，任务被误判为可重试。

# 触发方式

运行 `go test ./si19v -run '^TestIndiaPolicy019$' -count=1`，让两个策略同时超过活动告警上限。

# 错误信息

```text
limit error type was lost: policy evaluation failed: evaluate alert capacity: alert capacity exceeded: 5
```
