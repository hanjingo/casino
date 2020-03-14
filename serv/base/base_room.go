package base

import "time"

//游戏房间设计模板
type Room struct {
	Id         int64     //房间id
	Status     int       //房间状态
	CreateTime time.Time //创建时间
	CloseTime  time.Time //关闭时间
	GameId     int       //游戏类型
	PgId       int       //游戏服务器id
	BRun       bool      //房间是否已启动
}

//拿id
func (room *Room) GetId() interface{} {
	return room.Id
}

//初始化房间
func (room *Room) Init() {
	room.CreateTime = time.Now()
	//具体逻辑去子游戏中实现
}

//让房间跑起来
func (room *Room) Run() {
	//具体逻辑去子游戏中实现
}

//关闭房间
func (room *Room) Close() {
	room.CloseTime = time.Now()
	//具体逻辑去子游戏中实现
}

//结算房间
func (room *Room) Settle() {
	//具体逻辑去子游戏中实现
}

//房间是否可用
func (room *Room) IsValid() bool {
	// todo
	return true
}
