# Bug

并发刷新告警规则时，存储快照与缓存暴露内部可变引用，产生 data race，停用规则会混入结果，旧快照也会被后续写入改变。

# 触发方式

运行 `go test ./ka11x -run '^TestKappaSnapshot011$' -count=1 '-race'`，测试用同一启动屏障并发更新和刷新同一规则。

# 错误信息

```text
WARNING: DATA RACE
Write at 0x00c000096690 by goroutine 8
Previous read at 0x00c000096690 by goroutine 9
testing.go:1712: race detected during execution of test
```
