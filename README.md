# BingWallpapers
微软必应壁纸每日自动获取
Microsoft Bing Wallpapers Crewler
可爱飞行猪（朋友喜欢叫我🚗奥迪哥or宝马哥😂）-Windows系统小工具 :) 

对于Go语言我一直都有关注，但是最近才开始写Go程序，感觉这个编程语言很棒，为了熟悉Go语言的语法，加上我也很喜欢微软必应的壁纸，就是锁屏界面使用的Win10 spotlight的系列非常漂亮，所以就写了这个小东西。我的系统上执行一次程序大概3~4秒多，主要是获取cn.bing.com的网络壁纸消耗的时间多一些，这个我在下个版本将会改进。这个版本没有用到协程goroutine，下个版本会加入协程功能，希望能将执行效率再提高一些。

我的E-Mail📲: <golang83@outlook.com>、<gemarkcg@gmail.com> 🈶🈳💬

## Download MSI 下载安装包☢
安装包下载：
[BingWallpapersSetup.exe](https://github.com/gemark/bingwallpapers/releases/download/v0.1.0/BingWallpapersSetup.exe) 仅使用的话，建议下载安装程序。

Zip打包下载：
[BingWallpapers_v1.0.zip](https://github.com/gemark/bingwallpapers/releases/download/v0.1.0/BingWallpapers_v1.0.zip)

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
```
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
```
C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp
```
这样，每次开机到登陆到桌面，就会运行一次，但这样对于长时间开机的用户，可能就会错过Bing必应的每日更新。

**高级用户建议使用“任务计划程序”**
***在桌面使用快捷键“Win + R”打开运行，将下面的内容粘贴进入“运行”中：***
```
%windir%\system32\taskschd.msc /s
```
![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc00.png)

## 下面是“任务计划程序”的创建/修改步骤：
![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc01.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc02.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc03.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc04.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc05.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc06.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc07.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc08.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc09.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc10.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc11.png)

![avatar](https://github.com/gemark/bingwallpapers/blob/DocScreenShot/doc_screenshot/doc12.png)