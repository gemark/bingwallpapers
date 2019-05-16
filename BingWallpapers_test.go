package main

import (
	"fmt"
	"os"
	"testing"
	
	TaskSched "github.com/gemark/bingwallpapers/tools"
)

// Test_run 测试测试
func TestTaskSched(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	tso := TaskSched.NewTaskSchedObjct()
	stdout, stderr, err := tso.QueryTask("bingwallpapers")
	if err != nil {
		fmt.Println(err.Error())
	}
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
}
