package tools

import "strings"

//老版兼容函数
func In_array(array interface{}, item interface{}) bool {
	switch key := item.(type) {
	case string:
		for _, item := range array.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range array.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range array.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func InArray(array interface{}, item interface{}) bool {
	switch key := item.(type) {
	case string:
		for _, item := range array.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range array.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range array.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

func Join(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}
