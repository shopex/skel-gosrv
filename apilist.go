package main

import (
	"git.shopex.cn/luban/lib/apipkg"
	"git.shopex.cn/luban/@{SKEL}/controller"
)

var Apis []interface{}

func init() {
	Apis = apipkg.ApiList(
		apipkg.ApiInclude(&controller.Item{}),
	)
}
