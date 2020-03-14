package main

import (
	"github.com/hanjingo/casino/serv/games/tp/help"
)

// for win:   go build -o tp.exe main.go
// for linux: go build -o tp main.go
func main() {
	//初始化
	help.GameApp = new(help.App)
	help.GameApp.Init()

	//跑起来
	help.GameApp.Run()
}
