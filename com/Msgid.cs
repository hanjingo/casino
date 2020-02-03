using System;
using System.Collections;
using System.Collections.Generic;


public enum MsgId:UInt32
{
    /**********************通用消息************************/
    UnKnownMsg = 0,
    NewConnReq = 1,
    NewConnRsp = 2,
    PlayerOffLineNtfId = 3,
    CheckConnReq = 4,

    //////////////其他///////////////////
    //准备卡住
    FGFReadyWaitingStartNtf    = 23000,
    FGFReadyWaitingCountDoneNtf = 23001 
}
