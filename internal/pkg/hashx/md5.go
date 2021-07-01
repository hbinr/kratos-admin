package hashx

import (
	"crypto/md5"
	"encoding/hex"
	"kratos-admin/internal/pkg/constant/key"
)

// MD5String 密码加密
func MD5String(oPassword string) string {
	h := md5.New()
	h.Write([]byte(key.KeyMySecret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
