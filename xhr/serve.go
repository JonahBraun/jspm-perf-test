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

	http.HandleFunc("/blob.js", emitBlob)
	http.Handle("/", http.FileServer(http.Dir(cwd)))

	fmt.Println("http://127.0.0.1:10808/index.html")
	err = http.ListenAndServe(":10808", nil)
	if err != nil {
		panic(err)
	}
}

func emitBlob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")

	// send 5kb blob
	for i := 0; i < 5*16; i++ {
		w.Write(blob)
	}
}

// a nice 64 byte blob
var blob = []byte("var foo='foofoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfo';\n")
