package help

import (
	"errors"
	"time"

	"github.com/hanjingo/casino/serv/base"
	"github.com/hanjingo/casino/serv/lib"
)

type App struct {
	base.App
}

func NewApp() *App {
	back := new(App)
	back.App.App() //父类构造
	return back
}

func (app *App) Init() {
	app.App.Init()
	//初始化配置
	Config = NewConfig()

	//初始化处理器
	Handler = new(lib.Handler)
}

func (app *App) Run() {
	app.App.Run()
	dur := time.Duration(Config.ScanDur) * time.Millisecond
	//向网关拨号
	if app.Session == nil {
		panic(errors.New("没有设置网关客户端"))
	}
	app.Session.Bind(Config.Id)
	if err := app.Session.Dial(Config.GateAddr, Config.GateToken, Config.Net); err != nil {
		panic(err)
	}
	//启动其他组件
	app.PlayerM.Run(dur)
	app.RoomM.Run(dur)
}
