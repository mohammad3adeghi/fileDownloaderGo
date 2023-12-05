package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

const links = "https://img.freepik.com/free-photo/painting-mountain-lake-with-mountain-background_188544-9126.jpg"

func main() {
	resp, err := http.Get(links)
	if err != nil {
		log.Fatal()
	}
	defer resp.Body.Close()

	f, err := os.Create(path.Base(links))

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

}