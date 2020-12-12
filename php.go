package tools

import (
	"crypto/md5"
	"fmt"
	"github.com/PeterYangs/tools/com"
	"io"
	"regexp"
	"strings"
	"time"
)

//php中的常用函数，go重写

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

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

	fmt.Println(r1)

	times, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)

	return times.Unix()

}
