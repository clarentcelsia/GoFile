package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/nfnt/resize"
)

func ImageResize(filepath string) error {
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(bytes.NewReader(filebytes))
	if err != nil {
		return err
	}

	// Lanczos3 : preserve the quality
	img = resize.Resize(800, 600, img, resize.Lanczos3)

	var writeBytes bytes.Buffer
	if err = jpeg.Encode(&writeBytes, img, nil); err != nil {
		return err
	}

	if err = ioutil.WriteFile("FILE.jpeg", writeBytes.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
