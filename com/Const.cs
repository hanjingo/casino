using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

/// <summary>
/// 游戏类型
/// </summary>
public enum GameType : UInt32
{ 
    //未知
    UnKnown = 0,
    //登录
    Login = 100,
    //大厅
    Hall = 101,
    //扎金花
    ThreePoker = 102,
    //广东麻将
    GDMahJong = 103,
    //德州扑克
    Holdem = 104,
    //跑胡子
    PaoHuZi = 105,
    //斗牛
    DouNiu = 106,
    //跑得快
    RunFast = 107,
}

/// <summary>
/// 房间游戏状态
/// </summary>
public enum RoomStatus : UInt32
{ 
    //未知状态
    Unknown,
    //房间开始
    Start,
    //房间关闭
    Close,
}

//回调类型
public enum CallType : UInt32
{
    Start = 1900,//开始
    MsgUpdate = 1901, //消息更新
    BoxUpdate = 1902, //消息盒子更新
    PlayerUpdate = 1903, //玩家更新
    ClockUpdate = 1904, //闹钟更新
    BiPaiUpdate = 1905, //比牌更新
    MoveCardUpdate = 1906, //移动牌更新
    PlaceCardUpdate = 1907, //放置牌更新
    ChipUpdate = 1908, //筹码更新
    NextSceneUpdate = 1909, //跳转成精更新
    End = 1910, //结束了
}
public delegate void UpdateCallBack();

/// <summary>
/// 玩家状态
/// </summary>
public enum PlayerStatus : UInt32
{
    WATCH_BATTLE   = 1901, //观战
    SIT_DOWN       = 1902, //坐下
    SIT_AND_READY  = 1903, //坐下+准备
    PLAYING        = 1904, //正在玩
    DIS_CONNECTED  = 1905, //断线
    DEPOSIT        = 1906, //托管
}

/// <summary>
/// 资源路径
/// </summary>
public static class Const
{
    public static string PlayerItemPrefabPath = "prefab/playerItem";
    public static string ClockPrefabPath = "prefab/clock";
    public static string CardPrefabPath = "prefab/card";
    public static string ChipPrefabPath = "prefab/chip";
    public static string MsgBoxPrefabPath = "prefab/msgBox";
    public static string ChipImgPath = "fireGoldFlower/chip/";
    public static string PokerImgPath = "poker/";

    public static Dictionary<string, string> PokerPathMap = new Dictionary<string, string>{
        {PokerName.CLUB_2, PokerImgPath+"CLUB_2" },
        {PokerName.CLUB_3, PokerImgPath+"CLUB_3" },
        {PokerName.CLUB_4, PokerImgPath+"CLUB_4" },
        {PokerName.CLUB_5, PokerImgPath+"CLUB_5" },
        {PokerName.CLUB_6, PokerImgPath+"CLUB_6" },
        {PokerName.CLUB_7, PokerImgPath+"CLUB_7" },
        {PokerName.CLUB_8, PokerImgPath+"CLUB_8" },
        {PokerName.CLUB_9, PokerImgPath+"CLUB_9" },
        {PokerName.CLUB_10, PokerImgPath+"CLUB_10" },
        {PokerName.CLUB_J, PokerImgPath+"CLUB_J" },
        {PokerName.CLUB_Q, PokerImgPath+"CLUB_Q" },
        {PokerName.CLUB_K, PokerImgPath+"CLUB_K" },
        {PokerName.CLUB_A, PokerImgPath+"CLUB_A" },

        {PokerName.DIAMOND_2, PokerImgPath+"DIAMOND_2" },
        {PokerName.DIAMOND_3, PokerImgPath+"DIAMOND_3" },
        {PokerName.DIAMOND_4, PokerImgPath+"DIAMOND_4" },
        {PokerName.DIAMOND_5, PokerImgPath+"DIAMOND_5" },
        {PokerName.DIAMOND_6, PokerImgPath+"DIAMOND_6" },
        {PokerName.DIAMOND_7, PokerImgPath+"DIAMOND_7" },
        {PokerName.DIAMOND_8, PokerImgPath+"DIAMOND_8" },
        {PokerName.DIAMOND_9, PokerImgPath+"DIAMOND_9" },
        {PokerName.DIAMOND_10, PokerImgPath+"DIAMOND_10" },
        {PokerName.DIAMOND_J, PokerImgPath+"DIAMOND_J" },
        {PokerName.DIAMOND_Q, PokerImgPath+"DIAMOND_Q" },
        {PokerName.DIAMOND_K, PokerImgPath+"DIAMOND_K" },
        {PokerName.DIAMOND_A, PokerImgPath+"DIAMOND_A" },

        {PokerName.HEART_2, PokerImgPath+"HEART_2" },
        {PokerName.HEART_3, PokerImgPath+"HEART_3" },
        {PokerName.HEART_4, PokerImgPath+"HEART_4" },
        {PokerName.HEART_5, PokerImgPath+"HEART_5" },
        {PokerName.HEART_6, PokerImgPath+"HEART_6" },
        {PokerName.HEART_7, PokerImgPath+"HEART_7" },
        {PokerName.HEART_8, PokerImgPath+"HEART_8" },
        {PokerName.HEART_9, PokerImgPath+"HEART_9" },
        {PokerName.HEART_10, PokerImgPath+"HEART_10" },
        {PokerName.HEART_J, PokerImgPath+"HEART_J" },
        {PokerName.HEART_Q, PokerImgPath+"HEART_Q" },
        {PokerName.HEART_K, PokerImgPath+"HEART_K" },
        {PokerName.HEART_A, PokerImgPath+"HEART_A" },

        {PokerName.SPADE_2, PokerImgPath+"SPADE_2" },
        {PokerName.SPADE_3, PokerImgPath+"SPADE_3" },
        {PokerName.SPADE_4, PokerImgPath+"SPADE_4" },
        {PokerName.SPADE_5, PokerImgPath+"SPADE_5" },
        {PokerName.SPADE_6, PokerImgPath+"SPADE_6" },
        {PokerName.SPADE_7, PokerImgPath+"SPADE_7" },
        {PokerName.SPADE_8, PokerImgPath+"SPADE_8" },
        {PokerName.SPADE_9, PokerImgPath+"SPADE_9" },
        {PokerName.SPADE_10, PokerImgPath+"SPADE_10" },
        {PokerName.SPADE_J, PokerImgPath+"SPADE_J" },
        {PokerName.SPADE_Q, PokerImgPath+"SPADE_Q" },
        {PokerName.SPADE_K, PokerImgPath+"SPADE_K" },
        {PokerName.SPADE_A, PokerImgPath+"SPADE_A" },

        {PokerName.SMALL_JOKER, PokerImgPath+"SMALL_JOKER" },
        {PokerName.BIG_JOKER, PokerImgPath+"BIG_JOKER" },
    };
    
}