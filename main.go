package main

import (
	"fmt"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm2util"
)

func main() {

	base64 := sm2util.EncryptBase64([]byte(`[{"medins_list_codg":"WC_219114_205274","fixmedins_code":"H51011502705"}]`))
	fmt.Println(base64)

	data := sm2util.SignData([]byte(`{"body":"BNszPlpkAs33EhQPdq+vWVg5V5GicQkV/mMpXafGmUbr1DeA9T/FMB3z9s80PmcxoGtspyYEdPXBruu8Pyz9ehfis/H19gnAL3sjk2zMzyE3WXT0KRTA7MgopwsrBfKDZFBnM4MPX16FtLiUGnPAJlVjS0niRh04Ni74+j4lE7n4lP57tQxdKzTcVYx/7CfDwY/98HsA6uHGFWFgMFZYB61AUGVvZaWP5Sc=","secretKey":"f6ec579e4e51427bbfba260db061582a"}`))
	fmt.Println("sign:", data)

	verifyData := sm2util.VerifyData(`{"body":"BNszPlpkAs33EhQPdq+vWVg5V5GicQkV/mMpXafGmUbr1DeA9T/FMB3z9s80PmcxoGtspyYEdPXBruu8Pyz9ehfis/H19gnAL3sjk2zMzyE3WXT0KRTA7MgopwsrBfKDZFBnM4MPX16FtLiUGnPAJlVjS0niRh04Ni74+j4lE7n4lP57tQxdKzTcVYx/7CfDwY/98HsA6uHGFWFgMFZYB61AUGVvZaWP5Sc=","secretKey":"f6ec579e4e51427bbfba260db061582a"}`,
		`MEUCIQCJaBcbb+kc4Zkst6ccRU92DFLCzppTo93fTzGPp0suBAIgTA588dgsjd4dO1gmRYzIDVYarhTZDdKdy7YOdrhArgU=`)
	fmt.Println(verifyData)

	decodeBase64 := sm2util.DecodeBase64(`BNszPlpkAs33EhQPdq+vWVg5V5GicQkV/mMpXafGmUbr1DeA9T/FMB3z9s80PmcxoGtspyYEdPXBruu8Pyz9ehfis/H19gnAL3sjk2zMzyE3WXT0KRTA7MgopwsrBfKDZFBnM4MPX16FtLiUGnPAJlVjS0niRh04Ni74+j4lE7n4lP57tQxdKzTcVYx/7CfDwY/98HsA6uHGFWFgMFZYB61AUGVvZaWP5Sc=`)

	fmt.Println(string(decodeBase64))

}
