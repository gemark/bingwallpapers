/*
   _____       __   __             _  __ 
  ╱ ____|     |  ╲/   |           | |/ / 
 | |  __  ___ |  ╲ /  | __  _ _ __| ' /  
 | | |_ |/ _ ╲| |╲ /| |/ _`  | '__|  <   
 | |__| |  __/| |   | (  _|  | |  | . ╲  
  ╲_____|╲___ |_|   |_|╲__,_ |_|  |_|╲_╲ 
 可爱飞行猪❤: golang83@outlook.com  💯💯💯
 Author Name: GeMarK.VK.Chow奥迪哥  🚗🔞🈲
 Creaet Time: 2019/05/09 - 09:32:25
 ProgramFile: images.go
 Description:
 用于必应壁纸程序的配置输入输出
*/

package tools

import (
	"image"
	"image/jpeg"
	"os"
)

// IsNeedImages 检测是否是需要的jpeg文件
func IsNeedImages(fi os.FileInfo, localPath string) bool {
	var fs *os.File
	var err error
	// errors processing
	defer func() {
		if err := recover(); err != nil {
			if fs != nil {
				fs.Close()
			}
			panic(err)
		}
	}()

	// check its jpg/jpeg format with file
	// byte offset:
	// 2bytes -> hex 0xffd8 -> SOI marker		1 ~ 2 bytes
	// 2bytes -> hex 0xffe0 -> APP0 marker		3 ~ 4 bytes AAP0 marker start
	// -------note: APP1 -> hex 0xffe0 maybe 0xffe1(tiff)http://zh.wikipedia.org/wiki/TIFF can be ignored
	// 2bytes -> data segment length, include its self, but exclude marker code
	// 第5和第6字节描述数据的的长度，在这个从度从第5字节开始，通常是16个字节
	// 如果是0x0010的话，就是16字节，到第17，18的时候如果是0xffdb就到DQT or DHT标记了。
	// 2bytes -> hex 0xffdb -> DQT marker

	// why not use the image package directly?
	// its more efficient by extracting partial byte recognition types
	// if file exclude JFIF header, skip it, but If the JFIF header is included, the image package is called.

	// open file in wallpager dir
	file := localPath + "\\" + fi.Name()
	if fs, err = os.Open(file); err != nil {
		panic(err.Error())
	}

	// first 2 bytes
	jpegSOI := make([]byte, 2)
	if len, err := fs.Read(jpegSOI); (err != nil && len != 2) && (jpegSOI[0] != 0xff || jpegSOI[1] != 0xd8) {
		return false
	}

	// 3 ~ 4 bytes
	jpegAPP0 := make([]byte, 2)
	if len, err := fs.Read(jpegAPP0); (err != nil && len != 2) && (jpegAPP0[0] != 0xff || jpegAPP0[1] != 0xe0) {
		return false
	}
	fs.Seek(0, 0)
	var pic image.Image
	if pic, err = jpeg.Decode(fs); err != nil {
		return false
	}
	var rect image.Rectangle
	rect = pic.Bounds()
	if rect.Empty() {
		return false
	}

	width, height := rect.Dx(), rect.Dy()
	if width < height || width < 1920 {
		return false
	}

	return true
}
