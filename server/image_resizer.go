package main

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
)

type ImageResizer struct{}

func (i *ImageResizer) GetInterpolationFunctionByName(name string) vips.Kernel {
	switch name {
	default:
		return vips.KernelNearest
	case "nearest":
		return vips.KernelNearest
	case "bilinear":
		return vips.KernelLinear
	case "bicubic":
		return vips.KernelCubic
	case "mitchell-netravali":
		return vips.KernelMitchell
	case "lanczos2":
		return vips.KernelLanczos2
	case "lanczos3":
		return vips.KernelLanczos3
	}
}

func (i *ImageResizer) ResizeImageToOneDimension(img *vips.ImageRef, dimensionPixel int, interpolation vips.Kernel) error {
	if img == nil {
		return fmt.Errorf("img must not be nil")
	}

	if dimensionPixel < 1 {
		return fmt.Errorf("dimensionPixel must be greater than 0")
	}

	// Get the original image dimensions
	original_width := img.Width()
	original_height := img.Height()

	if (original_width <= dimensionPixel) && (original_height <= dimensionPixel) {
		return nil
	}

	scale := 1.0

	if original_width > original_height {
		scale = float64(dimensionPixel) / float64(original_width)
	} else {
		scale = float64(dimensionPixel) / float64(original_height)
	}

	// Resize the image
	err := img.Resize(scale, interpolation)
	if err != nil {
		return fmt.Errorf("image resize fail")
	}

	return nil
}
