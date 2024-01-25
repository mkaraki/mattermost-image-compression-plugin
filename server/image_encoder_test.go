package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"testing"

	"golang.org/x/image/webp"

	"github.com/stretchr/testify/assert"
)

func TestEncodeWebP(t *testing.T) {
	// Create a new ImageEncoder
	encoder := &ImageEncoder{}

	// Create a simple 1x1 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})

	// Create a buffer to hold the output
	buf := &bytes.Buffer{}

	// Test with valid quality
	indicated_size, err := encoder.EncodeWebP(img, buf, 90)
	assert.NoError(t, err, "EncodeWebP should not fail with valid quality (quality=90)")

	assert.Equal(t, int64(buf.Len()), indicated_size, "Indicated size should match the buf length")

	_, err = webp.Decode(buf)
	assert.NoError(t, err, "webp.Decode should not fail with EncodeWebP output")

	// Test with quality over 100
	_, err = encoder.EncodeWebP(img, buf, 110)
	assert.NoError(t, err, "EncodeWebP should not fail with quality over 100")

	// Test with quality under 0
	_, err = encoder.EncodeWebP(img, buf, -10)
	assert.NoError(t, err, "EncodeWebP should not fail with quality under 0")

	// Test with nil image
	_, err = encoder.EncodeWebP(nil, buf, 90)
	assert.Error(t, err, "EncodeWebP should fail with nil image")

	// Test with nil output
	_, err = encoder.EncodeWebP(img, nil, 90)
	assert.Error(t, err, "EncodeWebP should fail when output is nil")
}

func TestEncodeJpeg(t *testing.T) {
	// Create a new ImageEncoder
	encoder := &ImageEncoder{}

	// Create a simple 1x1 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})

	// Create a buffer to hold the output
	buf := new(bytes.Buffer)

	// Test with valid quality
	indicated_size, err := encoder.EncodeJpeg(img, buf, 90)
	assert.NoError(t, err, "EncodeJpeg should not fail with valid quality (quality=90)")

	assert.Equal(t, int64(buf.Len()), indicated_size, "Indicated size should match the buf length")

	_, err = jpeg.Decode(buf)
	assert.NoError(t, err, "jpeg.Decode should not fail with EncodeJpeg output")

	// Test with quality over 100
	_, err = encoder.EncodeJpeg(img, buf, 110)
	assert.NoError(t, err, "EncodeJpeg should not fail with quality over 100")

	// Test with quality under 0
	_, err = encoder.EncodeJpeg(img, buf, -10)
	assert.NoError(t, err, "EncodeJpeg should not fail with quality under 0")

	// Test with nil image
	_, err = encoder.EncodeJpeg(nil, buf, 90)
	assert.Error(t, err, "EncodeJpeg should fail with nil image")

	// Test with nil output
	_, err = encoder.EncodeJpeg(img, nil, 90)
	assert.Error(t, err, "EncodeJpeg should fail when output is nil")
}
