package hashx

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// MD5String 密码加密
func MD5String(oPassword string) string {
	h := md5.New()
	h.Write([]byte("xiangshou.md5.test.123sxbi1928bd~!@#"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
