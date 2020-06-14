package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/h2non/filetype"
)

var validFormats = []string{
	"image/jpeg",
	"image/png",
	"image/tiff",
	"image/webp",
	"image/gif",
}

func main() {

	var dir = flag.String("d", ".", "Directory where to resize images")
	var width = flag.Int("w", 1224, "Desired image width, in pixels. Setting it to 0 will keep the ratio.")
	var height = flag.Int("h", 0, "Desired image height, in pixels. Setting it to 0 will keep the ratio.")
	var quality = flag.Int("q", 80, "Desired image quality, from 0 to 100 (lower to better)")

	flag.Parse()

	var files []string

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		ok, mime := hasValidFormat(file)
		if ok {
			resizeImage(file, *width, *height, *quality, mime)
		}
	}
}

// hasValidFormat check if the file has an accepted format and also returns the MIME type
func hasValidFormat(fileName string) (valid bool, mime string) {

	fileMIME, err := getImageMIME(fileName)
	if err != nil {
		return false, ""
	}

	for _, v := range validFormats {
		if fileMIME == v {
			return true, fileMIME
		}
	}

	return false, ""
}

// resizeImage will resize the filename to de desired width and quality
func resizeImage(fileName string, width int, height int, quality int, mime string) {

	src, err := imaging.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	if width > 0 || height > 0 {
		bounds := src.Bounds()
		imageWidth := bounds.Dx()
		imageHeigth := bounds.Dy()
		if imageWidth > width || imageHeigth > height {
			src = imaging.Resize(src, width, height, imaging.Lanczos)
		}
	}

	switch mime {
	case "image/jpeg", "image/jpg":

		err = imaging.Save(src, fileName, imaging.JPEGQuality(quality))

	case "image/png":

		compression := png.NoCompression
		if quality > 80 && quality < 100 {
			compression = png.DefaultCompression
		} else if quality > 60 && quality < 80 {
			compression = png.BestSpeed
		} else if quality > 60 {
			compression = png.BestCompression
		}

		err = imaging.Save(src, fileName, imaging.PNGCompressionLevel(compression))

	default:

		err = imaging.Save(src, fileName)
	}

	if err != nil {
		fmt.Printf("failed to save image: %v", err)
	}

	fmt.Printf("File %s (%s) was modified\n", fileName, mime)

}

// getImageMIME returns the MIME type of the file
func getImageMIME(fileName string) (mime string, err error) {

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		return "", err
	}

	return kind.MIME.Value, nil
}
