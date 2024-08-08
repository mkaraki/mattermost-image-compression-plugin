package main

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"io"
)

type ImageEncoder struct{}

func (i *ImageEncoder) EncodeWebP(img *vips.ImageRef, output io.Writer, quality int) (int64, error) {
	if (img == nil) || (output == nil) {
		return 0, fmt.Errorf("img and output must not be nil")
	}

	if (quality < 0) || (quality > 100) {
		println("Invalid image quality value, use 10.")
		quality = 10
	}

	webp_export_params := vips.WebpExportParams{
		StripMetadata: true,
		Quality:       quality,
		Lossless:      false,
		NearLossless:  false,
	}

	webp_data, _, err := img.ExportWebp(&webp_export_params)
	if err != nil {
		println("Failed to encode image as webp", err.Error())
		return 0, err
	}

	output_size := len(webp_data)

	_, err = output.Write(webp_data)

	return int64(output_size), err
}
