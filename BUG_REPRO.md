# Bug

批量查询缺失投递路由时，错误链身份和第二条缺失上下文丢失，永久错误被归为服务故障并进入重试。

# 触发方式

运行 `go test ./mc13z -run '^TestCharlieRoute013$' -count=1`，同时查询两个缺失路由和一个已存在路由。

# 错误信息

```text
missing sentinel not preserved: route missing-a state: route lookup failed: delivery route missing
```
