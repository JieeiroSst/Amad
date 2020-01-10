package images

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploading file ")
	//1 parse input multipart/from data
	r.ParseMultipartForm(10 << 20)
	//2 retrive file from posted from data
	file, handler, err := r.FormFile("uploads")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fmt.Printf("upload file %+v\n", handler.Filename)
	fmt.Printf("file size %+v\n", handler.Size)
	fmt.Printf("mimi header%+v\n ", handler.Header)
	//write tempory file  to server
	tempFile, err := ioutil.TempFile("uploads", "uploads-*.png")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tempFile.Close()
	fileByes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	tempFile.Write(fileByes)
	//return all file images
	fmt.Println(tempFile)
}
