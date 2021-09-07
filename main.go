// GOROOT_BOOTSTRAP=/usr/local/go
// GOOS=js GOARCH=wasm go build -o lib.wasm build.go
package main

import (
	"fmt"
	"mp-color/myprocessor"
	"syscall/js"
)

func main() {
	c1 := make(chan struct{}, 0)
	registerCallbacks()
	//println("Init wasm2")
	<-c1
}

func process(this js.Value, data []js.Value) interface{} {
	//println("data : ", fmt.Sprintf("%s", data[0]))
	result := myprocessor.Do(fmt.Sprintf("%s", data[0]))
	//log.Println(fmt.Sprintf("%v", result))
	return result
}

func registerCallbacks() {
	js.Global().Set("process", js.FuncOf(process))
}

// GOROOT_BOOTSTRAP=/usr/local/go
// GOOS=js GOARCH=wasm go build -o lib.wasm main.go
// package main

// import (
// 	"log"
// 	"mp-color/myprocessor"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	myprocessor.Do("data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wAARCAJ/AyADASIAAhEBAxEB/8QAHQAAAgIDAQEBAAAAAAAAAAAABQYEBwMICQIBAP/")
// 	log.Printf("%s", time.Since(start))
// }
