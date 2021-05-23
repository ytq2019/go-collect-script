package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main10() {
	f, err := os.Open("./pics.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		download(string(b))
	}
}

func download(imagPath string) {
	fmt.Println("正在下载:", imagPath)
	imagPath = strings.ReplaceAll(imagPath, "\n", "")
	//通过http请求获取图片的流文件
	resp, _ := http.Get(imagPath)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fullPath := strings.ReplaceAll(imagPath, "https://w7.dapp100.cn/", "")
	//fileName := path.Base(dir)
	dir := path.Dir(fullPath)
	if err := CreateMutiDir(dir); err != nil {
		panic(err)
	}

	out, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	io.Copy(out, bytes.NewReader(body))
	return
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//调用os.MkdirAll递归创建文件夹
func CreateMutiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
