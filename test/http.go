package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
	"log"
)

func main() {

	p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

	//str, err :=tools.PostToString("http://list.com/index.php",tools.HttpSetting{Parameter: p})
	str, err := tools.GetToString("http://list.com/index.php", tools.HttpSetting{Parameter: p})

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println(str)

}

//str, err :
