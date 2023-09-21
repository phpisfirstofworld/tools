# tools

go的工具集

**安装**
```shell
go get github.com/PeterYangs/tools
```


**1.网络请求**

网络请求已拆分到新仓库https://github.com/PeterYangs/request

```go

import "github.com/PeterYangs/tools/http"

//获取客户端
client := http.Client()

//get请求
str, err := client.Request().GetToString("https://www.baidu.com")

//post请求
str, err := client.Request().PostToString("https://www.baidu.com")



//携带参数
p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

str, err := client.SetTimeout(1 * time.Second).Request().SetParameter(p).GetToString("https://www.google.com/")


//复杂参数
p := map[string]interface{}{
		"name":     []string{"123", "456"},
		"age":      1,
		"nickname": "123",
		"form": map[string]interface{}{
			"one":   "1",
			"two":   "2",
			"three": []string{"123", "456"},
			"four": map[string]interface{}{
				"one": "1",
				"two": "2",
			},
		},
	}
	
client.Request().SetParameter(p).GetToString("http://list.com/pass/get.php")
	


//自定义header
header := map[string]string{"user-agent": "Iphone100"}

//添加全局header	
client:=http.Client().SetHeader(header)

//局部header	
re, err :=client.Request().SetHeader(header).GetToString("http://list.com/pass/header.php")


//proxy
client := http.Client()

client.SetProxyAddress("http://127.0.0.1:4780")

html, err := client.Request().GetToString("https://www.google.com/")


//timeout
client := http.Client()

client.SetTimeout(1*time.Second)

html, err := client.Request().GetToString("https://www.google.com/")


//重试次数，默认为0
html,err:=client.Request().SetReTryTimes(3).GetToString("https://xxxccaacasdad.com")


```


<br/>
<br/>

**2.时间处理**

```go
//时间戳转时间格式，目前仅支持Y、m、d、H、i、s、w
tools.Date("Y-m-d", time.Now().Unix())
tools.Date("Y-m-d H:i:s", time.Now().Unix()))
tools.Date("Y-m-d H:i", time.Now().Unix())
tools.Date("Y", time.Now().Unix())
tools.Date("Ym", time.Now().Unix()))
tools.Date("Ymd", time.Now().Unix()))
tools.Date("H", time.Now().Unix()))
tools.Date("Hi", time.Now().Unix()))

//时间格式转时间戳，单位秒
tools.StrToTime("2020/12/12")
tools.StrToTime("2020-12-12")
tools.StrToTime("2020-12-12 11:32:00")
tools.StrToTime("2020/12/12 11:32:00")


```

<br/>
<br/>

**3.数组操作**

```go

//in_array
array := []string{"1", "2", "3", "4"}

b := tools.InArray(array, "4")

println(b)

array2 := []int{1, 2, 3, 4}

b2 := tools.InArray(array2, 4)

println(b2)




//implode

arr:=[]string{"1","2","3"}

fmt.Println(tools.Implode("-",arr))



```

<br/>
<br/>

**4.文件操作**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/tools/file"
)

func main() {
    
	//一次性读取
	str, err := file.Read("README.md")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(string(str))


	//逐行读取
	err := file.ReadLine("README.md", func(line []byte) {

		fmt.Println(string(line))

	})

	if err != nil {

		fmt.Println(err)

		return
	}

	//一次性写入
	file.Write("xx.txt", []byte("123"))
}
```

<br/>
<br/>

**5.字符串操作**

```go

//explode
str:="1,2,3"

fmt.Println(tools.Explode(",",str))




//md5
str := "123"

fmt.Println(tools.Md5(str))




//字符串截取
str := "我尼玛"

//起始1，长度1
tools.SubStr(str, 1, 1)

//起始字符串长度倒数第二，长度最大
tools.SubStr(str, -2, -1)

//起始0，长度倒数第二
tools.SubStr(str, 0, -2)



```
**6.3des加密**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/tools/secret"
)

func main() {

	d := secret.NewDes()

	source := "hello world"
	fmt.Println("原字符：", source)

	key := d.GenerateKey() //24位

	//加密
	code, err := d.Encyptog3DES([]byte(source), key)

	if err != nil {

		fmt.Println("加密错误", err)

		return
	}

	fmt.Println("密文：", string(code.ToBase64()))

	//解密
	real, err := d.Decrptog3DES(code.ToBase64(), key, secret.Base64)
	//
	if err != nil {

		fmt.Println("解密错误", err)

		return
	}

	fmt.Println("解密：", string(real))

}
```

**7.文件夹操作**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/tools/file"
)

func main() {

	//获取目标文件夹下的所有文件（包含子目录，返回false则不继续遍历）
	file.GetDirList("./", func(path string) bool {

		fmt.Println(path)

		return true
	})

}
```

**8.Hash**
