# Bug

批量导出遇到拒绝项时，资源租约未及时释放，事务收尾覆盖业务失败，服务预留状态也没有复位。

# 触发方式

运行 `go test ./lb12y -run '^TestBravoLease012$' -count=1`，批次前两项正常、第三项被拒绝。

# 错误信息

```text
expected rejected payload, got export resource limit reached: 2
```
