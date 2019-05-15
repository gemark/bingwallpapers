package bingwallpapers

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
	var (
		year  string
		month string
		day   string
	)
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	// start := time.Now()
	tso := NewTaskSchedObjct()
	year, month, day = "2019", "05", "01"
	date, err := tso.DatesFormat(year, month, day)
	if err != nil {
		panic(err.Error())
	}
	// end := time.Now().Sub(start)
	// fmt.Println(end)
	fmt.Println(date)
}
