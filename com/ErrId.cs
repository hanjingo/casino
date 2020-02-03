using System;
using System.Collections;
using System.Collections.Generic;

public static class ErrId { 
    private static Dictionary<UInt32, Error> m = new Dictionary<uint, Error>();
    public static void Add(UInt32 id, string desc) { 
        Error err = new Error(id, desc);
        m.Add(id, err);
    }
    public static string GetDesc(UInt32 id) { 
        if(!m.ContainsKey(id))    
            return "未知错误";
        return m[id].Description;
    }
    public static UInt32 OK = 1;
    static ErrId() { 
        Add(0, "未知错误");
        Add(1, "成功");
        Add(2, "失败");
        Add(3, "参数错误");
        Add(5, "不是当前用户");

        Add(100,"房间错误");
        Add(101,"房间无效");
        Add(102, "玩家不在房间中"); 
        Add(103, "没座位了");
        Add(104, "座位已存在");
        Add(105, "玩家没有座位");
        Add(106, "座位不存在");
        Add(107, "玩家已经坐下");
        Add(108, "房间不存在");
        Add(109, "回合尚未开始");
        Add(110, "对方不在房间");
        Add(111, "最后一局");
        Add(112, "庄家不存在");
        Add(113, "玩家无操作的权限");
        Add(114, "房间不在状态");
        Add(115, "没得面可以下");
        Add(116, "房间掉线列表为nil");
        Add(117, "创建房间失败");
        Add(118, "拿房间地址失败");
        Add(119, "房间状态错误");
        Add(120, "房卡数量不对");
        Add(121, "退还房卡失败");
        Add(122, "房间位置不可用");

        Add(1000, "数据结构错误");
        Add(1001, "管道已满");
        Add(1002, "管道已空");
        Add(1003, "队列满");
        Add(1004, "队列空");

        Add(2000, "数据库错误");
        Add(2001, "redis连接无效");
        Add(2002, "从redis读取数据失败");
        Add(2003, "将数据保存到redis失败");

        Add(3000, "序列化错误");
        Add(3001, "json格式化错误");
        Add(3002, "json反格式化错误"); 
        Add(3003, "proto格式化错误");
        Add(3004, "proto反格式化错误");
        Add(3005, "消息过短");
        Add(3006, "消息过长");
        Add(3007, "content为空");
        Add(3008, "消息不完整");
        Add(3009, "消息为空");

        Add(4000, "玩家错误");
        Add(4001, "玩家不存在");
        Add(4002, "玩家已登录");
        Add(4003, "玩家筹码不够");
        Add(4004, "玩家没坐下");
        Add(4005, "未知文字消息");
        Add(4006, "用户在黑名单中");
        Add(4007, "对方没坐下");
        Add(4008, "没有玩家准备");
        Add(4009, "比牌失败");
        Add(4010, "玩家不在状态");
        Add(4011, "庄家没有资格下注");
        Add(4012, "玩家信息保存失败");
        Add(4013, "玩家已存在");
        Add(4014, "注册玩家信息失败");
        Add(4015, "玩家token错误");
        Add(4016, "玩家未登录");
        Add(4017, "玩家不在大厅");
        Add(4018, "玩家已经在房间里了");
        Add(4019, "玩家已经被淘汰");

        Add(5000, "step错误");
        Add(5001, "prev错误");

        Add(6000, "pub sub错误");
        Add(6001, "已经订阅");

        Add(7000, "扑克错误");
        Add(7001, "扑克牌池已空");
        Add(7002, "此牌不存在");

        //http相关
        Add(8000, "http错误");
        Add(8001, "http解析资源失败");
        Add(8002, "资源不存在");

        //服务器相关
        Add(9000, "服务器错误");
        Add(9001, "服务器密码错误");
        Add(9002, "服务器信息错误");

        //筹码相关
        Add(10000, "筹码错误");
        Add(10001, "筹码过小");
        Add(10002, "筹码数量错误");

        //扎金花相关
        Add(11000, "扎金花错误");
        Add(11001, "扎金花比牌失败");
        Add(11002, "扎金花第一轮不能比牌");
    }
}

public class Error { 
    public UInt32 Id;
    public string Description;
    public Error(UInt32 id, string desc) { 
        Id = id;
        Description = desc;
    }
}