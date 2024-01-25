package main

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.oneofone.dev/resize"
)

func TestGetInterpolationFunctionByName(t *testing.T) {
	resizer := &ImageResizer{}

	// Test with "nearest"
	funcName := "nearest"
	expectedFunc := resize.NearestNeighbor
	actualFunc := resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'nearest'")

	// Test with "bilinear"
	funcName = "bilinear"
	expectedFunc = resize.Bilinear
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'bilinear'")

	// Test with "bicubic"
	funcName = "bicubic"
	expectedFunc = resize.Bicubic
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'bicubic'")

	// Test with "mitchell-netravali"
	funcName = "mitchell-netravali"
	expectedFunc = resize.MitchellNetravali
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'mitchell-netravali'")

	// Test with "lanczos2"
	funcName = "lanczos2"
	expectedFunc = resize.Lanczos2
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'lanczos2'")

	// Test with "lanczos3"
	funcName = "lanczos3"
	expectedFunc = resize.Lanczos3
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'lanczos3'")

	// Test with default case
	funcName = "nonexistent"
	expectedFunc = resize.NearestNeighbor
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for default case")
}

func TestResizeImageToOneDimension(t *testing.T) {
	resizer := &ImageResizer{}

	// Create a simple 2x2 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	// Test with nil image
	_, err := resizer.ResizeImageToOneDimension(nil, 1, resize.NearestNeighbor)
	assert.Error(t, err, "ResizeImageToOneDimension should fail with nil image")

	// Test with dimensionPixel less than 1
	_, err = resizer.ResizeImageToOneDimension(img, 0, resize.NearestNeighbor)
	assert.Error(t, err, "ResizeImageToOneDimension should fail with dimensionPixel less than 1")

	// Test with valid parameters
	resizedImg, err := resizer.ResizeImageToOneDimension(img, 1, resize.NearestNeighbor)
	assert.NoError(t, err, "ResizeImageToOneDimension failed with valid parameters")
	assert.Equal(t, 1, resizedImg.Bounds().Dx(), "Resized image width does not match the expected width")
	assert.Equal(t, 1, resizedImg.Bounds().Dy(), "Resized image height does not match the expected height")

	// Test with dimensionPixel greater than original dimensions
	resizedImg, err = resizer.ResizeImageToOneDimension(img, 3, resize.NearestNeighbor)
	assert.NoError(t, err, "ResizeImageToOneDimension failed with dimensionPixel greater than original dimensions")
	assert.Equal(t, 2, resizedImg.Bounds().Dx(), "Resized image width does not match the original width")
	assert.Equal(t, 2, resizedImg.Bounds().Dy(), "Resized image height does not match the original height")
}
