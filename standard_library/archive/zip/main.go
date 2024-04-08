package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a buffer to write our archive to.
	var buf = new(bytes.Buffer)
	// Create a new zip archive.
	w := zip.NewWriter(buf) // 创建zip文件，持续给buf写入
	// Add some files to the archive.
	var files = []struct{ Name, Body string }{ // 定义一个文件列表
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files { // 遍历文件列表，依次写入
		f, err := w.Create(file.Name) // 在buf中创建文件，写入文件名
		if err != nil {
			log.Fatal(1, err)
		}
		_, err = f.Write([]byte(file.Body)) // 在buf中写入文件内容
		if err != nil {
			log.Fatal(2, err)
		}
	}
	// Make sure to check the error on Close.
	if err := w.Close(); err != nil { // 关闭zip文件
		log.Fatal(3, err)
	}

	// Write the archive to a file.
	filePath := "test.zip"
	file, err := os.Create(filePath) // 创建文件
	if err != nil {
		log.Fatal(5, err)
	}
	if _, err := buf.WriteTo(file); err != nil { // 写入文件
		log.Fatal(6, err)
	}
	if err := file.Close(); err != nil { // 关闭文件
		log.Fatal(7, err)
	}

	// Open a zip archive for reading.
	r, err := zip.OpenReader(filePath) // 打开文件
	if err != nil {
		log.Fatal(8, err)
	}
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File { // 遍历文件列表
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open() // 打开文件
		if err != nil {
			log.Fatal(9, err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68) // 读取文件内容，并输出
		if err != io.EOF && err != nil {
			log.Fatal(10, err)
		}
		rc.Close() // 关闭文件
		fmt.Println()
	}
	if err := r.Close(); err != nil { // 关闭文件
		log.Fatal(11, err)
	}
}
