package bingwallpapers

import "syscall"

func CreateTaskSchd() {
	tsh, err := syscall.LoadLibrary("taskschd")
	if err != nil {
		panic(err.Error())
	}
	defer syscall.FreeLibrary(tsh)

}
