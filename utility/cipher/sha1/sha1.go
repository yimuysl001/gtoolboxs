package sha

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

func ToUpperSHA1(s string) string {
	return strings.ToUpper(SHA1(s))

}
