using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CacheManager : MonoBehaviour
{
    /// <summary>
    /// 图片缓存
    /// </summary>
    Dictionary<string, Sprite> Sprites = new Dictionary<string, Sprite>();

    /// <summary>
    /// 音乐缓存
    /// </summary>
    Dictionary<string, AudioClip> Clips = new Dictionary<string, AudioClip>();

    /// <summary>
    /// 预制体缓存
    /// </summary>
    public Dictionary<string, GameObject> GameObjs = new Dictionary<string, GameObject>();

    /// <summary>
    /// 文本缓存
    /// </summary>
    public Dictionary<string, TextAsset> Texts = new Dictionary<string, TextAsset>();


    private void Awake()
    {
        GameApp.Instance.Cache = this;
    }

    /// <summary>
    /// 加载图片
    /// </summary>
    /// <param name="path"></param>
    /// <returns></returns>
    public Sprite GetSprite(string spriteKey)
    { 
        Sprite sprite;
        if(Sprites.TryGetValue(spriteKey, out sprite)) { 
            return sprite;    
        }
        sprite = Resources.Load(spriteKey, typeof(Sprite)) as Sprite;
        if(sprite != null) {
            Sprites[spriteKey] = sprite; 
            return sprite;
        }
        return null;
    }

    /// <summary>
    /// 设置图片
    /// </summary>
    /// <param name="spriteKey"></param>
    /// <param name="data"></param>
    public void SetSprite(string spriteKey, Sprite img)
    { 
        
    }

    /// <summary>
    /// 加载音频
    /// </summary>
    /// <param name="path"></param>
    /// <returns></returns>
    public AudioClip GetClip(string key)
    { 
        AudioClip clip;
        if(Clips.TryGetValue(key, out clip)) { 
            return clip;    
        }
        clip = Resources.Load(key, typeof(AudioClip)) as AudioClip;
        if(clip != null) { 
            Clips[key] = clip;
            return clip;
        }
        return null;
    }

    /// <summary>
    /// 加载文本
    /// </summary>
    /// <param name="path"></param>
    /// <returns></returns>
    public string GetText(string key)
    { 
        TextAsset txt;
        if(Texts.TryGetValue(key, out txt)) { 
            return txt.text;    
        }
        txt = Resources.Load(key, typeof(TextAsset)) as TextAsset;
        if(txt == null)
            return string.Empty;
        Texts[key] = txt;
        return txt.text;
    }

    /// <summary>
    /// 加载GameObject
    /// </summary>
    /// <param name="path"></param>
    /// <returns></returns>
    public GameObject GetGameObj(string key)
    { 
        GameObject obj;
        if(GameObjs.TryGetValue(key, out obj)) { 
            return obj;    
        }
        obj = (GameObject)Resources.Load(key);
        if(obj == null)
            return null;
        GameObjs[key] = obj;
        return obj;
    }
}
