package sm2util

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"os"
)

const (
	pripath = "./sm2.pri.pem"
	pubpath = "./sm2.pub.pem"
)

var (
	pubKey *sm2.PublicKey
	priKey *sm2.PrivateKey
)

func init() {
	initPriKey()
	initPubKey()
}

func initPriKey() {
	file, err := os.ReadFile(pripath)
	if err != nil {
		logger.Logger.Error(err)
	}
	priKey, err = x509.ReadPrivateKeyFromPem(file, nil)
	if err != nil {
		logger.Logger.Error(err)
	}

}
func initPubKey() {
	file, err := os.ReadFile(pubpath)
	if err != nil {
		logger.Logger.Error(err)
	}
	pubKey, err = x509.ReadPublicKeyFromPem(file)
	if err != nil {
		logger.Logger.Error(err)
	}
}

func EncryptBase64(data []byte) string {
	encrypt, err := sm2.Encrypt(pubKey, data, rand.Reader, sm2.C1C2C3)
	logger.Logger.PanicErrorCtx(gctx.New(), err)
	return base64.StdEncoding.EncodeToString(encrypt)
}
func DecodeBase64(data string) []byte {
	decodeString, err2 := base64.StdEncoding.DecodeString(data)
	logger.Logger.PanicErrorCtx(gctx.New(), err2)
	//decrypt, err2 :=priKey.DecryptAsn1(decodeString)
	decrypt, err2 := sm2.Decrypt(priKey, decodeString, sm2.C1C2C3)
	logger.Logger.PanicErrorCtx(gctx.New(), err2)
	return decrypt
}

func SignData(data []byte) string {
	sign, err := priKey.Sign(rand.Reader, data, nil)
	logger.Logger.PanicErrorCtx(gctx.New(), err)
	return base64.StdEncoding.EncodeToString(sign)
}

func VerifyData(msg, sign string) bool {
	msgBytes := []byte(msg)
	//signBytes, _ := hex.DecodeString(sign)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	// sm2验签
	return pubKey.Verify(msgBytes, signBytes)
}
