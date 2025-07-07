package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var image = ImageToBase64("/home/lard/GolandProjects/exercises/man.png")
	fmt.Println(image)
}

func ImageToBase64(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(nil)
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}
