package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var file = []struct{ Name, Body string }{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, v := range file {
		hdr := &tar.Header{
			Name: v.Name,
			Mode: 0600,
			Size: int64(len(v.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(v.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.

	tar.FormatGNU.String()
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
		// if _, err := io.Copy(os.Stdout, tr); err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("\n-------")
	}

}
