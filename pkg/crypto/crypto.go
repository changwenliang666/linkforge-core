package crypto

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
