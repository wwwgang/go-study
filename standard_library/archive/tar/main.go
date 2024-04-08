package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)                 // 创建一个tar.Writer对象，持续给buf写入数据
	var file = []struct{ Name, Body string }{ // 定义一个文件列表
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, v := range file {
		hdr := &tar.Header{ // 创建一个tar.Header对象，设置文件名、大小、权限等属性
			Name: v.Name,
			Mode: 0600,
			Size: int64(len(v.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil { // 将Header写入到tar.Writer对象中
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(v.Body)); err != nil { // 写入文件内容
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil { // 关闭tar.Writer对象
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.

	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n-------")
	}

}
