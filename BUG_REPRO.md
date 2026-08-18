# Bug
窗口存储、缓存和聚合步骤共享同一张可变 map，后续写入会反向改变已经返回的快照。

# 触发
在项目根目录运行 `go test -race ./internal/windowcache -run TestSnapshotsDoNotShareMutableMaps`。

# 错误信息
`snapshot changed to 20`
