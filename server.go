package main

import (
	"git.shopex.cn/luban/lib"
	"log"
	"runtime"
	"strings"
)

var cmdServer = &lib.Command{
	UsageLine: "server",
	Short:     "服务管理",
	Long: `
    start           启动服务
`,
}

func init() {
	cmdServer.Run = server
}

func server(cmd *lib.Command, args []string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(args) < 1 {
		log.Println("缺少参数")
		return 1
	}
	switch args[0] {
	case "start":
		lib.ConnectEtcd("@{SKEL}", Cfg.EtcdAddr, Cfg.MyAddr)
		token := strings.Split(Cfg.MyAddr, ":")
		lib.ConnectDatabase()
		return lib.StartService(":"+token[1], Apis, &ApiDefine)
	}
	return 0
}
