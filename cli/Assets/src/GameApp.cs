using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameApp 
{
    private static GameApp instance;
    public static GameApp Instance { 
        get { 
            if(instance == null)
            {
                instance = new GameApp();
                instance.init();
            }
            return instance;
        }    
    }

    private bool bInit = false;
    private void init() 
    {
        if (bInit) 
        {
            return;    
        }
        //todo
        bInit = true;
    }

    /// <summary>
    /// 音乐管理器
    /// </summary>
    public AudioManager Audio;

    /// <summary>
    /// 配置管理器
    /// </summary>
    public ConfigManager Config = ConfigManager.Instance;

    /// <summary>
    /// 资源管理器
    /// </summary>
    public CacheManager Cache = new CacheManager();

    /// <summary>
    /// 数据管理器
    /// </summary>
    public DataManager Data = new DataManager();
}
