package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func switchMethod(commands ...string) error {
	var cmd string
	var targetIp string

	if len(commands) == 2 {
		cmd = commands[1]
	} else {
		targetIp = commands[1]
		cmd = commands[2]
	}

	switch cmd {
	case "list":
		// 未実装 UDPでipのリストをもらう 名前とip
		return nil
	default:
		return postFile(cmd, "http://"+targetIp+":8888")
	}
}

func postFile(filename string, targetUrl string) error {
	fmt.Println(filename + "を送信します")
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//キーとなる操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//ファイルハンドル操作をオープンする
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	fmt.Println("##################")
	fmt.Println(targetUrl + "へ送信します")
	fmt.Println("送信ファイル名" + filename)
	fmt.Println("##################")
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func Main(commands []string) error {
	fmt.Println("### client Start ###")
	return switchMethod(commands...)
}