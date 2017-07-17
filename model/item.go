package model

// 使用下面命令生成Controller
// lsv generate tmc model/item.go tmc

type Item struct {
	Id      int64
	Name    string
	Created int64
	Updated int64
}
