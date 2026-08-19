# Bug

路由策略零配置包含 nil map，关闭校验时返回 typed-nil 接口，空路由入口也未被拒绝。

# 触发方式

运行 `go test ./nd14q -run '^TestDeltaRouting014$' -count=1`，覆盖默认写入、关闭校验和空输入三条路径。

# 错误信息

```text
panic: assignment to entry in nil map [recovered, repanicked]
github.com/example/telemetry-alert/internal/routingpolicy.(*Service).Apply(...)
```
