package client

import (
	"../lib"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"
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
		return getList()
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

func getList() error {
	var rlen int

	ServerAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:8889")
	if err != nil {
		fmt.Println("Error: %v\n", err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		fmt.Println("Error: %v\n", err)
	}

	conn, err := net.ListenUDP("udp", LocalAddr)
	if err != nil {
		fmt.Println("Error: %v\n", err)
	}

	conn.SetDeadline(time.Now().Add(3 * time.Second))
	defer func() {
		fmt.Println("完了")
		os.Exit(0)
	}()
	defer conn.Close()

	s := lib.CurrentUserName()

	rlen, err = conn.WriteTo([]byte(s), ServerAddr)

	if err != nil {
		fmt.Printf("Send Error: %v\n", err)
		return err
	}

	fmt.Printf("Send: %v\n", s)

	buf := make([]byte, 1024)

	for {
		rlen, _, err = conn.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Receive Error: %v\n", err)
			return err
		}

		fmt.Printf("Receive: %v\n", string(buf[:rlen]))
	}

	return nil
}

func Main(commands []string) error {
	fmt.Println("### client Start ###")
	return switchMethod(commands...)
}
