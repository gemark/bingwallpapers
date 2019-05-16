/*
   _____       __   __             _  __ 
  ╱ ____|     |  ╲/   |           | |/ / 
 | |  __  ___ |  ╲ /  | __  _ _ __| ' /  
 | | |_ |/ _ ╲| |╲ /| |/ _`  | '__|  <   
 | |__| |  __/| |   | (  _|  | |  | . ╲  
  ╲_____|╲___ |_|   |_|╲__,_ |_|  |_|╲_╲ 
 可爱飞行猪❤: golang83@outlook.com  💯💯💯
 Author Name: GeMarK.VK.Chow奥迪哥  🚗🔞🈲
 Creaet Time: 2019/05/13 - 13:43:01
 ProgramFile: taskschd.go
 Description:
 这个程序主要是为bingwallpapers程序
 创建windows系统中的任务计划程序项目
*/

package tools

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	registry "golang.org/x/sys/windows/registry"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// TaskSchedObject object
type TaskSchedObject struct {
	DateFormat string
	TimeFormat string
}

// NewTaskSchedObjct 创建TaskSchedObject
func NewTaskSchedObjct() *TaskSchedObject {
	obj := new(TaskSchedObject)
	obj.getDateTimeFormat()
	return obj
}

// GetDateTimeFormat 获取系统的日期/时间格式
func (ts *TaskSchedObject) getDateTimeFormat() {
	// `reg query "HKEY_CURRENT_USER\Control Panel\International" /v sLongDate`
	k, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\International`, registry.ALL_ACCESS)
	if err != nil {
		panic(err.Error())
	}
	defer k.Close()

	// 通过Windows的注册表获取系统采用的日期格式
	DateFormat, _, err := k.GetStringValue("sShortDate")
	if err != nil {
		panic(err.Error())
	}

	// 通过Windows的注册表获取系统采用的时间格式
	TimeFormat, _, err := k.GetStringValue("sTimeFormat")
	if err != nil {
		panic(err.Error())
	}
	ts.TimeFormat = TimeFormat
	ts.DateFormat = DateFormat
}

// CreateTask 创建任务计划
func (ts *TaskSchedObject) CreateTask(taskName string, execPath string, startTime string) (string, string, error) {
	// taskName = `"` + taskName + `"`
	// execPath = `"` + execPath + `"`
	endDate := "2025/12/01"
	args := "/Create /TN " + taskName + " /TR " + execPath + " /SC daily /ST " + startTime + " /ED " + endDate
	arglist := strings.Split(args, " ")
	_, stderr, err := cmdRun("schtasks", arglist[0:]...)
	if err != nil {
		return "", "", err
	}
	if stderr != "" {
		return "", convertGBK2Str(stderr), nil
	}
	return "create succcessful.", "", nil
}

// QueryTask 查询任务计划
func (ts *TaskSchedObject) QueryTask(taskName string) (string, string, error) {
	args := "/Query /tn " + taskName
	arglist := strings.Split(args, " ")
	// cmdPath := os.e
	stdout, stderr, err := cmdRun("schtasks", arglist[0:]...)
	if err != nil {
		return convertGBK2Str(stdout), convertGBK2Str(stderr), err
	}
	if stderr != "" {
		return convertGBK2Str(stdout), convertGBK2Str(stderr), errors.New(stderr)
	}
	return convertGBK2Str(stdout), convertGBK2Str(stderr), nil
}

func cmdRun(cmder string, args ...string) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd := exec.Command(cmder, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// ConvertStr2GBK 转换字符串
func convertStr2GBK(str string) string {
	ret, err := simplifiedchinese.GBK.NewEncoder().String(str)
	if err != nil {
		panic(err.Error())
	}
	return ret

	// b, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	// return string(b)
}

func convertGBK2Str(gbkStr string) string {
	ret, err := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
	if err != nil {
		panic(err.Error())
	}
	return ret
}

func errorHandle(err error){
	if err != nil {
		panic(err)
	}
}

func lengthVerified(numStr string, start, end int) bool {
	if numStr == "" || start < 0 || end > 9 {
		return false
	}
	length := len(numStr)
	if length < start && length > end {
		return false
	} 
	return true
}

func numericVerified(numStr []string) bool  {
	var (
		symbolPattern  = "[-a-z-A-Z`~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]"
		numericPattern = `[\d+]`
	)
	numericReg1 := regexp.MustCompile(symbolPattern)
	numericReg2 := regexp.MustCompile(numericPattern)
	for _, v := range numStr {
		if numericReg1.MatchString(v){
			return false
		}
		if !numericReg2.MatchString(v){
			return false
		}
	}
	return true
}