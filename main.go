package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	log.Println("Starting")

	dir := os.Args[1]
	dirpath := filepath.Join(dir, "*.jpg")
	matches, err := filepath.Glob(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	size := os.Args[2]

	// twelve hundred or five hundred
	dirResize := filepath.Join(dir, size)
	os.Mkdir(dirResize, 0755)

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

		sizeInt, err := strconv.ParseUint(size, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		m := resize.Resize(uint(sizeInt), 0, img, resize.Lanczos3)

		basename := filepath.Base(value)
		newpath := filepath.Join(dirResize, basename)
		out, err := os.Create(newpath)
		if err != nil {
			log.Fatal(err)
		}

		jpeg.Encode(out, m, nil)
		out.Close()

		log.Println(newpath)
	}
}
