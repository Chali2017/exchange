package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"webapi/common/confy"
	"webapi/common/launcher"
	"webapi/common/logi"
	"webapi/metadata"
	"webapi/web.server/routers"
)

var (
	AppName = "web"
	r       *gin.Engine
)

func main() {
	//configUrl := flag.String("config", "config.json", "config filename")
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			logi.Errorf(AppName, "Process stopped - %+v", r)
		} else {
			logi.Infof(AppName, "Process stopped")
		}

		logi.FlushAll()
	}()

	// 打印版本信息
	metadata.PrintConsoleAndExit()

	// 初始化公共模块
	launcher.InitializeAll(AppName)

	// 打印系统基本信息
	metadata.Print(func(k, v string) {
		logi.Infof(AppName, "%s: %s", k, v)
	})

	// 初始化业务层
	/*launcher.LoadModule(AppName, "MyOA Core Service", func() {
		core.Initialize()
	})*/
	//common.InitConfig(*configUrl)

	// 初始化 Web 服务器
	launcher.LoadModule(AppName, "Gin Web Framework", func() {
		gin.SetMode(gin.DebugMode)
		r = routers.GinReady()
	})

	// 刷日志
	{
		logi.Infof(AppName, "GIN Server launched (port:%s)", confy.String("web.http.port"))
		logi.FlushAll()
	}

	// 启动 Web 服务器
	//if err := r.Run(fmt.Sprintf("%s:%s", confy.String("web.http.ip"), confy.String("web.http.port"))); err != nil {
	if err := r.Run(fmt.Sprintf(":%s", confy.String("web.http.port"))); err != nil {
		logi.Infof(AppName, "GIN Server fatal: %s", err.Error())
	}
}
