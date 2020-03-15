package base

import (
	"github.com/hanjingo/casino/serv/lib"
)

type App struct {
	BRun    bool           //是否已启动
	PlayerM *PlayerManager //玩家管理器
	RoomM   *RoomManager   //房间管理器
	Session *lib.Session   //根session
}

//构造函数
func (app *App) App() {
	app.BRun = false
	app.PlayerM = new(PlayerManager)
	app.PlayerM.PlayerManager()

	app.RoomM = new(RoomManager)
	app.RoomM.RoomManager()

	app.Session = lib.NewSession()
}

//子游戏服务器初始化
func (app *App) Init() {
	//初始化玩家管理器
	app.PlayerM.Init()

	//初始化房间管理器
	app.RoomM.Init()

	//具体逻辑去子游戏中实现
}

//子游戏服务器跑起来
func (app *App) Run() {
	if app.BRun {
		return
	}
	app.BRun = true
	//具体逻辑去子游戏中实现
}
