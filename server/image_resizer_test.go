package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterpolationFunctionByName(t *testing.T) {
	resizer := &ImageResizer{}

	// Test with "nearest"
	funcName := "nearest"
	expectedFunc := vips.KernelNearest
	actualFunc := resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'nearest'")

	// Test with "bilinear"
	funcName = "bilinear"
	expectedFunc = vips.KernelLinear
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'bilinear'")

	// Test with "bicubic"
	funcName = "bicubic"
	expectedFunc = vips.KernelCubic
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'bicubic'")

	// Test with "mitchell-netravali"
	funcName = "mitchell-netravali"
	expectedFunc = vips.KernelMitchell
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'mitchell-netravali'")

	// Test with "lanczos2"
	funcName = "lanczos2"
	expectedFunc = vips.KernelLanczos2
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'lanczos2'")

	// Test with "lanczos3"
	funcName = "lanczos3"
	expectedFunc = vips.KernelLanczos3
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for 'lanczos3'")

	// Test with default case
	funcName = "nonexistent"
	expectedFunc = vips.KernelNearest
	actualFunc = resizer.GetInterpolationFunctionByName(funcName)
	assert.Equal(t, expectedFunc, actualFunc, "Expected function does not match the actual function for default case")
}

func TestResizeImageToOneDimension(t *testing.T) {
	resizer := &ImageResizer{}

	// Create a simple 2x2 pixel image
	img, err := vips.Black(2, 2)
	assert.NoError(t, err, "[TEST SPEC FAIL] Failed to create image.")

	// Test with nil image
	err = resizer.ResizeImageToOneDimension(nil, 1, vips.KernelNearest)
	assert.Error(t, err, "ResizeImageToOneDimension should fail with nil image")

	// Test with dimensionPixel less than 1
	err = resizer.ResizeImageToOneDimension(img, 0, vips.KernelNearest)
	assert.Error(t, err, "ResizeImageToOneDimension should fail with dimensionPixel less than 1")

	// Test with valid parameters
	err = resizer.ResizeImageToOneDimension(img, 1, vips.KernelNearest)
	assert.NoError(t, err, "ResizeImageToOneDimension failed with valid parameters")
	assert.Equal(t, 1, img.Width(), "Resized image width does not match the expected width")
	assert.Equal(t, 1, img.Height(), "Resized image height does not match the expected height")

	// Create a simple 2x2 pixel image
	img, err = vips.Black(2, 2)
	assert.NoError(t, err, "[TEST SPEC FAIL] Failed to create image.")

	// Test with dimensionPixel greater than original dimensions
	err = resizer.ResizeImageToOneDimension(img, 3, vips.KernelNearest)
	assert.NoError(t, err, "ResizeImageToOneDimension failed with dimensionPixel greater than original dimensions")
	assert.Equal(t, 2, img.Width(), "Resized image width does not match the original width")
	assert.Equal(t, 2, img.Width(), "Resized image height does not match the original height")
}
