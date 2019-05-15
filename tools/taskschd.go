package bingwallpapers

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	registry "github.com/golang/sys/windows/registry"
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

// DatesFormat 将日期格式化为系统的日期格式
func (ts *TaskSchedObject) DatesFormat(year, month, day string) (string, error) {
	// start0 := time.Now()
	var (
		yearlen        = len(year)
		monthlen       = len(month)
		daylen         = len(day)
		err            error
		msg            string
		symbolStr      [4]bool
		symbolPattern  = "[-a-z-A-Z`~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]"
		numericPattern = `[\d+]`
	)
	// 数值字符串合法性验证
	if (yearlen != 2 && yearlen != 4) || (monthlen > 2) || (daylen > 2) {
		msg = `year format must "yy" or "yyyy" and month must M or MM and day must D or DD  .`
		return "", errors.New(msg)
	}
	msg = `args: "year", "month", "day" string must numeric.`
	symbolReg := regexp.MustCompile(symbolPattern)
	numericReg := regexp.MustCompile(numericPattern)
	symbolStr[0] = symbolReg.MatchString(year)
	symbolStr[1] = symbolReg.MatchString(month)
	symbolStr[2] = symbolReg.MatchString(day)
	symbolStr[3] = !numericReg.MatchString(year)
	for i := 0; i < len(symbolStr); i++ {
		if symbolStr[i] {
			return "", errors.New(msg)
		}
	}
	// 对字符串进行转换，检测是否符合月/日的大小范围
	intMonth, err := strconv.Atoi(month)
	intDay, err := strconv.Atoi(day)
	if err != nil {
		panic(err.Error())
	}
	if intMonth > 12 || intMonth == 0 || intDay > 31 || intDay == 0 {
		msg = "month or day number over range."
		panic(errors.New(msg))
	}
	// 将日期字符串与系统的日期格式进行匹配处理
	var (
		datef   string
		dateSep string
		result  string
	)

	datef = ts.DateFormat
	reg := regexp.MustCompile("[/.-]")
	sepstr := reg.FindAllString(datef, 1)
	if len(sepstr) < 1 {
		return "", errors.New("unknow date separator")
	}
	switch sepstr[0] {
	case "/":
		dateSep = "/"
	case ".":
		dateSep = "."
	case "-":
		dateSep = "-"
	}
	dateElems := strings.Split(datef, dateSep)
	mmdd := func(elem string) {
		switch elem {
		case "M":
			// yyyy or yy/M/d
			if len(month) > 1 {
				month = month[1:]
			}
			if len(day) > 1 {
				day = day[1:]
			}
		case "MM":
			// yyyy or yy/MM/dd
			if len(month) == 1 {
				month = "0" + month
			}
			if len(day) == 1 {
				day = "0" + day
			}
		case "yy":
			if len(year) > 2 {
				year = year[2:]
			}
		}
	}
	if dateElems != nil {
		switch dateElems[0] {
		case "M":
			// M/d/yy or yyyy
			mmdd(dateElems[0])
			mmdd(dateElems[2])
			result = month + dateSep + day + dateSep + year
		case "MM":
			// MM/dd/yy or yyyy
			mmdd(dateElems[0])
			mmdd(dateElems[2])
			result = month + dateSep + day + dateSep + year
		case "yy":
			// yy/M or MM/d or dd
			mmdd(dateElems[0])
			mmdd(dateElems[1])
			result = year + dateSep + month + dateSep + day
		case "yyyy":
			// yyyy/M or MM/d or dd
			mmdd(dateElems[1])
			result = year + dateSep + month + dateSep + day
		}
	}
	// end0 := time.Now().Sub(start0)
	// fmt.Println(end0)
	return result, nil
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
