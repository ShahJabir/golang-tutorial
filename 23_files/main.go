package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func writeFile(filePath string, data []byte) {
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

func removeFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(filePath, "Deleted successfully")
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

	// Write in file
	writeFile("newFile.txt", []byte("Hello, World!"))

	// Stream file reading using bufio and then writing to another file
	SourceFile, err := os.Open("./example.txt")
	if err != nil {
		panic(err)
	}
	// defer SourceFile.Close()

	DestFile, err := os.Create("./streamFile.txt")
	if err != nil {
		panic(err)
	}
	// defer DestFile.Close()

	reader := bufio.NewReader(SourceFile)
	writer := bufio.NewWriter(DestFile)

	for {
		n, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		WriteErr := writer.WriteByte(n)
		if WriteErr != nil {
			panic(WriteErr)
		}
	}
	writer.Flush()
	DestFile.Close()
	SourceFile.Close()

	fmt.Println("File copied successfully using buffered I/O")

	// Remove File
	time.Sleep(time.Second)
	removeFile("streamFile.txt")
	removeFile("newFile.txt")
}
