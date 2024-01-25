package main

import (
	"fmt"
	"image"

	"go.oneofone.dev/resize"
)

type ImageResizer struct{}

func (i *ImageResizer) GetInterpolationFunctionByName(name string) resize.InterpolationFunction {
	switch name {
	default:
		return resize.NearestNeighbor
	case "nearest":
		return resize.NearestNeighbor
	case "bilinear":
		return resize.Bilinear
	case "bicubic":
		return resize.Bicubic
	case "mitchell-netravali":
		return resize.MitchellNetravali
	case "lanczos2":
		return resize.Lanczos2
	case "lanczos3":
		return resize.Lanczos3
	}
}

func (i *ImageResizer) ResizeImageToOneDimension(img image.Image, dimensionPixel int, interpolation resize.InterpolationFunction) (image.Image, error) {
	if img == nil {
		return nil, fmt.Errorf("img must not be nil")
	}

	if dimensionPixel < 1 {
		return nil, fmt.Errorf("dimensionPixel must be greater than 0")
	}

	// Get the original image dimensions
	original_width := img.Bounds().Dx()
	original_height := img.Bounds().Dy()

	if (original_width <= dimensionPixel) && (original_height <= dimensionPixel) {
		return img, nil
	}

	// Resize the image
	resized_image := resize.Thumbnail(uint(dimensionPixel), uint(dimensionPixel), img, interpolation)

	return resized_image, nil
}
