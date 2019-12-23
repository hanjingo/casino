//登录
export class LoginReq{
    public UserId:number;
    public Token:string;
}
export class LoginRsp{
    public UserId:number;
    public Result:number;
}

//进房
export class JoinRoomReq{
    public UserId:number;
    public RoomId:number;
}
export class JoinRoomRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
    public RoomInfo:RoomInfo;
}
export class JoinRoomNtf{
    public RoomId:number;
    public UserId:number;
}
export class RoomInfo{
    //todo
}

//退房
export class ExitRoomReq{
    public RoomId:number;
    public UserId:number;
}
export class ExitRoomRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
}
export class ExitRoomNtf{
    public RoomId:number;
    public UserId:number;
}

//坐下
export class SeatDownReq{
    public RoomId:number;
    public SeatId:number;
    public UserId:number;
}
export class SeatDownRsp{
    public Result:number;
    public RoomId:number;
    public SeatId:number;
    public UserId:number;
}
export class SeatDownNtf{
    public RoomId:number;
    public SeatId:number;
    public UserId:number;
}

//起身
export class GetUpReq{
    public RoomId:number;
    public UserId:number;
}
export class GetUpRsp{
    public RoomId:number;
    public UserId:number;
    public SeatId:number;
}
export class GetUpNtf{
    public RoomId:number;
    public UserId:number;
}

//准备
export class ReadyReq{
    public RoomId:number;
    public UserId:number;
}
export class ReadyRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
}
export class ReadyNtf{
    public RoomId:number;
    public UserId:number;
}

//加注
export class FillingChipsReq{
    public RoomId:number;
    public UserId:number;
    public Num:number;
}
export class FillingChipsRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
    public Num:number;
}
export class FillingChipsNtf{
    public RoomId:number;
    public UserId:number;
    public Num:number;
}

//跟注
export class CallReq{
    public RoomId:number;
    public UserId:number;
}
export class CallRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
    public Num:number;
}
export class CallNtf{
    public RoomId:number;
    public UserId:number;
    public Num:number;
}

//弃牌
export class GiveUpReq{
    public RoomId:number;
    public UserId:number;
}
export class GiveUpRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
}
export class GiveUpNtf{
    public RoomId:number;
    public UserId:number;
}

//比牌
export class CompetReq{
    public RoomId:number;
    public PromoterId:number; //发起人
    public TargetId:number; //被比牌人
}
export class CompetRsp{
    public Result:number;
    public RoomId:number;
    public PromoterId:number; //发起人
    public TargetId:number; //被比牌人
    public WinnerId:number; //赢家
}
export class CompetNtf{
    public RoomId:number;
    public PromoterId:number; //发起人
    public TargetId:number; //被比牌人
    public WinnerId:number; //赢家
}

//看牌
export class WatchCardReq{
    public RoomId:number;
    public UserId:number;
}
export class WatchCardRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
    public Cards:Array<number>;
}
export class WatchCardNtf{
    public RoomId:number;
    public UserId:number;
}

//发消息
export class SendPlayerMsgReq{
    public RoomId:number;
    public UserId:number;
    public MsgType:number; //消息类型 1.语音 2.固定语音 3.文字 4.表情
    public Msg:string;
}
export class SendPlayerMsgRsp{
    public Result:number;
    public RoomId:number;
    public UserId:number;
    public MsgType:number;
    public Msg:string;
}
export class SendPlayerMsgNtf{
    public MsgType:number;
    public Msg:string;
}

//发牌通知
export class DealCards{
    public SeatId:number;
    public CardNum:number;
}

//翻牌通知
export class ShowCardNtf{
    public RoomId:number;
    public UserId:number;
    public Cards:Array<number>; //牌id
}

//回合开始通知
export class RoundStartNtf{
    public RoomId:number;
    public RoundNum:number;
}

//回合结束通知
export class RoundEndNtf{
    public RoomId:number;
    public RoundNum:number;
}

//回合结算通知
export class RoundSettleNtf{
    public RoomId:number;
    public SettleResult:string;
}

//掉线通知
export class PlayerDisconnectNtf{
    public RoomId:number;
    public UserId:number;
}

//还座位通知
export class ReturnSeatNtf{
    public RoomId:number;
    public SeatId:number;
}

//准备倒计时
export class ReadCounterNtf{
    public RoomId:number;
    public LeftMs:number; //剩下的毫秒数
}

//结算倒计时
export class SettleCounterNtf{
    public RoomId:number;
    public LeftMs:number; //剩下的毫秒数
}

//发牌倒计时
export class DealCounterNtf{
    public RoomId:number;
    public LeftMs:number;
}

//玩家操作倒计时
export class PlayerOpCounterNtf{
    public RoomId:number;
    public LeftMs:number;
}

//房间关闭通知
export class RoomCloseNtf{
    public RoomId:number;
}