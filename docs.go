package main

import (
	"encoding/json"

	"git.shopex.cn/luban/lib/apipkg"
)

const (
    api_string string = `{
    "apis": [
        {
            "name": "location_list",
            "title": "列出位置",
            "summary": "district_id, city_id, province_id 填写一个即可. 参数有优先级顺序.",
            "parameters": [
                {
                    "name": "district_id",
                    "description": "地区id",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "city_id",
                    "description": "城市id",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "province_id",
                    "description": "省份id",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "location_add",
            "title": "添加位置",
            "parameters": [
                {
                    "name": "name",
                    "description": "名称",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": true,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "district_id",
                    "description": "地区",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "max_x",
                    "description": "经度",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "max_y",
                    "description": "纬度",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "memo",
                    "description": "备注",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "location_remove",
            "title": "删除位置",
            "parameters": [
                {
                    "name": "location_id",
                    "description": "地区id",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "location_provinces",
            "title": "获取省份列表"
        },
        {
            "name": "location_citys",
            "title": "获取城市列表",
            "parameters": [
                {
                    "name": "province_id",
                    "description": "省份id",
                    "type": "int",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "location_districts",
            "title": "获取地区列表",
            "parameters": [
                {
                    "name": "city_id",
                    "description": "城市id",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "captcha_new",
            "title": "生成验证码"
        },
        {
            "name": "captcha_image",
            "title": "获取验证码图片",
            "parameters": [
                {
                    "name": "captcha_id",
                    "description": "验证码id",
                    "type": "int64",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "captcha_refresh",
            "title": "刷新验证码",
            "parameters": [
                {
                    "name": "captcha_id",
                    "description": "验证码id",
                    "type": "int64",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        },
        {
            "name": "captcha_validate",
            "title": "校验验证码",
            "parameters": [
                {
                    "name": "captcha_id",
                    "description": "验证码id",
                    "type": "int64",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                },
                {
                    "name": "input",
                    "description": "输入字符",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        }
    ]
}`
)

var ApiDefine apipkg.ApiDeclaration

func init() {
	json.Unmarshal([]byte(api_string), &ApiDefine)
}
