package tools

import (
	"github.com/PeterYangs/tools/com"
	"regexp"
	"strings"
	"time"
)

//先用比较蠢的办法
func Date(format string, timestamp int64) string {

	//去空格
	format = strings.Trim(format, " ")

	t := time.Unix(timestamp, 0)

	f := "2006-01-02 15:04:05"

	switch format {

	case "Y-m-d":

		f = "2006-01-02"

	case "Y/m/d":

		f = "2006/01/02"

	case "Y-m":

		f = "2006-01"

	case "Y/m":

		f = "2006-01"

	case "m-d":

		f = "01-02"

	case "m/d":
		f = "01/02"

	case "H:i:s":

		f = "15:04:05"

	case "H:i":

		f = "15:04"

	case "Y-m-d H:i":

		f = "2006-01-02 15:04"

	case "Y/m/d H:i":

		f = "2006/01/02 15:04"
	case "Y":

		f = "2006"
	case "Ym":

		f = "200601"
	case "Ymd":

		f = "20060102"
	case "H":
		f = "15"
	case "Hi":
		f = "1504"

	case "md":

		f = "0102"

	}

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
