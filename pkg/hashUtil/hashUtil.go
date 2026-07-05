package hashUtil

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const defaultCost = 12

// 密码加密
func HashEncode(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		defaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("密码加密错误")
	}
	return string(hashedBytes), nil
}

// 检查密码是否正确
func CheckHashDecode(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// base62 编码
func base62Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	base := uint64(len(charset))
	var result []byte

	for num > 0 {
		rem := num % base
		result = append([]byte{charset[rem]}, result...)
		num = num / base
	}

	return string(result)
}

// 你要的核心方法：id + longURL → shortCode
func GenerateShortCode(id uint64, longURL string) string {
	// longURL 在这里不参与编码，只作为业务语义保留
	// （未来可用于加盐、防冲突、或自定义短链）
	const offset = 100000
	num := id + offset

	code := base62Encode(num)

	// 可选：控制长度（防止极端情况下过长）
	if len(code) > 10 {
		code = code[:10]
	}

	return code
}
