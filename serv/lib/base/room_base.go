package lib

import "time"

//游戏房间设计模板
type RoomBase struct {
	id         int64     //房间id
	status     int       //房间状态
	createTime time.Time //创建时间
	closeTime  time.Time //关闭时间
	gameId     int       //游戏类型
	pgId       int       //游戏服务器id
}

func (base *RoomBase) GetId() int64 {
	return base.id
}

func (base *RoomBase) SetId(id int64) {
	base.id = id
}

func (base *RoomBase) SetStatus(status int) {
	base.status = status
}

func (base *RoomBase) GetStatus() int {
	return base.status
}

func (base *RoomBase) SetCreateTime(t time.Time) {
	base.createTime = t
}

func (base *RoomBase) GetCreateTime() time.Time {
	return base.createTime
}

func (base *RoomBase) SetCloseTime(t time.Time) {
	base.closeTime = t
}

func (base *RoomBase) GetCloseTime() time.Time {
	return base.closeTime
}

func (base *RoomBase) GetGameId() int {
	return base.gameId
}

func (base *RoomBase) SetGameId(id int) {
	base.gameId = id
}

func (base *RoomBase) GetPgId() int {
	return base.pgId
}

func (base *RoomBase) SetPgId(id int) {
	base.pgId = id
}

//房间跑起来
func (base *RoomBase) Run() {
}
