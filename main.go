package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Starting")

	dir := os.Args[1]
	dirpath := filepath.Join(dir, "*.jpg")
	matches, err := filepath.Glob(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	dir500 := filepath.Join(dir, "1200")
	os.Mkdir(dir500, 0755)

	log.Println(len(matches))

	for _, value := range matches {
		file, err := os.Open(value)
		if err != nil {
			log.Fatal(err)
		}

		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		m := resize.Resize(1200, 0, img, resize.Lanczos3)

		basename := filepath.Base(value)
		newpath := filepath.Join(dir500, basename)
		out, err := os.Create(newpath)
		if err != nil {
			log.Fatal(err)
		}

		jpeg.Encode(out, m, nil)
		out.Close()

		log.Println(newpath)
	}
}
