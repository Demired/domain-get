package main

import (
	domain "domain-get/golang/pull"
	"fmt"
	"time"
)

var (
	token    = "PehsQdzWkBxnwVJr"            //示例token
	version  = "new_dns"                     //new|new_dns|all|all_dns|del
	date     = time.Now().Format("20060102") //最多支持30天新增/删除查询
	savePath = "."
)

func main() {
	e := domain.Pull(token, version, date, savePath)
	if e != nil {
		panic(e)
	}
	fmt.Println("下载完毕")
}
