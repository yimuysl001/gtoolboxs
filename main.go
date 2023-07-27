package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm3"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm4/ymsm4"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"sort"
	"strings"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())
	//
	//all, err := dbutil.DB("log").Model("TBZDBQ").All()
	//logger.Logger.Info(err)
	//logger.Logger.Info(all)
	//
	//all, err = dbutil.DB("log").Model("TBZDBQ").All()

	//automatic.Auto(func() {
	//	fmt.Println("=======设置完成======")
	//})

	//logger.Logger.Info("===================")
	//logger.Logger.Info("========自动开机设置测试===========")

	//m := gmap.New()
	//
	//m.Set("id", 4)
	//m.Set("name", "de2")
	//
	//dbutil.DB().Model("test").Insert(m)

	var cstr = `e803b69e92995e6d8752e768d213a661a4dc1fdbd1f689c746e35795e64d0bda3b24acf7081f5790e69900afaeb6ba6743a75aa5fc075e46b666772f7aa9bf733928e4a6200567838450f8d15f22b09a0974326c23d5ccd1133810035e201d862cc6ac802abe7993385302331ee5756beb978af7ff9d1afb91019fb7bb780231ec9cc4a244622618e048ff47ccd312aa92cf6b4dc72afd880aec768bc641201cdcd3d6776fac84af9f11e4661e967c89e7e26a1410cc0c1d2edac88c87adfae6969469eb2158d50fcfdc8f99f7be5d5d4446a4b4dafeb288b2b53dd176b9b40a067e841eb1bc2a6c1fec0876cdb43ba32295728e7d1a67ae85fb09e299f822cb`
	decodeString, err := hex.DecodeString(cstr)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	out, err := ymsm4.Sm4Ecb([]byte("7bf0acf801ae4ff1"), decodeString, false)
	fmt.Println(string(out), err)
	sm3str := `ciphertext=e803b69e92995e6d8752e768d213a661a4dc1fdbd1f689c746e35795e64d0bda3b24acf7081f5790e69900afaeb6ba6743a75aa5fc075e46b666772f7aa9bf733928e4a6200567838450f8d15f22b09a0974326c23d5ccd1133810035e201d862cc6ac802abe7993385302331ee5756beb978af7ff9d1afb91019fb7bb780231ec9cc4a244622618e048ff47ccd312aa92cf6b4dc72afd880aec768bc641201cdcd3d6776fac84af9f11e4661e967c89e7e26a1410cc0c1d2edac88c87adfae6969469eb2158d50fcfdc8f99f7be5d5d4446a4b4dafeb288b2b53dd176b9b40a067e841eb1bc2a6c1fec0876cdb43ba32295728e7d1a67ae85fb09e299f822cbencrypt_type=SM4hospital_code=HNZYCSYY_001request_id=b029bbf4-cb9b-4ae1-aacc-36b32fe32e5csign_type=SM3timestamp=1686188289031&sign_key=6fd6d7117f15434c9237e269a640b0f93ad31e3d89004f23a31c6c694add9a09`

	encrypt := sm3.Encrypt([]byte(sm3str))

	fmt.Println(encrypt)
	sjson := `{
"request_id": "b029bbf4-cb9b-4ae1-aacc-36b32fe32e5c",
"hospital_code": "HNZYCSYY_001",
"timestamp": 1686188289031,
"sign_type": "SM3",
"encrypt_type": "SM4",
"ciphertext":
"e803b69e92995e6d8752e768d213a661a4dc1fdbd1f689c746e35795e64d0bda3b24acf7081f5790e69900afaeb6ba6743a75aa5fc075e46b666772f7aa9bf733928e4a6200567838450f8d15f22b09a0974326c23d5ccd1133810035e201d862cc6ac802abe7993385302331ee5756beb978af7ff9d1afb91019fb7bb780231ec9cc4a244622618e048ff47ccd312aa92cf6b4dc72afd880aec768bc641201cdcd3d6776fac84af9f11e4661e967c89e7e26a1410cc0c1d2edac88c87adfae6969469eb2158d50fcfdc8f99f7be5d5d4446a4b4dafeb288b2b53dd176b9b40a067e841eb1bc2a6c1fec0876cdb43ba32295728e7d1a67ae85fb09e299f822cb"
}`
	m := gjson.New(sjson).Map()
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(gconv.String(m[key]))
	}
	sign_key := "6fd6d7117f15434c9237e269a640b0f93ad31e3d89004f23a31c6c694add9a09"
	sb.WriteString("&sign_key=")
	sb.WriteString(sign_key)

	fmt.Println(sb.String())
	encrypt = sm3.Encrypt([]byte(sm3str))

	fmt.Println(encrypt)

	jsont := `{"name": "泰老师","id_card_no": "110101199003072551","doctor_code": "888888","doctor_name": "GLY","department": "感染科病区","start_time": "2023-04-25 11:24:12","end_time": "2023-07-24 11:24:12","hospitalItems": ["002501010150000"]}`
	out, err = ymsm4.Sm4Ecb([]byte("70e296a276a94ade"), []byte(jsont), true)
	fmt.Println(err)
	logger.Logger.Info(hex.EncodeToString(out))

	select {}
}
