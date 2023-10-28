package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func relaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	cmd.Dir = ""
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Automatically!</h1>")
	relaunch()
}
func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
