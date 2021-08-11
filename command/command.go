package baidupcscmd

import (
	"os"

	"github.com/iikira/BaiduPCS-Go/baidupcs"
	pcsconfig "github.com/iikira/BaiduPCS-Go/config"
)

var (
	info = new(baidupcs.PCSApi)
)

func init() {
	ReloadInfo()
}

// ReloadInfo 重载配置
func ReloadInfo() {
	pcsconfig.Reload()
	if pcsconfig.ActiveBaiduUser != nil {
		info = baidupcs.NewPCS(pcsconfig.ActiveBaiduUser.BDUSS)
	}
}

// ReloadIfInConsole 程序在 Console 模式下才回重载配置
func ReloadIfInConsole() {
	if len(os.Args) == 1 {
		ReloadInfo()
	}
}
