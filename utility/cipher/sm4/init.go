package sm4

import (
	"context"
	"encoding/base64"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm4/ymsm4"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

const (
	pempath = "key.pem"          // 密文路径
	keystr  = "AC1hI4j5kLmPQr8S" //默认密文
)

var sm4key ymsm4.SM4Key

func init() {
	readPem()
}

// 写入密文
func WritePem(key []byte) {
	logger.Logger.PanicErrorCtx(context.Background(), ymsm4.WriteKeyToPemFile(pempath, key, nil), "写文件失败")
	readPem()
}

// 读取密文
func readPem() {
	var err error
	sm4key, err = ymsm4.ReadKeyFromPemFile(pempath, nil)
	//logger.IfError(err)
	if err != nil {
		sm4key = []byte(keystr)
		//ymsm4.WriteKeyToPemFile(pempath, sm4key, nil)
	}
}

// EncryptEcbBase64 加密
func EncryptEcbBase64(data []byte) string {
	ecbMsg, err2 := ymsm4.Sm4Ecb(sm4key, data, true)
	logger.Logger.PanicErrorCtx(context.Background(), err2)
	return base64.StdEncoding.EncodeToString(ecbMsg)
}

// DectyptEcbBase64 解密
func DectyptEcbBase64(data string) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(data)
	ecbMsg, err2 := ymsm4.Sm4Ecb(sm4key, decodeString, false)
	logger.Logger.PanicErrorCtx(context.Background(), err2)
	return ecbMsg
}

// EncryptByKeyEcbBase64 加密
func EncryptByKeyEcbBase64(data []byte, key string) string {
	ecbMsg, err2 := ymsm4.Sm4Ecb([]byte(key), data, true)
	logger.Logger.PanicErrorCtx(context.Background(), err2)
	return base64.StdEncoding.EncodeToString(ecbMsg)
}

// DectyptByKeyEcbBase64 解密
func DectyptByKeyEcbBase64(data, key string) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(data)
	ecbMsg, err2 := ymsm4.Sm4Ecb([]byte(key), decodeString, false)
	logger.Logger.PanicErrorCtx(context.Background(), err2)
	return ecbMsg
}
