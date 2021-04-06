# tools
go的工具集


**安装**

go get github.com/PeterYangs/tools

<br/>

**1.网络请求**

```
//get请求
str, err := tools.GetToString("http://www.baidu.com")

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


```

