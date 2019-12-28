package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func main() {
	ch := make(chan int, 4)
	wg := sync.WaitGroup{}

	files, _ := ioutil.ReadDir("./image")
	fmt.Println(len(files))
	for _, f := range files {
		ch <- 1
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			save_image("./image/" + name)
			<-ch
		}(f.Name())
	}
	wg.Wait()
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
	m := resize.Resize(1280, 0, img, resize.Lanczos3)

	if data == "png" {
		out, err := os.Create(path + uuid + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		png.Encode(out, m)
	} else if data == "jpeg" {
		out, err := os.Create(path + uuid + ".jpeg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
	}

}

func create_uuid() string {
	u, _ := uuid.NewRandom()
	uu := u.String()
	return uu
}
