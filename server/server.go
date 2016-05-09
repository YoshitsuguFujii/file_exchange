package server

import (
	"../lib"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const IntervalSec = 5 * time.Second

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
	fmt.Println(handler.Filename + "を受信しました")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filename := lib.AppendSuffixNumberIfExistsFile(handler.Filename)
	f, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	fmt.Println("保存が完了いたしました")

}

func Main(server_alive chan bool) {
	fmt.Println("### server Start ###")

	defer func() {
		server_alive <- false
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
