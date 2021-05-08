# tools

go的工具集

**安装**

go get github.com/PeterYangs/tools

<br/>

**1.网络请求**




```

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


```


<br/>
<br/>

**2.时间处理**

```
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

```

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

```

//读取文件
data, err :=tools.ReadFile("http.go")




//创建多级文件夹
path := "a/b/c/"

tools.MkDirDepth(path)





//获取文件拓展名
name := "1212.png"

f1, _ := tools.GetExtensionName(name)

fmt.Println(f1)


```

<br/>
<br/>

**5.字符串操作**

```

//explode
str:="1,2,3"

fmt.Println(tools.Explode(",",str))




//md5
str := "123"

fmt.Println(tools.Md5(str))




//字符串截取
str := "我尼玛"

tools.SubStr(str, 1, 1)
tools.SubStr(str, 1, -1)



```

