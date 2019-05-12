# <p style='color:#375EAB'>BingWallpapers</p>
微软必应壁纸每日自动获取
Microsoft Bing Wallpapers Crewler

(奥迪哥)可爱飞行猪-Windows系统小工具 :) 

Go语言一直都有关注，但是最近才开始写Go程序，感觉很这个语言很棒，为了熟悉Go语言，所以写点小工具。我也很喜欢微软必应的壁纸，然后就是锁屏界面windows spotlight的系列都很漂亮，所以就写了这个小东西。我的系统上执行一次程序大概3~4秒多，主要是获取cn.bing.com的网络壁纸消耗的时间多一些，这个我在下个版本将会改进。这个版本没有用到协程goroutine，下一个将会改一下。希望能将执行效率再提高一些。

我的E-Mail: <golang83@outlook.com>、<gemarkcg@gmail.com>

## Download MSI 下载安装包
## Download Source and Compile 下载源码并编译
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
