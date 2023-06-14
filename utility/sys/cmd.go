// Package sys
/*
程序名称 程序pid 程序重启 目前只支持windows
*/
package sys

import (
	"bytes"
	"context"
	"fmt"
	"github.com/yimuysl001/gtoolboxs/utility/logger"

	"os"
	"os/exec"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

// 程序名称
var sysName string = ""

// 程序pid
var pid int = -1

// 程序路径
var allsysname = ""

// 初始化
func init() {
	//获取程序路径
	executable, err := os.Executable()

	//错误校验
	logger.Logger.PanicErrorCtx(context.Background(), err)
	pid = os.Getpid()
	allsysname = strings.ReplaceAll(executable, "\\", "/")
	//获取文件名称
	sysName = path.Base(allsysname)

	logger.Logger.Info("============\t", "程序名称：", sysName, "\t============")
	logger.Logger.Info("============\t", "PID:", pid, "\t============")

}

func GetAllSysName() string {
	return allsysname
}

func GetPid() int {
	return pid
}

func GetSysName() string {
	return sysName
}

// Start 启动程序
// 默认为 sysName 如果入参有值 为入参第一个值
func Start(sysname ...string) {
	name := strings.TrimSpace(sysName)
	if len(sysname) > 0 && strings.TrimSpace(sysname[0]) != "" {
		name = strings.TrimSpace(sysname[0])
	}
	command := exec.Command("cmd", "/c", "start", name)
	err := command.Start()
	logger.Logger.PanicErrorCtx(context.Background(), err)
	exit()

}

// CheckSys 校验程序是否可用
// 如果存在相同文件不同pid 直接退出程序
func CheckSys(sysname ...string) bool {

	defer func() {
		logger.Logger.Info("同名文件校验完成......")
	}()

	logger.Logger.Info("校验同名文件......")
	//避免旧程序未能及时关闭
	time.Sleep(time.Second * 2)
	name := strings.TrimSpace(sysName)
	if len(sysname) > 0 && strings.TrimSpace(sysname[0]) != "" {
		name = strings.TrimSpace(sysname[0])
	}
	//pid不等于当前程序pid
	pidstr := fmt.Sprintf("PID NE  %d", pid)
	if len(sysname) > 1 && strings.TrimSpace(sysname[1]) != "" {
		pidstr = fmt.Sprintf("PID NE  %v", strings.TrimSpace(sysname[1]))
	}
	//校验同名文件 列出文件名称等于当前程序名称的程序，pid不等于当前程序的进程
	cmd := exec.Command("TASKLIST", "/FI", "IMAGENAME eq "+name, "/FI", pidstr)
	output, err := cmd.Output()
	if err != nil {
		logger.Logger.Error(err)
		exit()
		return false
	}
	//字符编码不一致，内容可能出现乱码，所以校验 ======
	if bytes.Contains(output, []byte("====")) {
		logger.Logger.Info("已有相同名称程序启动：", name)
		exit()
		return true

	}
	getkillsysbat()
	return false

}

func getkillsysbat() {
	os.WriteFile("killsys.bat",
		[]byte(fmt.Sprintf("taskkill /f /im %d \n del /f /s /q killsys.bat", GetPid())),
		0666)

}

// BlockingThread 守护线程
/*用法示例
func main() {

    //异常退出重启
	defer func() {
		mylog.IfError(context.Background(), recover())
		sys.Start()

	}()

   	//线程启动需要使用的程序
	go cmd.Main.Run(gctx.New())

	//阻塞程序，并在程序将要退出的时候调用
	sys.BlockingThread()

}
*/
func BlockingThread(name ...string) {
	logger.Logger.Info("程序启动完成，守护线程等待...")
	sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//signal.Notify(sigs, syscall.Signal(17), syscall.Signal(19, 23)
	//signal.Notify(sigs)
	yy := <-sigs
	logger.Logger.Error("退出原因：", yy)

	Start(name...)
	//exit()
}

// 退出程序
func exit() {
	os.Exit(0)
}
