package server

import (
	"../lib"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const IntervalSec = 5 * time.Second
const Port = ":8888"

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.RemoteAddr + "より送信を確認しました")
	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("uploadfile")

	sep_file_name := strings.Split(handler.Filename, "\\")
	filename := sep_file_name[len(sep_file_name)-1]

	sep_file_name = strings.Split(filename, "/")
	filename = sep_file_name[len(sep_file_name)-1]

	fmt.Println(filename + "を受信しました")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filename = lib.AppendSuffixNumberIfExistsFile(filename)
	f, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	fmt.Println("保存が完了いたしました")

}

func Run(server_alive chan bool) {
	fmt.Println("### server Start on" + Port + " ###")

	defer func() {
		server_alive <- false
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(Port, nil)
}
