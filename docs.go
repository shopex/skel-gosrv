package main

import (
	"encoding/json"

	"git.shopex.cn/luban/lib/apipkg"
)

const (
	api_string string = ``
)

var ApiDefine apipkg.ApiDeclaration

func init() {
	json.Unmarshal([]byte(api_string), &ApiDefine)
}
