package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncodeSha1(value string) string {
	s := sha1.New()
	s.Write([]byte(value))
	return hex.EncodeToString(s.Sum(nil))
}
