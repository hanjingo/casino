package base

type App struct {
	BRun bool //是否已启动
}

//子游戏服务器初始化
func (app *App) Init() {
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
