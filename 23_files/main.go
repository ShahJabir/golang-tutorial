package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./example.txt")
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("File Info:", fileInfo)
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("File Permission:", fileInfo.Mode())
	defer f.Close()
}
