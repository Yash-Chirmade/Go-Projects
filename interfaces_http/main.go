package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
type logWriter struct{}

func main() {
	resp, err := http.Get("http://gooogle.com")
	if err != nil {
		fmt.Print("error", err)
		os.Exit(1)
	}
	// byteSlice := make([]byte, 99999)
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))


	lw:=logWriter{}

	io.Copy(lw,resp.Body)//inbuilt writer

	// Custom writer

}
func (logWriter) Write(bs []byte)(int,error){
	fmt.Println(string(bs))
	fmt.Println("length",len(bs))
	return len(bs),nil
}