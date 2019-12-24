package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func main() {

	files, _ := ioutil.ReadDir("./image")
	for _, f := range files {
		save_image("./image/" + f.Name())
	}

}

func save_image(filename string) {

	path := "./datas/"

	uuid := create_uuid()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	img, data, err := image.Decode(file)
	if err != nil {
		log.Println("hoge")
		log.Fatal(err)
	}
	file.Close()
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	if data == "png" {
		out, err := os.Create(path + uuid + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		fmt.Println("png")
		png.Encode(out, m)
	} else if data == "jpeg" {
		out, err := os.Create(path + uuid + ".jpeg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
		fmt.Println("jpeg")
	}

}

func create_uuid() string {
	u, _ := uuid.NewRandom()
	uu := u.String()
	return uu
}