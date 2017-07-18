package main

import (
	"git.shopex.cn/luban/lib"
)

func main() {
	cmd_runner := &lib.Commands{
		cmdServer,
	}

	cmd_runner.Run()
}
