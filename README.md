# tools

go的工具集

**安装**

go get github.com/PeterYangs/tools

<br/>

**1.网络请求**

```
//get请求
str, err := tools.GetToString("http://www.baidu.com",tools.HttpSetting{})

if err != nil {

    log.Fatal(err)

}

fmt.Println(str)





//携带参数
p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

str, err := tools.GetToString("http://www.baidu.com",tools.HttpSetting{Parameter: p})

if err != nil {

    log.Fatal(err)

}

fmt.Println(str)




//post
p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

str, err := tools.PostToString("http://www.baidu.com",tools.HttpSetting{Parameter: p})

if err != nil {

    log.Fatal(err)

}

fmt.Println(str)






//自定义header
header:=map[string]string{"user-agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"}

tools.GetToString("https://www.baidu.com",tools.HttpSetting{Header: header})





//下载图片(会判断header的类型)
img := "https://item-shopping.c.yimg.jp/i/n/sakuranokoi_udj008_1"
	
err := tools.DownloadImage(img, "image.png", tools.HttpSetting{})






//下载文件
img := "https://item-shopping.c.yimg.jp/i/n/sakuranokoi_udj008_1"

err := tools.DownloadFile(img, "image.png", tools.HttpSetting{})





//proxy
url := "https://www.google.com/"

html, err := tools.GetToString(url, tools.HttpSetting{
		
    ProxyAddress: "socks5://127.0.0.1:4781",
})

if err != nil {

    fmt.Println(err)
}

fmt.Println(html)


//超时设置
tools.GetToString("https://www.baidu.com",tools.HttpSetting{TimeOut: 10})


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

