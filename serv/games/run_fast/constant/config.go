package constant

type RoomConfig struct {
	RoomLifeTimeOut        int `json:"RoomLifeTimeOut"`        //房间寿命(min)
	ReadyTimeOut           int `json:"ReadyTimeOut"`           //准备超时(s)
	PutCardTimeOut         int `json:"PutCardTimeOut"`         //发牌超时(s)
	TurnSettleTimeOut      int `json:"TurnSettleTimeOut"`      //回合结算超时(s)
	PlayerOperationTimeOut int `json:"PlayerOperationTimeOut"` //玩家操作超时(s)
	ReadyWheelingDur       int `json:"ReadyWheelingDur"`       //准备等待扫描间隔(s)
}
