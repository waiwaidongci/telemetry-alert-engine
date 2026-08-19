# Bug 是什么

规则默认 map 未初始化，禁用校验器以 typed-nil 接口返回，服务写入和入口判空分别失效。

# 如何触发

执行 `go test ./internal/ruledefaults -run '^TestRuleZeroValueInitializesPolicyAndRejectsTypedNil$' -count=1`。

# 错误信息

测试先报告 default config left nil rules 与 disabled factory returned typed nil，随后 panic：`assignment to entry in nil map`。
