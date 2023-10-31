package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func openFile(path string) (*os.File, func(), error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("can't open file: ", err)
	}
	closer := func() { f.Close() }
	return f, closer, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("can't execute program, need more arguments")
	}
	file, closer, _ := openFile(os.Args[1])
	defer closer()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("can't read data from file: ", err)
	}
	fmt.Println(string(data))
}
