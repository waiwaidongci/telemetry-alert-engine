# Bug

多源对账的错误通道、producer、WaitGroup 和 consumer 生命周期错位，非法读数会截断合法结果并造成等待超时。

# 触发方式

运行 `go test ./rh18u -run '^TestHotelPipeline018$' -count=1 "-race"`，两个 producer 经同一屏障并发启动，其中一个源包含负数。

# 错误信息

```text
error bus blocked before configured capacity
producer lost valid reading
consumer did not drain both channels
incomplete readings
```
