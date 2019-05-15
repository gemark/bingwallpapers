//	Program Name:	BingWallpapers
//	Author:			https://github.com/gemark
//	E-Mail:			golang83@outlook.com
//	First Date:		2019/05/09 14:27
//	Last Date:		2019/05/11 16:11
//	Description:	获取win10系统中，Bing必应开机的每日壁纸
//					获取cn.bing.com Image of day 系列壁纸

package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	bingTools "github.com/gemark/bingwallpapers/tools"
)

// BingCrewlerConfig 配置
type BingCrewlerConfig struct {
	Local       bingWallpapers
	Web         bingWallpapers
	WinTaskName string
	StartTime   string
}

// BingCrewlerConfig 字段
type bingWallpapers struct {
	SavedRootPath string
	LocalPath     string
	FilesHash     string
	PowerOn       bool
}

// bing image of day json format
// 本来可以直接在读入的[]byte里做个简单的解析
// 但考虑到后面不知道还需不需要扩展本程序，所以
// 就将json的结构留在这里。
type bingjson struct {
	Images []struct {
		Startdate     string `json:"startdate"`
		Fullstartdate string `json:"fullstartdate"`
		Enddate       string `json:"enddate"`
		URL           string `json:"url"`
		Urlbase       string `json:"urlbase"`
		Copyright     string `json:"copyright"`
		Copyrightlink string `json:"copyrightlink"`
		Title         string `json:"title"`
		Quiz          string `json:"quiz"`
		Wp            bool   `json:"wp"`
		Hsh           string `json:"hsh"`
		Drk           int32  `json:"drk"`
		Top           int32  `json:"top"`
		Bot           int32  `json:"bot"`
		Hs            []byte `json:"hs"`
	} `json:"images"`
	Tooltip struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltip"`
}

var (
	// %USERPROFILE%\\AppData\\Local\\Packages...
	// \\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\
	// 本地的 win10 spotlight 路径，但好像并不会每天都会更新存储
	localPath  string
	conf       *bingTools.Config  // 配置文件对象 	config file object
	slog       *bingTools.Logger  // 日志对象		logger object
	bingConfig *BingCrewlerConfig // 配置结构		config struct
)

func init() {
	// 初始化配置信息对象 init config struct object
	bingConfig = new(BingCrewlerConfig)
	// 初始化配置工具对象 init config file object
	conf = new(bingTools.Config)
	// 创建日志对象	create logger object
	slog = slog.CreateLogger("./log", "info.log")
	// 载入配置文件(这里已经拿到全部配置了) get all config setting from file
	conf.LoadConfig("./data/config.json", bingConfig)
}

// errorHandling 错误处理处
func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		// 处理异常 all panic in there
		if err := recover(); err != nil {
			// 错误信息写入日志
			slog.ErrorOutput(err)
			slog.StdOut("程序运行错误，请检测日志文件 “./info.log”")
			slog.CloseLogger()
			os.Exit(1)
		}
		// 关闭日志对象 close logger object
		slog.CloseLogger()
	}()
	slog.InfoOutput("程序开始执行...")
	start := time.Now()
	// 检测配置文件中路径的合法性 check config file
	checkConfigObject()

	execPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	execPath = execPath + "\\" + "BingWallpapers.exe"
	if err != nil {
		panic(err)
	}

	tso := bingTools.NewTaskSchedObjct()
	_, stderr, err := tso.QueryTask(bingConfig.WinTaskName)
	if err != nil {
		panic(err.Error())
	}
	if stderr != "" {
		stdout2, stderr2, err := tso.CreateTask(bingConfig.WinTaskName, execPath, bingConfig.StartTime)
		if err != nil {
			slog.ErrorOutput("CreateTaskError->" + err.Error())
		}
		if stderr2 != "" {
			slog.ErrorOutput("CreateTask->" + stderr2)
		}
		if stdout2 != "" {
			slog.InfoOutput("CreateTask->" + stdout2)
		}
	}

	// 获取壁纸 get local and web wallpapers in microsoft bing
	GetWallpapers()
	end := time.Now().Sub(start)
	useTime := fmt.Sprintf("程序执行完毕，耗时：%v", end.Seconds())
	slog.InfoOutput(useTime)
}

// GetWallpapers 获取必应壁纸
func GetWallpapers() {
	// 获取本地壁纸
	// if config Local & Web -> PowerOn setting true
	if bingConfig.Local.PowerOn {
		getLocal()
	}
	if bingConfig.Web.PowerOn {
		getWeb()
	}
}

func getWeb() {
	url := bingConfig.Web.LocalPath
	urls := strings.Split(url, "${n}")
	dict := make(map[string]string)
	for i := 0; i < 30; i++ {
		bingURL := fmt.Sprintf("%s%d%s", urls[0], i, urls[1])
		resp, err := http.Get(bingURL)
		if err != nil {
			slog.ErrorOutput(err.Error())
			continue
		}
		if resp.StatusCode == 200 {
			br, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err.Error())
			}
			bj := new(bingjson)
			json.Unmarshal(br, bj)
			_, isExists := dict[bj.Images[0].URL]
			if !isExists {
				dict[bj.Images[0].URL] = "https://cn.bing.com" + bj.Images[0].URL
				jpegResp, err := http.Get(dict[bj.Images[0].URL])
				if err != nil {
					slog.InfoOutput(err.Error())
					continue
				}
				if jpegResp.StatusCode == 200 {
					jpegBufRead, err := ioutil.ReadAll(jpegResp.Body)
					if err != nil {
						panic(err.Error())
					}
					hashdata := sha1.Sum(jpegBufRead)
					pstr := byteToString([]byte(hashdata[:]))
					filePath := bingConfig.Web.SavedRootPath + *pstr + ".jpg"
					webDict := readHash(bingConfig.Web.FilesHash)
					_, isWebFileKeyExists := (*webDict)[*pstr]
					if !isWebFileKeyExists {
						fs, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0)
						if err != nil {
							panic(err.Error())
						}
						_, err = fs.Write(jpegBufRead)
						if err != nil {
							fs.Close()
							panic(err.Error())
						}
						wirteHash(*pstr, bingConfig.Web.FilesHash)
						fs.Close()
					}
				}
			} else {
				break
			}
		}
		resp.Body.Close()
	}
}

func getLocal() {
	var fs *os.File
	var userHome string
	var err error
	defer func() {
		if fs != nil {
			fs.Close()
		}
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	// 获取系统用户home目录 get winsys user home dir
	if userHome, err = os.UserHomeDir(); err != nil {
		panic(err.Error())
	}
	localPath = userHome + bingConfig.Local.LocalPath

	// 打开系统spotlight壁纸所在路径 open path
	if fs, err = os.Open(localPath); err != nil {
		panic(err.Error())
	}

	// get path FileInfo
	if _, err = fs.Stat(); err != nil {
		panic(err.Error())
	} else {
		// all fileinfo in path
		var content []os.FileInfo
		if content, err = fs.Readdir(0); err != nil {
			panic(err.Error())
		}
		for _, v := range content {
			if bingTools.IsNeedImages(v, localPath) {
				SaveToDisk(v, localPath)
			}
		}
	}
}

// SaveToDisk 将壁纸保存到磁盘
func SaveToDisk(v os.FileInfo, localPath string) {
	var (
		fs  *os.File
		err error
	)
	FilePath := localPath + "\\" + v.Name()
	if fs, err = os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0); err != nil {
		panic(err.Error())
	}
	defer fs.Close()
	hash := sha1.New()
	if _, err = io.Copy(hash, fs); err != nil {
		panic(err.Error())
	}
	hashdata := hash.Sum(nil)
	pstr := byteToString(hashdata)
	dict := readHash(bingConfig.Local.FilesHash)
	_, isExists := (*dict)[*pstr]
	if !isExists {
		// 不存在，写入磁盘
		des := bingConfig.Local.SavedRootPath + *pstr + ".jpg"
		saveToPath(FilePath, des)
		// sha1 sum 写入文件
		wirteHash(*pstr, bingConfig.Local.FilesHash)
	}
}

func wirteHash(h string, path string) {
	var (
		fs  *os.File
		err error
	)
	if fs, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0); err != nil {
		panic(err.Error())
	}
	defer fs.Close()

	if _, err = fs.WriteString(h + "\n"); err != nil {
		panic(err.Error())
	}
}

func readHash(filepath string) *map[string]bool {
	var (
		fs  *os.File
		br  *bufio.Reader
		err error
	)
	if fs, err = os.OpenFile(filepath, os.O_RDONLY, 0); err != nil {
		panic(err.Error())
	}
	defer fs.Close()

	// 通过buffer io来读取文件
	// use bufio read file lines.
	br = bufio.NewReader(fs)
	dict := make(map[string]bool)
	for {
		line, _, eof := br.ReadLine()
		if eof == io.EOF {
			break
		}
		linestr := string(line[:])
		dict[linestr] = true
	}
	return &dict
}

// saveToPath 将文件写入目标目录
// src 含文件名
// des 不含文件名，仅为目标目录
func saveToPath(src string, des string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err.Error())
	}
	if err := ioutil.WriteFile(des, data, 0); err != nil {
		panic(err.Error())
	}
}

func byteToString(b []byte) *string {
	var fmtstr = ""
	for _, v := range b {
		if v < 0x10 {
			fmtstr += fmt.Sprintf("0%x", v)
		} else {
			fmtstr += fmt.Sprintf("%x", v)
		}
	}
	return &fmtstr
}

// 主要是根据配置看看能否将配置中的路径到目录创建，如果配置的路径合法，应该Pass
func checkConfigObject() {
	// 检测本地壁纸配置
	Local := []string{
		bingConfig.Local.SavedRootPath,
		bingConfig.Local.LocalPath,
		bingConfig.Local.FilesHash,
		bingConfig.Web.SavedRootPath,
		bingConfig.Web.LocalPath,
		bingConfig.Web.FilesHash,
	}
	for i, v := range Local {
		// 处理路径中的分隔符
		v = filepath.ToSlash(v)
		t := (reflect.TypeOf(v)).Name()
		if t != "string" || v == "" {
			panic("配置文件Local项，路径配置错误！")
		}
		if i != 1 && i != 4 {
			// 检测路径是否存在
			is, e := dirIsExists(v)
			if e != nil {
				panic(e.Error())
			}
			if !is {
				dirs := filepath.Dir(v)
				if e := os.MkdirAll(dirs, 0); e != nil {
					slog.ErrorOutput(e.Error())
				}
			}
			if i == 2 || i == 5 {
				var (
					fs  *os.File
					err error
				)
				if fs, err = os.OpenFile(v, os.O_CREATE, 0); err != nil {
					panic(err.Error())
				}
				fs.Close()
			}
		}
	}
}

// dirIsExists 检测路径（目录）是否存在
func dirIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
