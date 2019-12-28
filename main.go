package main

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/disintegration/imaging"
)

func main() {
	ch := make(chan int, 4)
	wg := sync.WaitGroup{}

	files, _ := ioutil.ReadDir("./image")
	for _, f := range files {
		ch <- 1
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			save_image("./image/"+name, name)
			<-ch
		}(f.Name())
	}
	wg.Wait()
}

func save_image(filename string, name string) {

	path := "./datas/"

	file, err := imaging.Open(filename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	m := imaging.Resize(file, 1280, 0, imaging.Lanczos)

	err = imaging.Save(m, path+name)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
