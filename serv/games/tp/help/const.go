package help

import (
	"github.com/hanjingo/casino/serv/lib"
	"github.com/hanjingo/logger"
)

var GameApp *App //整个进程的根变量

var Handler *lib.Handler //回调器

var Config *TpConfig //配置

var log = *logger.GetDefaultLogger() //使用默认的日志

//房间状态
const (
	GAMING_STATUS_UNKNOW uint32 = 0 //未知状态

	GAMING_STATUS_START uint32 = 1 //房间开始

	GAMING_STATUS_READYING   uint32 = 2 //正在准备
	GAMING_STATUS_READY_DONE uint32 = 3 //准备完成

	GAMING_STATUS_PUT_CARDING   uint32 = 4 //正在发牌
	GAMING_STATUS_PUT_CARD_DONE uint32 = 5 //发牌完成

	GAMING_STATUS_LOOPING   uint32 = 6 //正在轮转
	GAMING_STATUS_LOOP_DONE uint32 = 7 //轮转完成

	GAMING_STATUS_TURN_SETTLEING   uint32 = 8 //正在回合结算
	GAMING_STATUS_TURN_SETTLE_DONE uint32 = 9 //回合结算完成

	GAMING_STATUS_ROOM_DONE uint32 = 10 //房间关闭
)
