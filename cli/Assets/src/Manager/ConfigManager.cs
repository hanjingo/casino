using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ConfigManager
{
    private static ConfigManager instance;
    public static ConfigManager Instance
    {
        get
        {
            if (instance == null)
            {
                instance = new ConfigManager();
                instance.Net = new NetConfig();
                instance.Audio = new AudioConfig();
            }
            return instance;
        }
    }

    /// <summary>
    /// 游戏设置
    /// </summary>
    public GameInfo Game = new GameInfo();
    public class GameInfo { 
        //是否开启调试
        public bool DoDebug = false;
    }

    /// <summary>
    /// 网络设置
    /// </summary>
    public NetConfig Net = new NetConfig();
    public class NetConfig { 
        //登录服务器ip
        //public string LoginAddr = "127.0.0.1:10087";
        public string LoginAddr = "47.107.73.146:10087";
        //大厅ip
        public string HallIp;
        //大厅端口
        public int HallPort;
        //新建连接等待(ms)
        public int NewConnWaitTimeOut = 3000;
    }

    /// <summary>
    /// 音乐设置
    /// </summary>
    public AudioConfig Audio;
    public class AudioConfig {
        public bool DoPlayBgm;
        public float BgmVolume;
        public bool DoPlayEff;
        public float EffVolume;
    }

    /// <summary>
    /// 保存到本地
    /// </summary>
    public void Save() { 
        
    }

    /// <summary>
    /// 加载
    /// </summary>
    public void Load(string path) { 
        
    }
}
