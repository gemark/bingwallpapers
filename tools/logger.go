/*
   _____       __   __             _  __ 
  ╱ ____|     |  ╲/   |           | |/ / 
 | |  __  ___ |  ╲ /  | __  _ _ __| ' /  
 | | |_ |/ _ ╲| |╲ /| |/ _`  | '__|  <   
 | |__| |  __/| |   | (  _|  | |  | . ╲  
  ╲_____|╲___ |_|   |_|╲__,_ |_|  |_|╲_╲ 
 可爱飞行猪❤: golang83@outlook.com  💯💯💯
 Author Name: GeMarK.VK.Chow奥迪哥  🚗🔞🈲
 Creaet Time: 2019/05/09 - 14:27:15
 ProgramFile: logger.go
 Description:
 用于必应壁纸程序的日志输出
*/

package tools

import (
	"fmt"
	"log"
	"os"
)

// Logger 对象
type Logger struct {
	Format   string
	LogPath  string
	Filename string
	logobj   *log.Logger
	logFS    *os.File
}

// CreateLogger 创建日志对象
func (logger *Logger) CreateLogger(logpath string, filename string) *Logger {
	logger = new(Logger)
	logger.Format = ""
	logger.LogPath = logpath
	logger.Filename = filename
	logger.logobj = new(log.Logger)
	logger.openLogFile(logpath, filename)
	logger.logobj.SetOutput(logger.logFS)
	logger.logobj.SetFlags(log.LstdFlags)
	return logger
}

// ErrorOutput 错误输出到文件
// msg 为字符串类型的错误信息
func (logger *Logger) ErrorOutput(msg interface{}) {
	logger.logobj.SetPrefix("[ErrorMSG]")
	switch msg.(type) {
	case string:
		logger.logobj.Printf("%s\r\n", msg)
	case error:
		var err = msg.(error)
		logger.logobj.Printf("%s\r\n", err.Error())
	}
}

// StdOut 使用fmt.Println输入字符串
func (logger *Logger) StdOut(msg string) {
	fmt.Printf("%s\r\n", msg)
}

// InfoOutput 输出普通信息到文件
// msg 为字符串类型的文本信息
func (logger *Logger) InfoOutput(msg interface{}) {
	logger.logobj.SetPrefix("[InfoMSG]")
	switch msg.(type) {
	case string:
		logger.logobj.Printf("%s\r\n", msg)
	case error:
		var err = msg.(error)
		logger.logobj.Printf("%s\r\n", err.Error())
	}
}

func (logger *Logger) openLogFile(logpath string, filename string) {
	var err error
	path := logpath + "/" + filename
	if logger.logFS, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0); err != nil {
		panic(err.Error())
	}
}

// CloseLogger 关闭日志对象
func (logger *Logger) CloseLogger() {
	// 关闭文件对象
	logger.logFS.Close()
}
