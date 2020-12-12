package com

func GetString(str string, defaults string) string {

	if str == "" {

		return defaults
	}

	return str
}
