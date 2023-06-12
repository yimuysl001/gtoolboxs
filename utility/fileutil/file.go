/**
 * fileutil
 * @author : yangsonglin
 * @date : 2023-03-09
 */
package fileutil

import (
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"os"
	"path"
	"strings"
)

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

// CreateDir 创建文件夹
func CreateDir(path string) {
	_exist, _err := PathExists(path)
	if _err != nil {
		logger.Logger.Error("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		logger.Logger.Trace("文件夹已存在:" + path)
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Logger.Error("创建目录异常 -> %v\n", err)
		} else {
			logger.Logger.Trace(path, "创建成功!")
		}
	}
}

// GetFileSuffix 获取文件后缀名称
func GetFileSuffix(pathname string) string {
	return strings.ToLower(path.Ext(pathname))
}

// GetFileName 获取文件名称
func GetFileName(pathname string) string {
	return strings.ToLower(path.Base(pathname))
}
