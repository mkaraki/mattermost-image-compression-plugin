package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"

	"github.com/chai2010/webp"
)

type ImageEncoder struct{}

func (i *ImageEncoder) EncodeWebP(img image.Image, output io.Writer, quality int) (int64, error) {
	if (img == nil) || (output == nil) {
		return 0, fmt.Errorf("img and output must not be nil")
	}

	if (quality < 0) || (quality > 100) {
		println("Invalid image quality value, use 100.")
		quality = 90
	}

	webp_data, err := webp.EncodeRGB(img, float32(quality))
	if err != nil {
		println("Failed to encode image as webp", err.Error())
		return 0, err
	}

	output_size := len(webp_data)

	_, err = output.Write(webp_data)

	return int64(output_size), err
}

func (i *ImageEncoder) EncodeJpeg(img image.Image, output io.Writer, quality int) (int64, error) {
	if (img == nil) || (output == nil) {
		return 0, fmt.Errorf("img and output must not be nil")
	}

	if (quality < 0) || (quality > 100) {
		println("Invalid image quality value, use 100.")
		quality = 90
	}

	jpeg_data := &bytes.Buffer{}
	jpeg_writer := io.MultiWriter(output, jpeg_data)

	err := jpeg.Encode(jpeg_writer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		println("Failed to encode image as webp", err.Error())
		return 0, err
	}

	output_size := jpeg_data.Len()

	return int64(output_size), nil
}
