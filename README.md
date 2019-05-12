# BingWallpapers
微软必应壁纸每日自动获取

Bing Wallpapers Crewler

---

可爱飞行猪（朋友喜欢叫我🚗奥迪哥or宝马哥😂）-Windows系统小工具 :) 

对于Go语言我一直都有关注，但是最近才开始写Go程序，感觉这个编程语言很棒，为了熟悉Go语言的语法，加上我也很喜欢微软必应的壁纸，就是锁屏界面使用的Win10 spotlight的系列非常漂亮，所以就写了这个小东西。我的系统上执行一次程序大概3~4秒多，主要是获取cn.bing.com的网络壁纸消耗的时间多一些，这个我在下个版本将会改进。这个版本没有用到协程goroutine，下个版本会加入协程功能，希望能将执行效率再提高一些。

我的E-Mail📲: 

<golang83@outlook.com>

<gemarkcg@gmail.com> 🈶🈳💬



## Download MSI 下载安装包☢
安装包下载：(推荐)
[BingWallpapersSetup.exe](https://github.com/gemark/bingwallpapers/releases/download/v0.1.0/BingWallpapersSetup.exe) 

Zip打包下载：
[BingWallpapers_v1.0.zip](https://github.com/gemark/bingwallpapers/releases/download/v0.1.0/BingWallpapers_v1.0.zip)

仅使用的话，建议下载Zip包直接使用程序。

## Download Source and Compile 下载源码并编译☣
### Download 下载：
```shell
git clone git@github.com:gemark/bingwallpapers.git
```
**如果没有安装git和ssh工具**
请直接下载：
```shell
https://github.com/gemark/bingwallpapers/archive/master.zip
```
### Compile 编译：
本程序使用了go module，所以请设置环境变量：

Windows:
> cli> set GO111MODULE=on

Linux:
> cli> export GO111MODULE=on

编译命令与参数：
```shell
go build -ldflags="-H windowsgui -w -s"
```
如果直接go build不带参数，将不会有应用程序图标。

根目录下的resource.syso包含了win32应用程序的icons等资源文件。
使用上面的go build参数，可以自动的将.syso文件编译到.exe文件中。

---

## 关于程序运行的方式💔
使用安装程序安装到最后一步的时候，可以选择马上运行一次。

但本程序并没有干涉用户的启动方式，如：修改/设置 “任务计划程序” 或“开机启动项”。

**关于本程序的运行，必须得用户手动去添加“开机启动”或“任务计划程序”**

最简单的方法是，将程序安装后的快捷方式复制一份到下面这个路径中：
```shell
C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp
```
这样，每次开机到登陆到桌面，就会运行一次，但这样对于长时间开机的用户，可能就会错过Bing必应的每日更新。

**关于“任务计划程序的脚本：”**
创建一个文本文件，然后更改为“ScheduleTask.bat”文件，将下列内容添加到文件中(文件编码不要使用UTF-8请使用ANSI)
BW=xxxxxxxxx 是BingWallpapers.exe的完整路径。

>这里可以下载这个批处理：<https://github.com/gemark/bingwallpapers/blob/master/ScheduleTask.bat>


```shell
@echo off
set BW="D:\BingWallpapers\BingWallpapers.exe"
schtasks /create /tn "必应壁纸自动获取A" /tr %BW% /sc daily /st 14:00:00 /ed 2025/05/02
```
上面的批处理脚本的意思就是每天的14点就是pm2点，准时运行壁纸抓取程序，这个计划持续到2025年的5月2日，如果是英文操作系统，日期可能是dd/mm/yyyy也有可能是mm/dd/yyyy，如果日期格式不符，批处理程序会执行错误后会提示的，修改至对应的格式即可。
如果你想控制的再细致一些，请参考下面的“任务计划程序”流程。

---

**高级用户建议使用“任务计划程序”**

***在桌面使用快捷键“Win + R”打开运行，将下面的内容粘贴进入“运行”中：***
```
%windir%\system32\taskschd.msc /s
```
关于“任务计划程序”的指引：[Windows任务计划程序](https://blog.csdn.net/GeMarK/article/details/90143616)
