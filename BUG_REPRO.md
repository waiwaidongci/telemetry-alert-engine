# Bug

窗口过滤、存储和缓存共享切片底层数组及标签 map，调用方修改返回值或追加采样会污染源数据和历史快照。

# 触发方式

运行 `go test ./pf16s -run '^TestFoxtrotWindow016$' -count=1`，先过滤并发布窗口，再修改返回标签并追加新点。

# 错误信息

```text
stored window was aliased: []metricwindow.Point{metricwindow.Point{Sequence:4, Value:20, Tags:map[string]string{"state":"hot"}}}
```
