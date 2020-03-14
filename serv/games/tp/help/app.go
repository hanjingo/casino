package help

var GameApp *App

type App struct {
	Room   *RoomManager
	Player *PlayerManager
}

func GetApp() *App {
	return GameApp
}

func (app *App) Init() {

}

func (app *App) Run() {

}
