package main

import (
	"fmt"
	"io"
	"os"
)
func main(){
	filename:=os.Args[1]
	file,err:=os.Open(filename)
	if err!=nil{
		fmt.Print(err)
		os.Exit(0)
	}
	io.Copy(os.Stdout, file)
}

//Output:
// go run main.go sample.txt
// this is a sample text file to read by the reader.