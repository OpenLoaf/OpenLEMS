package c_base

// EDataFormat 数据格式/编码
type EDataFormat int

const (
	DataFormatUInt16      EDataFormat = iota // 16位无符号整数
	DataFormatInt16                          // 16位有符号整数 (二进制补码)
	DataFormatUInt32                         // 32位无符号整数
	DataFormatInt32                          // 32位有符号整数
	DataFormatFloat32                        // 32位IEEE 754浮点数
	DataFormatFloat64                        // 64位IEEE 754浮点数
	DataFormatBCD                            // 二进制编码的十进制数 (16位寄存器)
	DataFormatBCD32                          // 二进制编码的十进制数 (32位，占2个寄存器)
	DataFormatStringASCII                    // ASCII字符串 (每个字节一个字符)
	DataFormatStringUTF16                    // UTF-16字符串 (每2字节一个字符)
	DataFormatBits                           // 位图(比特位)，用于线圈、状态字等
	DataFormatBitRange                       // 位范围，用于提取特定范围的位
	DataFormatCustom                         // 自定义格式，使用自定义解析函数
)
