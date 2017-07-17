package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Cfg *Config

const config_file = "config.json"

func init() {
	Cfg = &Config{}
	buf, err := ioutil.ReadFile(config_file)
	if err != nil {
		fmt.Printf("%s:%s\n", config_file, err)
	}

	err = json.Unmarshal(buf, Cfg)

	if err != nil {
		fmt.Printf("error parsing %s:%s\n", config_file, err)
	}
}

type Config struct {
	EtcdAddr string `json:"etcd_addr"`
	MyAddr   string `json:"my_addr"`
}
