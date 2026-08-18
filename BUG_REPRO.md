# Bug
缺省告警配置的 map 为 nil，写入时 panic；装有 nil 指针的校验器接口又被当成有效对象。

# 触发
在项目根目录运行 `go test ./internal/ruleconfig -run TestDefaultRuleConfigHandlesZeroValues`。

# 错误信息
`apply panicked: assignment to entry in nil map`、`disabled validator is a typed nil`
