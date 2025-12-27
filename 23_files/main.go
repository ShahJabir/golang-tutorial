package main

import (
	"fmt"
	"os"
)

func fileInfo(filePath string) os.FileInfo {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return fileInfo
}

func readDir(dirPath string) []os.FileInfo {
	f, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfos, err := f.Readdir(-1)
	if err != nil {
		panic(err)
	}
	return fileInfos
}

func readFile(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo := fileInfo(filePath)
	data := make([]byte, fileInfo.Size())
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func createFile(filePath string, data []byte) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get file information
	info := fileInfo("./example.txt")
	fmt.Println("File Name:", info.Name())
	fmt.Println("File Size:", info.Size())
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Println("Last Modified:", info.ModTime())

	// Read file content
	data := readFile("./example.txt")
	fmt.Println("File Content:", string(data))

	// Read directory contents
	fileInfos := readDir(".")
	fmt.Println("Directory Contents:")
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}

	// Create a new file
	createFile("./newfile.txt", []byte("Hello, World!"))
	fmt.Println("New file created: newfile.txt")
}
