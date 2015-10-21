// standard file server to serve up HTML/JS
// includes blob emitter to test XHR request speed

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(cwd)))

	fmt.Println("http://127.0.0.1:10805/index.html")
	err = http.ListenAndServe(":10805", nil)
	if err != nil {
		panic(err)
	}
}
