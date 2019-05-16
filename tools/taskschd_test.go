/*
   _____       __   __             _  __ 
  ╱ ____|     |  ╲/   |           | |/ / 
 | |  __  ___ |  ╲ /  | __  _ _ __| ' /  
 | | |_ |/ _ ╲| |╲ /| |/ _`  | '__|  <   
 | |__| |  __/| |   | (  _|  | |  | . ╲  
  ╲_____|╲___ |_|   |_|╲__,_ |_|  |_|╲_╲ 
 可爱飞行猪❤: golang83@outlook.com  💯💯💯
 Author Name: GeMarK.VK.Chow奥迪哥  🚗🔞🈲
 Creaet Time: 2019/05/13 - 13:44:56
 ProgramFile: taskschd_test.go
 Description:
 这个程序主要是为bingwallpapers程序
 创建windows系统中的任务计划程序项目
*/

package tools

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestTaskSchdCreate(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tso := NewTaskSchedObjct()
	stdout, stderr, e := tso.CreateTask("bingwallpapers", "D:\\BingWallpapers\\BingWallpapers.exe", "20:00:00")
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(stderr)
	fmt.Println(stdout)
}

func TestTaskSchdQuery(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	tso := NewTaskSchedObjct()
	stdout, _, err := tso.QueryTask("bingwallpapers")
	if err != nil {
		fmt.Println(err.Error())
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
}

func Test_DateFormat(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
}
