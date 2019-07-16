package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// response, err2 := ioutil.ReadAll(resp.Body)
	// // var response []byte
	// response := make([]byte, 99999)
	// resp.Body.Read(response)
	// // if err2 != nil {
	// // 	fmt.Println("Error: ", err2)
	// // 	os.Exit(1)
	// // }
	// fmt.Println(string(response))
	// resp.Body.Close()
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fin := string(bs)
	fmt.Println(fin)
	fmt.Println("Just wrote so many bytes", len(bs))
	return len(bs), nil
}
