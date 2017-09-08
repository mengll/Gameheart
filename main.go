// Gameheart project main.go
package main

import (
	"Gameheart/controller"
	"flag"
	"fmt"
	"net/http"
	"runtime"
)

//the app init
var port string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.StringVar(&port, "port", ":8080", "the server port")
	fmt.Println("Hello World!")
	http.HandleFunc("/heart", controller.GameHeartRequest)
	http.ListenAndServe(port, nil)
}
