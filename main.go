package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLanch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is my deploy server!</h1>")
	//reLanch()
}

func main() {
	http.HandleFunc("/", firstPage)
	fmt.Println("开始监听5000端口............")
	http.ListenAndServe(":5000", nil)
}
