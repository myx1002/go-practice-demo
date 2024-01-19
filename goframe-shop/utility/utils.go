package utility

import (
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// 密码加密
func EncrypPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

func GenerateOrderNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return gconv.String(gtime.TimestampNanoStr() + gconv.String(r.Intn(1000)))
}
