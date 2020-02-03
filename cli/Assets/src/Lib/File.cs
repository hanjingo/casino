using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

public class FileUtil
{
    /// <summary>
    /// 获取当前运行路径
    /// </summary>
    /// <returns></returns>
    public static string GetRunDirectory()
    {
        return Environment.CurrentDirectory;
    }
    /// <summary>
    /// 获取当前运行路径的父路径
    /// </summary>
    /// <returns></returns>
    public static string GetRunDirectoryInParentPath(int PathNum = 3)
    {
        //对于我们当前程序的运行程序来说，当前存储该程序的目录即为他的父对象,
        //获取到当前运行目录的父文件夹信息
        //对于我们的当前程序来说，此目录为Bin
        //GetRunDirectory()的结果为Debug/Release
        DirectoryInfo pathInfo = Directory.GetParent(GetRunDirectory());
        //根据获取路径的层级来用当前路径获取到父路径
        while (PathNum > 0 && pathInfo.Parent != null)
        {
            DirectoryInfo info = pathInfo.Parent;
            pathInfo = info;
            PathNum--;
        }
        //获取到一个完整的文件夹路径
        return pathInfo.FullName;
    }
    /// <summary>
    /// 创建一个文件夹
    /// </summary>
    /// <param name="path">文件夹路径</param>
    /// <param name="foldname">文件夹名称</param>
    /// <returns></returns>
    public static string CreateFolder(string path)
    {
        //如果目录不存在，则创建一个目录文件夹
        if (!Directory.Exists(path))
            Directory.CreateDirectory(path);
        return path;
    }
    /// <summary>
    /// 创建一个文件夹
    /// </summary>
    /// <param name="path">文件夹路径</param>
    /// <param name="foldname">文件夹名称</param>
    /// <returns></returns>
    public static string CreateFolder(string path, string foldname)
    {
        //拼接路径
        string fold = path + "//" + foldname;
        //如果目录不存在，则创建一个目录文件夹
        if (!Directory.Exists(fold))
            Directory.CreateDirectory(fold);
        return fold;
    }
    /// <summary>
    /// 创建一个文本文件
    /// </summary>
    /// <param name="path">文件路径</param>
    /// <param name="filename">文件名称</param>
    /// <param name="info">文件信息</param>
    public static void CreateFile(string path, string filename, string info)
    {
        //写入流对象
        StreamWriter steamWrite;
        //拼接一个文本文件对象
        FileInfo finfo = new FileInfo(path + "//" + filename);
        //判断文件夹是否存在,如果不存在，则创建一个目录
        if (!Directory.Exists(path))
            Directory.CreateDirectory(path);
        //如果该文件已存在，则直接删除
        if (finfo.Exists)
            finfo.Delete();
        //创建一个文本对象，并返回给流对象
        steamWrite = finfo.CreateText();
        //写入数据
        steamWrite.WriteLine(info);
        //写入完成后关闭
        steamWrite.Close();
        //销毁
        steamWrite.Dispose();
    }
    /// <summary>
    /// 像一个文件内添加一段数据
    /// </summary>
    /// <param name="path"></param>
    /// <param name="filename"></param>
    /// <param name="info"></param>
    public static void AddFile(string path, string filename, string info)
    {
        //数据流对象
        StreamWriter steamWrite;
        //拼接一个文本文件对象
        FileInfo finfo = new FileInfo(path + "//" + filename);
        //判断文件夹是否存在,如果不存在，则创建一个目录
        if (!Directory.Exists(path))
            Directory.CreateDirectory(path);
        //如果文件不存在，则创建一个文件，否则加载该文件，并将文件对象赋值给流对象
        steamWrite = !finfo.Exists ? finfo.CreateText() : finfo.AppendText();
        //写入数据
        steamWrite.WriteLine(info);
        //写入完成后关闭
        steamWrite.Close();
        //销毁
        steamWrite.Dispose();
    }
    /// <summary>
    /// 读取文件
    /// </summary>
    /// <param name="path"></param>
    /// <param name="filename"></param>
    /// <returns></returns>
    public static string LoadFile(string path, string filename)
    {
        //如果文件不存在，则返回空
        if (!IsExistsFile(path, filename)) return null;
        //使用读入流对象打开一个文本文件
        StreamReader sr = File.OpenText(path + "//" + filename);
        //创建数组对象
        ArrayList arr = new ArrayList();
        while (true)
        {
            //按行读取文本内容
            string line = sr.ReadLine();
            //如果读取内容到最后一行的下一行，跳出循环
            if (line == null)
                break;
            arr.Add(line);
        }
        string str = "";
        //将读取的内容添加到str字符串中
        foreach (string i in arr)
            str += i;
        //关闭流对象
        sr.Close();
        //销毁流对象
        sr.Dispose();
        //返回读取的内容
        return str;
    }
    /// <summary>
    /// 文件是否存在
    /// </summary>
    /// <param name="path"></param>
    /// <param name="filename"></param>
    /// <returns></returns>
    public static bool IsExistsFile(string path, string filename)
    {
        //如果文件夹不存在的话  直接返回不存在
        if (!Directory.Exists(path))
            return false;
        //如果文件不存在的话   返回不存在
        if (!File.Exists(path + "//" + filename))
            return false;
        return true;
    }
}

