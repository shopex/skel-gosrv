package main

import (
	"git.shopex.cn/luban/lib"
	"log"
)

func main() {
	cmd_runner := &lib.Commands{
		cmdServer,
	}
	log.SetFlags(log.LstdFlags | log.Llongfile)

	cmd_runner.Run()
}
