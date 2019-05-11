package bingwallpapers

import (
	"encoding/json"
	"os"
)

// Config 对象
type Config struct {
	LoadedConfig interface{}
}

// LoadConfig 载入配置文件函数
// confPath string 为config配置文件的路径和文件名
// v *interface{} 可以是一个结构体
func (conf *Config) LoadConfig(confPath string, v interface{}) interface{} {
	if conf.LoadedConfig != nil {
		return conf.LoadedConfig
	}
	conf.LoadedConfig = v
	return conf.bingConfigInit(confPath)
}

func (conf *Config) bingConfigInit(confPath string) interface{} {
	var (
		fs  *os.File
		err error
	)
	if fs, err = os.Open(confPath); err != nil {
		panic(err.Error())
	}
	defer fs.Close()

	var filesize int64
	var fileinfo os.FileInfo
	if fileinfo, err = fs.Stat(); err != nil {
		panic(err.Error())
	}
	filesize = fileinfo.Size()
	var data = make([]byte, filesize)
	if _, err := fs.Read(data); err != nil {
		panic(err.Error())
	}
	// conf.LoadedConfig its *interface{}, *struct can work, or other type
	if err := json.Unmarshal(data, conf.LoadedConfig); err != nil {
		panic(err.Error())
	}
	return conf.LoadedConfig
}

// WriteConfig 将结构体转换成json，在写入path路径这种指定文件
func (conf *Config) WriteConfig(path string, v interface{}) error {
	var (
		fs  *os.File
		err error
	)
	fs, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0)
	if err != nil {
		panic(err.Error())
	}
	defer fs.Close()

	jsondata, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	dlen := len(jsondata)
	wlen, err := fs.Write(jsondata)
	if err != nil || wlen == dlen {
		panic(err.Error())
	}
	return nil
}
