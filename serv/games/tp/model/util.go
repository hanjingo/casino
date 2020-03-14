package model

//玩家回合结算信息
type PlayerTurnSettleInfo struct {
	UserId int64 //用户id
	UserName string //用户名
	HeadImgUrl string //头像
	ChangeScore int32 //分数变更
	IsWinner bool //是否是赢家
	Cards []int //牌集合
}

//房间信息
type RoomInfo struct {
	Id int64 //房间id
	Status int64 //房间状态
	GameId int32 //游戏id
	MaxPlayerNum int32 //最大人数限制
	TotalTurn int32 //总回合数
	ChipProjId uint32 //筹码方案id
}

//玩家信息
type PlayerInfo struct {
	Id int64 //玩家id
	Status uint32 //玩家状态
	SeatNum uint32 //座位号
	Name string //名字
	Score int32 //分数
	HeadImgUrl string //头像
}

//玩家回合结算信息
type PlayerTotalSettleInfo struct {
	UserId int64 //用户id
	HeadImgUrl string //头像
	UserName string //用户名
	Tag string //标签
	ChangeScore int32 //分数变更
}
