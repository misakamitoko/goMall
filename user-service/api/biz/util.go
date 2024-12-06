package biz

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func GenerateId() (int64, error) {
	// 定义 10^8 的最大值，用于生成随机数
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(8), nil)

	// 使用安全随机数生成器生成一个随机数
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassWord(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
