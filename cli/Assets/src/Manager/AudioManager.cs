using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class AudioManager : MonoBehaviour
{
    /// <summary>
    /// bgm
    /// </summary>
    AudioSource bgmSource;
    AudioSource BgmSource { 
        get { 
            if(bgmSource == null)
            { 
                if(gameObject.GetComponent<AudioSource>())    
                    bgmSource = gameObject.GetComponent<AudioSource>();
                else
                    bgmSource = gameObject.AddComponent<AudioSource>();
            }
            return bgmSource;
        }
    }

    /// <summary>
    /// eff
    /// </summary>
    List<AudioSource> EffSource = new List<AudioSource>();

    /// <summary>
    /// bgm开关
    /// </summary>
    public bool DoPlayBgm = true;

    /// <summary>
    /// bgm音量
    /// </summary>
    public float BgmVolume = 0.5f;

    /// <summary>
    /// eff开关
    /// </summary>
    public bool DoPlayEff = true;

    /// <summary>
    /// eff音量
    /// </summary>
    public float EffVolume = 0.5f;

    private void Awake()
    {
        this.DoPlayBgm = GameApp.Instance.Config.Audio.DoPlayBgm;
        this.BgmVolume = GameApp.Instance.Config.Audio.BgmVolume;
        this.DoPlayEff = GameApp.Instance.Config.Audio.DoPlayEff;
        this.EffVolume = GameApp.Instance.Config.Audio.EffVolume;
        GameApp.Instance.Audio = this;
    }

    /// <summary>
    /// 保存设置
    /// </summary>
    public void SaveConfig()
    { 
        GameApp.Instance.Config.Audio.DoPlayBgm = this.DoPlayBgm;
        GameApp.Instance.Config.Audio.BgmVolume = this.BgmVolume;
        GameApp.Instance.Config.Audio.DoPlayEff = this.DoPlayEff;
        GameApp.Instance.Config.Audio.EffVolume = this.EffVolume;
        GameApp.Instance.Config.Save();
    }

    /// <summary>
    /// 获得音频
    /// </summary>
    /// <param name="key"></param>
    /// <returns></returns>
    public AudioClip GetClip(string key)
    { 
        AudioClip clip = GameApp.Instance.Cache.GetClip(key);
        return clip;
    }

    /// <summary>
    /// 播放bgm
    /// </summary>
    /// <param name="key"></param>
    public void PlayBgm(string key)
    { 
        if(!DoPlayBgm)
            return;
        AudioClip clip = GetClip(key);
        if(clip == null) { 
            return;    
        }
        BgmSource.clip = clip; //bgm只可能有一个
        BgmSource.loop = true;
        BgmSource.volume = this.BgmVolume;
        BgmSource.Play();
    }

    /// <summary>
    /// 停掉bgm
    /// </summary>
    public void StopBgm()
    { 
        BgmSource.Stop();    
    }

    /// <summary>
    /// 播放音效
    /// </summary>
    /// <param name="key"></param>
    public void PlayEff(string key)
    { 
        if(!DoPlayEff)    
            return;
        AudioClip clip = GetClip(key);
        if(clip == null)
            return;
        AudioSource source = GetEffSource();
        source.clip = clip;
        source.loop = false;
        source.Play();
    }

    /// <summary>
    /// 关闭音效
    /// </summary>
    /// <param name="key"></param>
    public void CloseEff(string key)
    { 
        for(int i = 0; i < EffSource.Count; i++)
        { 
            if(EffSource[i].clip.name == key) {
                EffSource[i].Stop();
                return;
            }   
        }
    }

    /// <summary>
    /// 获得一个音效播放器
    /// </summary>
    /// <returns></returns>
    public AudioSource GetEffSource() { 
        for(int i = 0; i < EffSource.Count; i++)
        { 
            if(!EffSource[i].isPlaying)
            { 
                return EffSource[i]    ;
            }
        }
        AudioSource source = gameObject.AddComponent<AudioSource>();
        source.volume = GameApp.Instance.Config.Audio.EffVolume;
        source.loop = false;
        EffSource.Add(source);
        return source;
    }
}
