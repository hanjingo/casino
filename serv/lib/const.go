package lib

//游戏id
const (
	GAME_ID_TP uint32 = 10000 //赢三张
)

//房间状态
const (
	ROOM_STATUS_CREATED     uint32 = 1 //已创建
	ROOM_STATUS_GAMING      uint32 = 2 //游戏中
	ROOM_STATUS_CLOSED      uint32 = 3 //已关闭
	ROOM_STATUS_AUTO_CLOSED uint32 = 4 //自动关闭
	ROOM_STATUS_ERR         uint32 = 5 //房间状态异常
)

//玩家状态
const (
	PLAYER_STATUS_OFF_LINE      uint32 = 1 //已离线
	PLAYER_STATUS_LOGIN         uint32 = 2 //已登录
	PLAYER_STATUS_GAMING        uint32 = 3 //游戏中
	PLAYER_STATUS_DIS_CONNECTED uint32 = 4 //断线
)
