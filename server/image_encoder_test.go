package main

import (
	"bytes"
	"github.com/davidbyttow/govips/v2/vips"
	"testing"

	"golang.org/x/image/webp"

	"github.com/stretchr/testify/assert"
)

func TestEncodeWebP(t *testing.T) {
	// Create a new ImageEncoder
	encoder := &ImageEncoder{}

	// Create a simple 1x1 pixel image
	img, err := vips.Black(1, 1)
	assert.NoError(t, err, "[TEST SPEC FAIL] Failed to create image.")

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
