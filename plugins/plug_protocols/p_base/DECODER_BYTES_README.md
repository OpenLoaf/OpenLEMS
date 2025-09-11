# DecoderBytes 函数实现说明

## 概述

`DecoderBytes` 函数是一个通用的字节解析函数，支持多种工业协议的数据解析，包括 Modbus、CANbus、IEC 61850、S7 等协议。

## 功能特性

### 1. 支持的数据格式

- **整数类型**: 16位/32位有符号和无符号整数
- **浮点数类型**: 32位/64位IEEE 754浮点数
- **BCD码**: 16位和32位二进制编码十进制数
- **字符串**: ASCII和UTF-16字符串
- **位图**: 位级数据解析
- **位范围**: 精确提取字节中特定范围的位
- **自定义格式**: 支持扩展自定义解析

### 2. 字节序支持

- **字节序**: 大端序(ByteEndianBig) / 小端序(ByteEndianLittle)
- **字序**: 高字在前(WordOrderHighLow) / 低字在前(WordOrderLowHigh)

### 3. 数据处理

- **系数处理**: 支持浮点数系数乘法
- **偏移量处理**: 支持数值偏移
- **类型转换**: 自动转换为目标系统类型

## 函数签名

```go
func DecoderBytes(
    bytes []byte,           // 原始字节数据
    byteIndex uint16,       // 字节起始索引
    byteLength uint16,      // 字节长度（0表示纯位模式）
    bitIndex uint16,        // 位起始索引
    bitLength uint16,       // 位长度（0表示纯字节模式）
    byteEndian EByteEndian, // 字节序
    wordOrder EWordOrder,   // 字序
    dataFormat EDataFormat, // 数据格式
    returnFormat EValueType, // 返回格式类型
    offset int,             // 偏移量
    factor float32,         // 系数
    min int64,              // 最小值验证（数值类型）或最小长度验证（字符串类型，0表示不验证）
    max int64,              // 最大值验证（数值类型）或最大长度验证（字符串类型，0表示不验证）
) (any, error)             // 返回解析结果和错误
```

## 三种读取模式

### 1. 纯字节模式 (bitLength=0)

- **用途**: 读取指定字节范围的数据
- **参数**: `byteIndex` 和 `byteLength` 有效，`bitIndex=0, bitLength=0`
- **示例**: `DecoderBytes(data, 0, 2, 0, 0, ...)` - 读取第0字节开始的2字节数据

### 2. 纯位模式 (byteLength=0)

- **用途**: 从数据开头读取指定位范围
- **参数**: `bitIndex` 和 `bitLength` 有效，`byteIndex=0, byteLength=0`
- **示例**: `DecoderBytes(data, 0, 0, 5, 3, ...)` - 从第5位开始读取3位数据

### 3. 混合模式 (两者都不为0)

- **用途**: 先读取指定字节范围，再从中提取指定位
- **参数**: 所有参数都有效
- **示例**: `DecoderBytes(data, 2, 2, 3, 5, ...)` - 从第2字节开始读取2字节，然后提取第3-7位

## 使用示例

### 模式1：Modbus 16位整数解析（纯字节模式）

```go
data := []byte{0x12, 0x34}
result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, EInt16, 0, 1.0, 0, 10000)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: int16(0x1234)
```

### 模式2：CANbus 位数据解析（纯位模式）

```go
data := []byte{0xB6} // 0b10110110
result, err := DecoderBytes(data, 0, 0, 5, 3, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EInt32, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: int32(3) - 第5-7位的值为011(二进制) = 3(十进制)
```

### 模式3：混合模式 - 从第2字节开始读取2字节，然后提取第3-7位

```go
data := []byte{0x12, 0x34, 0xB6, 0x78}
result, err := DecoderBytes(data, 2, 2, 3, 5, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EInt32, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: int32(22) - 从0xB678中提取第3-7位的值
```

### CANbus BCD码解析（纯字节模式，无范围验证）

```go
data := []byte{0x12, 0x34}
result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatBCD, EInt32, 0, 0.1, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: 123.4 (1234 * 0.1)
```

### IEEE 754浮点数解析（纯字节模式，带范围验证）

```go
data := []byte{0x40, 0x49, 0x0F, 0xDB} // 3.14159的IEEE 754表示
result, err := DecoderBytes(data, 0, 4, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatFloat32, EFloat64, 0, 1.0, -100, 100)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: 3.14159
```

### 小端字节序解析（纯字节模式，带范围验证）

```go
data := []byte{0x34, 0x12}
result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianLittle, WordOrderHighLow, DataFormatUInt16, EInt16, 0, 1.0, 0, 10000)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: int16(0x1234)
```

### ASCII字符串解析（纯字节模式，带长度验证）

```go
data := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'}
result, err := DecoderBytes(data, 0, 11, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatStringASCII, EString, 0, 1.0, 3, 20)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: "Hello World"
```

### 位范围解析示例（使用纯位模式）

#### 获取第3个bit的值（从低位开始，0-based）

```go
data := []byte{0xB6} // 0b10110110
result, err := DecoderBytes(data, 0, 0, 3, 1, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EUint8, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: uint8(0) - 第3个bit的值为0
```

#### 获取从第4个bit起后2位bit的值

```go
data := []byte{0xB6} // 0b10110110
result, err := DecoderBytes(data, 0, 0, 4, 2, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EUint8, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: uint8(3) - 第4-5位bit的值为11(二进制) = 3(十进制)
```

#### 获取最高位（MSB）

```go
data := []byte{0xB6} // 0b10110110
result, err := DecoderBytes(data, 0, 0, 7, 1, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EUint8, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: uint8(1) - 最高位为1
```

#### 获取最低位（LSB）

```go
data := []byte{0xB6} // 0b10110110
result, err := DecoderBytes(data, 0, 0, 0, 1, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EUint8, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误
    log.Printf("解析失败: %v", err)
    return
}
// 结果: uint8(0) - 最低位为0
```

## 性能优化

### 1. 内存优化

- 预分配结果切片，避免多次内存分配
- 缓存浮点数转换，减少重复计算
- 零值检查，避免不必要的计算

### 2. 算法优化

- 高效的字节序重排算法
- 优化的BCD码解析
- 快速类型转换

### 3. 错误处理

- 边界检查，防止数组越界
- 参数验证，确保输入有效性
- 完整的错误传播，所有错误都会向上传递
- 详细的错误信息，便于调试和问题定位

### 4. 数据验证

- 支持数值范围验证（min/max）
- 支持字符串长度验证（min/max作为长度限制）
- 自动跳过布尔类型的验证
- 灵活的验证控制（min=0, max=0 表示不验证）
- 验证失败时返回详细的错误信息

## 协议支持

### Modbus TCP/RTU

- 支持所有标准数据类型
- 支持大端和小端字节序
- 支持字序交换

### CANbus

- 支持位级数据读取
- 支持BCD码解析
- 支持多字节序

### IEC 61850

- 支持IEEE 754浮点数
- 支持UTF-16字符串
- 支持复杂数据结构

### S7 (西门子PLC)

- 支持西门子特有的数据格式
- 支持位操作
- 支持字符串处理

## 测试覆盖

函数包含完整的单元测试，覆盖：

- 所有数据格式的解析
- 字节序和字序的组合
- 边界情况处理
- 错误情况处理
- 性能测试

## 扩展性

函数设计具有良好的扩展性：

- 可以轻松添加新的数据格式
- 支持自定义解析函数
- 可以扩展新的字节序类型
- 支持新的协议类型

## 注意事项

1. **系数处理**: 如果系数为0，会自动设置为1.0
2. **边界检查**: 函数会检查数据长度，防止越界访问
3. **类型安全**: 所有类型转换都是安全的，不会导致panic
4. **错误处理**: 必须检查返回的错误，不能忽略
5. **数据验证**: min和max参数用于数值范围验证或字符串长度验证，0表示不验证
6. **性能考虑**: 对于高频调用场景，建议缓存解析结果

## 错误处理

函数现在返回 `(any, error)`，调用者必须处理可能的错误：

```go
result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, EInt16, 0, 1.0, 0, 0)
if err != nil {
    // 处理错误情况
    switch {
    case strings.Contains(err.Error(), "empty input data"):
        // 处理空数据
    case strings.Contains(err.Error(), "insufficient data"):
        // 处理数据不足
    case strings.Contains(err.Error(), "unsupported data format"):
        // 处理不支持的数据格式
    case strings.Contains(err.Error(), "both byteLength and bitLength cannot be zero"):
        // 处理参数错误
    default:
        // 处理其他错误
    }
    return
}
// 使用解析结果
```

## 数据验证

函数支持对解析后的数据进行范围验证：

```go
// 带范围验证的解析
result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, EInt16, 0, 1.0, 0, 1000)
if err != nil {
    // 处理错误，包括验证失败
    if strings.Contains(err.Error(), "below minimum") {
        // 处理值低于最小值的情况
    } else if strings.Contains(err.Error(), "above maximum") {
        // 处理值高于最大值的情况
    }
    return
}
```

### 验证规则

1. **数值类型验证**: 所有数值类型（整数、浮点数）都会进行范围验证
2. **字符串类型**: 进行长度验证（min/max作为长度限制）
3. **布尔类型**: 自动跳过范围验证
4. **位范围类型**: 自动跳过范围验证（位值本身已经有限制）
5. **验证控制**:
   - `min=0, max=0`: 不进行验证
   - `min>0 或 max>0`: 进行验证
6. **验证时机**: 在数据解析、系数处理、类型转换完成后进行验证

### 位范围解析规则

1. **位位置**: 位位置从0开始计算，0是最低位（LSB），7是最高位（MSB）
2. **位长度**: 可以提取1-64个连续的位（支持跨字节操作）
3. **边界检查**: 确保位范围不超出数据长度
4. **返回值**: 根据位长度返回适当类型（uint8, uint16, uint32, uint64）
5. **字节序**: 位范围解析不受字节序和字序影响，按字节顺序提取
6. **跨字节支持**: 支持跨多个字节的位范围提取

### 验证错误类型

- `"value %d is below minimum %d"`: 值低于最小值
- `"value %d is above maximum %d"`: 值高于最大值
- `"string length %d is below minimum %d"`: 字符串长度低于最小值
- `"string length %d is above maximum %d"`: 字符串长度高于最大值
- `"value validation failed: %w"`: 验证过程中的其他错误

## EncoderBytes 编码器

除了解码器，还提供了对应的编码器 `EncoderBytes`，用于生成各种协议的发送指令。

### 编码器函数签名

```go
func EncoderBytes(
    value any,              // 要编码的值
    dataFormat EDataFormat, // 数据格式
    byteEndian EByteEndian, // 字节序
    wordOrder EWordOrder,   // 字序
    offset int,             // 偏移量（编码前应用）
    factor float32,         // 系数（编码前应用）
    min int64,              // 最小值验证
    max int64,              // 最大值验证
) ([]byte, error)          // 返回编码后的字节数组和错误
```

### 编码器使用示例

```go
// 编码16位整数用于Modbus发送
data, err := EncoderBytes(1234, DataFormatUInt16, ByteEndianBig, WordOrderHighLow, 0, 1.0, 0, 1000)

// 编码BCD码用于CANbus发送
data, err := EncoderBytes(1234, DataFormatBCD, ByteEndianLittle, WordOrderHighLow, 0, 0.1, 0, 0)

// 编码浮点数用于IEC 61850发送
data, err := EncoderBytes(3.14, DataFormatFloat32, ByteEndianBig, WordOrderHighLow, 0, 1.0, -100, 100)

// 编码字符串用于发送
data, err := EncoderBytes("Hello", DataFormatStringASCII, ByteEndianBig, WordOrderHighLow, 0, 1.0, 3, 20)
```

### 编码器特点

1. **协议支持**: 支持Modbus、CANbus、IEC 61850、S7等协议
2. **数据格式**: 支持所有解码器支持的数据格式
3. **字节序处理**: 支持大端/小端字节序和字序处理
4. **系数偏移**: 编码前反向应用系数和偏移量
5. **数据验证**: 编码前进行数据范围验证
6. **类型安全**: 安全的类型转换和错误处理

## 维护和更新

- 函数设计遵循Go语言最佳实践
- 代码注释完整，便于维护
- 测试覆盖率高，确保稳定性
- 支持向后兼容，不会破坏现有代码
- 编解码器配对设计，功能完整
