package errid

var (
	UNKNOWN_ERR     = NewMsg(0, "未知错误")
	OK              = NewMsg(1, "成功")
	FAIL            = NewMsg(2, "失败")
	ARG_ERR         = NewMsg(3, "参数错误")
	NOT_CURR_PLAYER = NewMsg(5, "不是当前用户") //这里比较危险，出现这个错误说明有人在利用用户id漏洞攻击服务器

	//room相关
	ROOM_ERR                    = NewMsg(100, "房间错误")
	ROOM_INVALID                = NewMsg(101, "房间无效")
	PLAYER_NOT_IN_THIS_ROOM     = NewMsg(102, "玩家不在房间中")
	NO_SEATS                    = NewMsg(103, "没座位了")
	SEATS_ALREADY_EXIST         = NewMsg(104, "座位已存在")
	PLAYER_NO_SEATS             = NewMsg(105, "玩家没有座位")
	SEAT_NOT_EXIST              = NewMsg(106, "座位不存在")
	PLAYER_ALREADY_SEAT_DOWN    = NewMsg(107, "玩家已经坐下")
	ROOM_NOT_EXIST              = NewMsg(108, "房间不存在")
	CURR_TURN_NOT_BEGIN         = NewMsg(109, "回合尚未开始")
	TARGET_NOT_IN_THIS_ROOM     = NewMsg(110, "对方不在房间")
	IS_THE_LAST_TURN            = NewMsg(111, "最后一局")
	BANKER_NOT_EXIST            = NewMsg(112, "庄家不存在")
	PLAYER_NOT_RIGHT_TO_OPERATE = NewMsg(113, "玩家无操作的权限")
	ROOM_NOT_IN_STATUS          = NewMsg(114, "房间不在状态")
	NO_FACES                    = NewMsg(115, "没得面可以下")
	PLAYER_OFFLINE_LIST_NIL     = NewMsg(116, "房间掉线列表为nil")
	CREATE_ROOM_FAIL            = NewMsg(117, "创建房间失败")
	GET_ROOM_ADDR_FAIL          = NewMsg(118, "拿房间地址失败")
	ROOM_STATUS_ERR             = NewMsg(119, "房间状态错误")
	ROOM_CARD_NUM_ERR           = NewMsg(120, "房卡数量不对")
	RETURN_ROOM_CARD_FAIL       = NewMsg(121, "退还房卡失败")
	ROOM_SEAT_INVALID           = NewMsg(122, "房间位置不可用")

	//数据结构相关
	DATA_STRUCT_ERR = NewMsg(1000, "数据结构错误")
	CHAN_FULL       = NewMsg(1001, "管道已满")
	CHAN_EMPTY      = NewMsg(1002, "管道已空")
	QUEUE_FULL      = NewMsg(1003, "队列满")
	QUEUE_EMPTY     = NewMsg(1004, "队列空")

	//数据库相关
	DB_ERR              = NewMsg(2000, "数据库错误")
	REDIS_CONN_INVALID  = NewMsg(2001, "redis连接无效")
	GET_FROM_REDIS_FAIL = NewMsg(2002, "从redis读取数据失败")
	SET_REDIS_FAIL      = NewMsg(2003, "将数据保存到redis失败")

	//序列化相关
	SERIALIZE_ERR      = NewMsg(3000, "序列化错误")
	JSON_FORMAT_ERR    = NewMsg(3001, "json格式化错误")
	JSON_UNFORMAT_ERR  = NewMsg(3002, "json反格式化错误")
	PROTO_FORMAT_ERR   = NewMsg(3003, "proto格式化错误")
	PROTO_UNFORMAT_ERR = NewMsg(3004, "proto反格式化错误")
	MSG_TOO_SHROT      = NewMsg(3005, "消息过短")
	MSG_TOO_LONG       = NewMsg(3006, "消息过长")
	CONTENT_NIL        = NewMsg(3007, "content为空")
	MSG_NOT_COMPLETE   = NewMsg(3008, "消息不完整")
	MSG_NIL            = NewMsg(3009, "消息为空")

	//player相关
	PLAYER_ERR                     = NewMsg(4000, "玩家错误")
	PLAYER_NOT_EXIST               = NewMsg(4001, "玩家不存在")
	PLAYER_ALREADY_ONLINE          = NewMsg(4002, "玩家已登录")
	PLAYER_CHIP_NOT_ENOUGH         = NewMsg(4003, "玩家筹码不够")
	PLAYER_NOT_SIT_DOWN            = NewMsg(4004, "玩家没坐下")
	UNKNOWN_TEXT_TYPE              = NewMsg(4005, "未知文字消息")
	BLACK_PLAYER                   = NewMsg(4006, "用户在黑名单中")
	TARGET_NOT_SIT_DOWN            = NewMsg(4007, "对方没坐下")
	NO_PLAYER_READY                = NewMsg(4008, "没有玩家准备")
	BI_PAI_FAIL                    = NewMsg(4009, "比牌失败")
	PLAYER_NOT_IN_STATUS           = NewMsg(4010, "玩家不在状态")
	BANKER_NO_RIGHT_TO_BOTTOM_POUR = NewMsg(4011, "庄家没有资格下注")
	PLAYER_INFO_SAVE_FAILE         = NewMsg(4012, "玩家信息保存失败")
	PLAYER_ALREADY_EXIST           = NewMsg(4013, "玩家已存在")
	REG_PLAYER_FAIL                = NewMsg(4014, "注册玩家信息失败")
	PLAYER_TOKEN_ERR               = NewMsg(4015, "玩家token错误")
	PLAYER_NOT_LOGIN               = NewMsg(4016, "玩家未登录")
	PLAYER_NOT_IN_HALL             = NewMsg(4017, "玩家不在大厅")
	PLAYER_ALREADY_IN_THIS_ROOM    = NewMsg(4018, "玩家已经在房间里了")
	PLAYER_IS_FAILED               = NewMsg(4019, "玩家已经被淘汰")

	//step相关
	STEP_ERR = NewMsg(5000, "step错误")
	PREV_NIL = NewMsg(5001, "prev错误")

	//pubsub相关
	PUB_SUB_ERR = NewMsg(6000, "pub sub错误")
	ALREADY_SUB = NewMsg(6001, "已经订阅")

	//poker相关
	POKER_ERR            = NewMsg(7000, "扑克错误")
	POKER_POOL_EMPTY     = NewMsg(7001, "扑克牌池已空")
	POKER_CARD_NOT_EXIST = NewMsg(7002, "此牌不存在")

	//http相关
	HTTP_ERR           = NewMsg(8000, "http错误")
	HTTP_PARSE_ERR     = NewMsg(8001, "http解析资源失败")
	RESOURCE_NOT_EXIST = NewMsg(8002, "资源不存在")

	//服务器相关
	SERVER_ERR          = NewMsg(9000, "服务器错误")
	SERVER_PASSWORD_ERR = NewMsg(9001, "服务器密码错误")
	SERVER_INFO_NIL     = NewMsg(9002, "服务器信息错误")

	//筹码相关
	CHIP_ERR           = NewMsg(10000, "筹码错误")
	CHIP_NUM_TOO_SMALL = NewMsg(10001, "筹码过小")
	CHIP_NUM_ERR       = NewMsg(10002, "筹码数量错误")

	//扎金花相关
	FGF_ERR                     = NewMsg(11000, "扎金花错误")
	FGF_BIPAI_FAIL              = NewMsg(11001, "扎金花比牌失败")
	FGF_FIRST_LOOP_CANNOT_BIPAI = NewMsg(11002, "扎金花第一轮不能比牌")
)
