package permission2

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm4"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/timeutil"
	"os"
	"strings"
	"time"
)

const permissionpath = "./config/.permission"

var (
	// 用于控制流程是否向下开展
	flags = false
	//获取key
	ckey = ""
	//加密字符
	ctime = ""
	//有效期
	dtime = time.Now()
)

// GetKey 获取程序数量的key
func GetKey() string {
	if ckey != "" {
		return ckey
	}
	name, _ := os.Hostname()
	dir, _ := os.Getwd()
	ckey = name + " " + dir
	ckey = gmd5.MustEncryptBytes([]byte(ckey))
	ckey = ckey[:16]
	return ckey
}

// SetCtime 设置权限日期
func SetCtime(str string) (errn error) {
	str = strings.TrimSpace(str)

	if str == "请输入注册码" || str == "" {
		return errors.New("未设置注册码")
	}
	if str == ckey {
		dtime = time.Now().AddDate(0, 0, 1)
		logger.Logger.Trace("临时权限")
		return nil
	}
	defer func() {
		if err := recover(); err != nil {
			logger.Logger.Error(err)
			errn = errors.New(fmt.Sprintf("%v", err))
		}
	}()

	base64 := sm4.DectyptEcbBase64(str)
	//key 解密
	ecbBase64 := sm4.DectyptByKeyEcbBase64(string(base64), ckey)
	dtime = timeutil.StrToTime(string(ecbBase64))
	if dtime.IsZero() {
		return errors.New("注册码输入不正确！！！")
	}

	if dtime.Before(time.Now()) {
		return errors.New("设置有效期:" + dtime.String() + "，在当前时间之前。")
	}

	return os.WriteFile(permissionpath, []byte(str), 0666)

}

func GetTime() time.Time {
	return dtime
}

// 获取有效期加密字符
func getCtime() {
	GetKey()
	if ctime != "" {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			//解密失败
			ctime = ""
		}
	}()
	ctime = ""
	file, err := os.ReadFile(permissionpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := strings.TrimSpace(string(file))

	if c == ckey {
		dtime = time.Now().AddDate(0, 0, 1)
		return
	}
	// TODO:自带解密
	base64 := sm4.DectyptEcbBase64(c)
	//key 解密
	ecbBase64 := sm4.DectyptByKeyEcbBase64(string(base64), ckey)

	dtime = timeutil.StrToTime(string(ecbBase64))

	if dtime.Before(time.Now()) {
		fmt.Println("程序已到期:" + dtime.String() + "，请联系管理员:" + ckey)
		ctime = ""
	}
	ctime = c
	return
}

func GetCtime() (t time.Time, err error) {
	defer func() {
		if r := recover(); r != nil {
			//解密失败
			err = fmt.Errorf("%v", r)
		}
	}()
	GetKey()
	file, err := os.ReadFile(permissionpath)
	c := strings.TrimSpace(string(file))
	if c == ckey {
		dtime = time.Now().AddDate(0, 0, 1)
		return dtime, nil
	}
	// TODO:自带解密
	base64 := sm4.DectyptEcbBase64(c)
	//key 解密
	ecbBase64 := sm4.DectyptByKeyEcbBase64(string(base64), ckey)

	dtime = timeutil.StrToTime(string(ecbBase64))

	if dtime.Before(time.Now()) {
		//fmt.Println("程序已到期:" + dtime.String() + "，请联系管理员:" + ckey)
		ctime = ""
		return dtime, fmt.Errorf("程序已到期:%v，请联系管理员: %v", dtime, ckey)

	}
	ctime = c
	return dtime, nil
}
