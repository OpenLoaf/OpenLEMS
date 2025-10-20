package t_machine_id

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

// GetMachineId 获取10位机器ID
// 使用机器硬件ID生成不可逆的10位十六进制字符串
func GetMachineId() string {
	// 获取原始机器ID
	rawId, err := machineid.ID()
	if err != nil {
		// 如果获取失败，返回一个基于错误信息的固定ID
		hash := md5.Sum([]byte("fallback-machine-id"))
		id := fmt.Sprintf("%x", hash)[:10]
		log.Printf("获取机器ID失败: %v, 使用ID: %v", err, id)
		return id
	}

	// 使用MD5哈希将长ID转换为10位十六进制字符串
	hash := md5.Sum([]byte(rawId))
	return fmt.Sprintf("%x", hash)[:10]
}

// GetProtectedMachineId 获取受保护的10位机器ID
// 使用应用特定的密钥对机器ID进行HMAC-SHA256哈希
func GetProtectedMachineId(appId string) string {
	// 获取受保护的机器ID
	protectedId, err := machineid.ProtectedID(appId)
	if err != nil {
		log.Printf("获取受保护机器ID失败: %v", err)
		// 如果获取失败，返回一个基于错误信息的固定ID
		hash := md5.Sum([]byte("fallback-protected-id"))
		return fmt.Sprintf("%x", hash)[:10]
	}

	// 使用MD5哈希将长ID转换为10位十六进制字符串
	hash := md5.Sum([]byte(protectedId))
	return fmt.Sprintf("%x", hash)[:10]
}
