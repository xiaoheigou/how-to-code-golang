package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//1.获取要爬的页面内容（string)
	res, err := http.Get("https://www.douyu.com/g_yz")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	buf := make([]byte, 4096)
	var result string
	for {
		m, err := res.Body.Read(buf)
		if m == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return
		}
		result += string(buf[:m])

	}
	//2.使用正则表达式获取爬的图片
	ret := regexp.MustCompile(`<img src="(?s:(.*?))"`)
	alls := ret.FindAllStringSubmatch(result, -1)
	var pictureSlice []string
	for _, v := range alls {
		fmt.Println(v[1])
		pictureSlice = append(pictureSlice, v[1])
	}
	pictureSlice = pictureSlice[:15]
	os.Mkdir("images", 0777)
	for i, v := range pictureSlice {
		path := "images/" + strconv.Itoa(i+1) + ".jpg"
		f, err := os.Create(path)
		fmt.Println(path, "+++++++++++++++++++++++++++++++++++++++++++")
		if err != nil {
			fmt.Println(err, "==========================")
		}
		res, err := http.Get(v)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		buf := make([]byte, 4096)
		for {
			m, err := res.Body.Read(buf)
			if m == 0 {
				break
			}
			if err != nil && err != io.EOF {
				return
			}
			f.Write(buf[:m])
		}

	}

}
