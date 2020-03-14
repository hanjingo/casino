package help

import (
	"time"

	"github.com/hanjingo/casino/serv/base"
	"github.com/hanjingo/casino/serv/lib"
	"github.com/hanjingo/gate/cli/golang/v1"
)

type App struct {
	base.App
	Session *lib.Session        //根session
	PlayerM *base.PlayerManager //玩家管理器
	RoomM   *base.RoomManager   //房间管理器
}

func (app *App) Init() {
	app.App.Init()
	//初始化配置
	Config = NewConfig()

	//初始化处理器
	Handler = new(lib.Handler)

	//初始化玩家管理器
	app.PlayerM.Init()

	//初始化房间管理器
	app.RoomM.Init()
}

func (app *App) Run() {
	app.App.Run()
	dur := time.Duration(Config.ScanDur) * time.Millisecond
	//向网关拨号
	app.Session = lib.NewSession(v1.NewGateCli())
	app.Session.Bind(Config.Id)
	if err := app.Session.Dial(Config.GateAddr, Config.GateToken, Config.Net); err != nil {
		panic(err)
	}
	//启动其他组件
	app.PlayerM.Run(dur)
	app.RoomM.Run(dur)
}
