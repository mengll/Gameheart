// Gameheart project main.go
package main

import (
	"fmt"
	"net/http"
	"runtime"
)

//the app init

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main() {
	fmt.Println("Hello World!")
	servbe ;= http.HandleFunc("/heart",)
	

}
