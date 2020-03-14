package main

import (
	"github.com/hanjingo/casino/serv/games/tp/handle"
	v1 "github.com/hanjingo/gate/cli/golang/v1"
)

// for win:   go build -o tp.exe main.go
// for linux: go build -o tp main.go
func main() {
	//加载配置

	//初始化日志

	//初始化回调
	handle.Gate = v1.NewGateCli()
	handle.Init()

	//跑起来
}
