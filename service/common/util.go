package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateUID() string {
	buf := make([]byte, 10)

	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%d%x", time.Now().Unix(), buf[0:10])
	return base64.StdEncoding.EncodeToString([]byte(str))
}
