package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
	"github.com/PeterYangs/tools/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var num int32

func main() {

	client := http.Client()

	client.SetTimeout(1 * time.Minute)

	wait := sync.WaitGroup{}

	startTime := time.Now()

	for i := 0; i < 100; i++ {

		wait.Add(1)

		go func(w *sync.WaitGroup) {

			defer wait.Done()

			for {

				nums := tools.MtRand(1, 10)

				res, err := client.Request().GetToString("http://www.177.com/test1?id=1&num=" + strconv.Itoa(int(nums)))

				fmt.Println(res)

				if err != nil || res == "卖光啦！！" {

					break
				}

				if res == "success" {

					//fmt.Println("success")

					//num++

					atomic.AddInt32(&num, int32(nums))
				}

			}

			//fmt.Println(res)

		}(&wait)

	}

	wait.Wait()

	endTime := time.Now()

	diff := endTime.Sub(startTime)

	fmt.Println("执行完毕,耗时：", diff.Seconds())

	fmt.Println(num)

	//client_ = http.Client(http.Setting{})
	//
	//max := make(chan int, 30)
	//
	//for {
	//
	//	max <- 1
	//
	//	go run(max)
	//
	//}

}

//func run(max chan int) {
//
//	defer func(maxs chan int) {
//
//		<-maxs
//
//	}(max)
//
//	h, err := client_.GetToString("https://www.youxi369.com/")
//
//	if err != nil {
//
//		fmt.Println(err)
//
//		return
//
//	}
//
//	fmt.Println(h)
//
//}
