# Common包日志代理

这个包提供了一个GoFrame g.Log()的代理实现，使得common包可以在不直接依赖GoFrame的情况下使用日志功能。

## 架构设计

1. **ILogger接口**: 定义了日志操作的接口，兼容GoFrame的g.Log()方法
2. **日志代理**: 提供Log()函数，模拟g.Log()的使用方式  
3. **GoFrame适配器**: 将GoFrame的日志实现适配到我们的接口
4. **默认实现**: 使用标准库log作为fallback实现

## 使用方式

### 在应用启动时注入GoFrame日志实现（cmd.go中）

```go
import "common/c_log"

// 注入GoFrame日志实现到common包
c_log.SetLogger(c_log.NewGoFrameLoggerAdapter(g.Log()))
```

### 在common包中使用日志

```go
import "common/c_log"

func SomeFunction(ctx context.Context) {
    // 使用方式与g.Log()完全一样
    c_log.Log().Infof(ctx, "这是一条信息日志: %s", "示例")
    c_log.Log().Debugf(ctx, "这是一条调试日志: %d", 123)
    c_log.Log().Warningf(ctx, "这是一条警告日志")
    c_log.Log().Errorf(ctx, "这是一条错误日志: %v", "错误信息")
}
```

## 优势

1. **解耦**: common包不需要直接依赖GoFrame
2. **兼容**: 使用方式与g.Log()完全一致
3. **灵活**: 可以轻松切换不同的日志实现
4. **默认支持**: 即使没有注入实现，也有标准库的fallback

## 支持的日志级别

- Debug/Debugf
- Info/Infof  
- Notice/Noticef
- Warning/Warningf
- Error/Errorf
- Critical/Criticalf
- Panic/Panicf
- Fatal/Fatalf
