package tools

import (
	"github.com/PeterYangs/tools/com"
	"regexp"
	"strings"
	"time"
)

//时间格式函数
func Date(format string, timestamp int64) string {

	//去空格
	format = strings.Trim(format, " ")

	t := time.Unix(timestamp, 0)
	//f := "2006-01-02 15:04:05"
	f := strings.Replace(format, "Y", "2006", -1)
	f = strings.Replace(f, "m", "01", -1)
	f = strings.Replace(f, "d", "02", -1)
	f = strings.Replace(f, "H", "15", -1)
	f = strings.Replace(f, "i", "04", -1)
	f = strings.Replace(f, "s", "05", -1)
	f = strings.Replace(f, "w", "Monday", -1)

	return t.Format(f)

}

func StrToTime(str string) int64 {

	//去空格
	str = strings.Trim(str, " ")

	r1 := regexp.MustCompile(`^(\d+)[-/](\d+)[-/](\d+)[\s]*(\d*)[:]*(\d*)[:]*(\d+)*$`).FindStringSubmatch(str)

	switch true {

	case len(r1) > 0:
		//拼接模板格式
		str = r1[1] + "-" + r1[2] + "-" + r1[3] + " " + com.GetString(r1[4], "00") + ":" + com.GetString(r1[5], "00") + ":" + com.GetString(r1[6], "00")

	}
	times, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)

	return times.Unix()

}
