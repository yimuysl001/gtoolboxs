package autolnk

import (
	"fmt"
	"github.com/jxeng/shortcut"
	"github.com/yimuysl001/gtoolboxs/utility/fileutil"
	"github.com/yimuysl001/gtoolboxs/utility/sys"
	"os/user"
	"strings"
)

const startup = "%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\%s.lnk"

// 开机自启动快捷方式
const desktopPth = "%s\\Desktop\\%s.lnk"

func GetHome() string {

	current, err := user.Current()
	if err != nil {
		return "C:\\Users\\User"
	}

	return current.HomeDir

}

func GetSysName() string {
	return strings.ReplaceAll(strings.ToLower(sys.GetSysName()), ".exe", "")
}

func CreateDesktop() error {
	dpath := fmt.Sprintf(desktopPth, GetHome(), GetSysName())
	exists, err := fileutil.PathExists(dpath)
	if err == nil && exists {
		return err
	}
	return shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     dpath,
		Target:           sys.GetAllSysName(),
		IconLocation:     "%SystemRoot%\\System32\\SHELL32.dll,0",
		Arguments:        "",
		Description:      "",
		Hotkey:           "",
		WindowStyle:      "1",
		WorkingDirectory: "",
	})

}

func CreateStartup() error {
	dpath := fmt.Sprintf(startup, GetHome(), GetSysName())
	exists, err := fileutil.PathExists(dpath)
	if err == nil && exists {
		return err
	}
	return shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     dpath,
		Target:           sys.GetAllSysName(),
		IconLocation:     "%SystemRoot%\\System32\\SHELL32.dll,0",
		Arguments:        "",
		Description:      "",
		Hotkey:           "",
		WindowStyle:      "1",
		WorkingDirectory: "",
	})

}
