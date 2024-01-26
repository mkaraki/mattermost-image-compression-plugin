package main

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/webp"
)

func TestResizeImageOrPassThrough(t *testing.T) {
	plugin := &Plugin{}
	plugin.API = &MattermostMockAPI{}
	config := &configuration{
		DoImageResize:              true,
		ImageInterpolationFunction: "nearest",
		ImageMaxDimension:          100,
	}
	plugin.setConfiguration(config)

	// Create a simple 2x2 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	t.Run("Test with DoImageResize enabled but passthrough size", func(t *testing.T) {
		expectedImg := img
		actualImg := plugin.ResizeImageOrPassThrough(img)
		assert.Equal(t, expectedImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize enabled and passthrough size")
		assert.Equal(t, expectedImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize enabled and passthrough size")
	})

	t.Run("Test with DoImageResize enabled and over size", func(t *testing.T) {
		img = image.NewRGBA(image.Rect(0, 0, 200, 200))
		expectedImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
		actualImg := plugin.ResizeImageOrPassThrough(img)

		assert.Equal(t, expectedImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize enabled and over size")
		assert.Equal(t, expectedImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize enabled and over size")
	})

	t.Run("Test with DoImageResize enabled and invalid config size", func(t *testing.T) {
		plugin.configuration.ImageMaxDimension = -1
		img := image.NewRGBA(image.Rect(0, 0, 2000, 2000))
		expectedImg := image.NewRGBA(image.Rect(0, 0, 1280, 1280))
		actualImg := plugin.ResizeImageOrPassThrough(img)

		assert.Equal(t, expectedImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize enabled and invalid config size")
		assert.Equal(t, expectedImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize enabled and invalid config size")

		plugin.configuration.ImageMaxDimension = 100
	})

	t.Run("Test with DoImageResize disabled", func(t *testing.T) {
		plugin.configuration.DoImageResize = false
		expectedImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
		actualImg := plugin.ResizeImageOrPassThrough(img)

		assert.Equal(t, expectedImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize disabled")
		assert.Equal(t, expectedImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize disabled")

		plugin.configuration.DoImageResize = true
	})

	t.Run("Test image aspect ratio is preserved (width is greater)", func(t *testing.T) {
		expectImg := image.NewRGBA(image.Rect(0, 0, 100, 50))
		actualImg := plugin.ResizeImageOrPassThrough(image.NewRGBA(image.Rect(0, 0, 200, 100)))

		assert.Equal(t, expectImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize enabled and over size")
		assert.Equal(t, expectImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize enabled and over size")
	})

	t.Run("Test image aspect ratio is preserved (height is greater)", func(t *testing.T) {
		expectImg := image.NewRGBA(image.Rect(0, 0, 50, 100))
		actualImg := plugin.ResizeImageOrPassThrough(image.NewRGBA(image.Rect(0, 0, 100, 200)))

		assert.Equal(t, expectImg.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with DoImageResize enabled and over size")
		assert.Equal(t, expectImg.Bounds().Dy(), actualImg.Bounds().Dy(), "Expected image height does not match the actual image height with DoImageResize enabled and over size")
	})
}

func TestExportImageAsWebp(t *testing.T) {
	plugin := &Plugin{}
	plugin.API = &MattermostMockAPI{}
	config := &configuration{
		ImageQuality: 10,
	}
	plugin.setConfiguration(config)

	// Create a simple 2x2 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	info := &model.FileInfo{}

	t.Run("Test with valid parameters", func(t *testing.T) {
		var buf bytes.Buffer
		info := &model.FileInfo{
			Name:      "test.ext",
			Extension: "ext",
			MimeType:  "application/octet-stream",
			Size:      0,
		}

		actualInfo, err := plugin.ExportImageAsWebp(info, img, &buf)

		assert.NoError(t, err, "ExportImageAsWebp failed with valid parameters")

		assert.Equal(t, "webp", actualInfo.Extension, "Expected extension does not match the actual extension with valid parameters")
		assert.Equal(t, "test.ext.webp", actualInfo.Name, "Expected name does not match the actual name with valid parameters")
		assert.Equal(t, "image/webp", actualInfo.MimeType, "Expected mime type does not match the actual mime type with valid parameters")

		assert.Equal(t, int64(buf.Len()), actualInfo.Size, "Indicated size should match the buf length")

		_, err = webp.Decode(&buf)
		assert.NoError(t, err, "webp.Decode should not fail with ExportImageAsWebp output")
	})

	t.Run("Test with invalid image quality", func(t *testing.T) {
		plugin.configuration.ImageQuality = -1
		var buf bytes.Buffer
		info := &model.FileInfo{
			Name:      "test.ext",
			Extension: "ext",
			MimeType:  "application/octet-stream",
			Size:      0,
		}

		actualInfo, err := plugin.ExportImageAsWebp(info, img, &buf)

		assert.NoError(t, err, "ExportImageAsWebp failed with invalid image quality")

		assert.Equal(t, "webp", actualInfo.Extension, "Expected extension does not match the actual extension with valid parameters")
		assert.Equal(t, "test.ext.webp", actualInfo.Name, "Expected name does not match the actual name with valid parameters")
		assert.Equal(t, "image/webp", actualInfo.MimeType, "Expected mime type does not match the actual mime type with valid parameters")

		assert.Equal(t, int64(buf.Len()), actualInfo.Size, "Indicated size should match the buf length")

		_, err = webp.Decode(&buf)
		assert.NoError(t, err, "webp.Decode should not fail with ExportImageAsWebp output")

		plugin.configuration.ImageQuality = 10
	})

	t.Run("Test with nil image", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := plugin.ExportImageAsWebp(info, nil, &buf)

		assert.Error(t, err, "ExportImageAsWebp should fail with nil image")
	})

	t.Run("Test with nil FileInfo", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := plugin.ExportImageAsWebp(nil, img, &buf)

		assert.Error(t, err, "ExportImageAsWebp should fail with nil FileInfo")
	})

	t.Run("Test with nil Writer", func(t *testing.T) {
		_, err := plugin.ExportImageAsWebp(info, img, nil)

		assert.Error(t, err, "ExportImageAsWebp should fail with nil Writer")
	})
}

func TestExportImageAsJpeg(t *testing.T) {
	plugin := &Plugin{}
	plugin.API = &MattermostMockAPI{}
	config := &configuration{
		ImageQuality: 10,
	}
	plugin.setConfiguration(config)

	// Create a simple 2x2 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	info := &model.FileInfo{}

	t.Run("Test with valid parameters", func(t *testing.T) {
		var buf bytes.Buffer
		info := &model.FileInfo{
			Name:      "test.ext",
			Extension: "ext",
			MimeType:  "application/octet-stream",
			Size:      0,
		}

		actualInfo, err := plugin.ExportImageAsJpeg(info, img, &buf)

		assert.NoError(t, err, "ExportImageAsJpeg failed with valid parameters")

		assert.Equal(t, "jpg", actualInfo.Extension, "Expected extension does not match the actual extension with valid parameters")
		assert.Equal(t, "test.ext.jpg", actualInfo.Name, "Expected name does not match the actual name with valid parameters")
		assert.Equal(t, "image/jpeg", actualInfo.MimeType, "Expected mime type does not match the actual mime type with valid parameters")

		assert.Equal(t, int64(buf.Len()), actualInfo.Size, "Indicated size should match the buf length")

		_, err = jpeg.Decode(&buf)
		assert.NoError(t, err, "jpeg.Decode should not fail with ExportImageAsJpeg output")
	})

	t.Run("Test with invalid image quality", func(t *testing.T) {
		plugin.configuration.ImageQuality = -1
		var buf bytes.Buffer
		info := &model.FileInfo{
			Name:      "test.ext",
			Extension: "ext",
			MimeType:  "application/octet-stream",
			Size:      0,
		}

		actualInfo, err := plugin.ExportImageAsJpeg(info, img, &buf)

		assert.NoError(t, err, "ExportImageAsJpeg failed with invalid image quality")

		assert.Equal(t, "jpg", actualInfo.Extension, "Expected extension does not match the actual extension with valid parameters")
		assert.Equal(t, "test.ext.jpg", actualInfo.Name, "Expected name does not match the actual name with valid parameters")
		assert.Equal(t, "image/jpeg", actualInfo.MimeType, "Expected mime type does not match the actual mime type with valid parameters")

		assert.Equal(t, int64(buf.Len()), actualInfo.Size, "Indicated size should match the buf length")

		_, err = jpeg.Decode(&buf)
		assert.NoError(t, err, "jpeg.Decode should not fail with ExportImageAsJpeg output")

		plugin.configuration.ImageQuality = 10
	})

	t.Run("Test with nil image", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := plugin.ExportImageAsJpeg(info, nil, &buf)

		assert.Error(t, err, "ExportImageAsJpeg should fail with nil image")
	})

	t.Run("Test with nil FileInfo", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := plugin.ExportImageAsJpeg(nil, img, &buf)

		assert.Error(t, err, "ExportImageAsJpeg should fail with nil FileInfo")
	})

	t.Run("Test with nil Writer", func(t *testing.T) {
		_, err := plugin.ExportImageAsJpeg(info, img, nil)

		assert.Error(t, err, "ExportImageAsJpeg should fail with nil Writer")
	})
}

func TestReadImage(t *testing.T) {
	plugin := &Plugin{}
	plugin.API = &MattermostMockAPI{}

	t.Run("Test with valid image", func(t *testing.T) {
		img := image.NewRGBA(image.Rect(0, 0, 2, 2))

		t.Run("Test with valid JPEG", func(t *testing.T) {
			var buf bytes.Buffer
			err := jpeg.Encode(&buf, img, nil)
			assert.NoError(t, err, "[TEST SPEC FAIL] Failed to encode image")

			reader := bytes.NewReader(buf.Bytes())
			actualImg, err := plugin.ReadImage(reader)

			assert.NoError(t, err, "ReadImage failed with valid image")
			assert.NotNil(t, actualImg, "ReadImage should return a non-nil image with valid image")

			assert.Equal(t, img.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with valid JPEG")
		})

		t.Run("Test with valid PNG", func(t *testing.T) {
			var buf bytes.Buffer
			err := png.Encode(&buf, img)
			assert.NoError(t, err, "[TEST SPEC FAIL] Failed to encode image")

			reader := bytes.NewReader(buf.Bytes())
			actualImg, err := plugin.ReadImage(reader)

			assert.NoError(t, err, "ReadImage failed with valid image")
			assert.NotNil(t, actualImg, "ReadImage should return a non-nil image with valid image")

			assert.Equal(t, img.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with valid PNG")
		})

		t.Run("Test with valid GIF", func(t *testing.T) {
			var buf bytes.Buffer
			err := gif.Encode(&buf, img, nil)
			assert.NoError(t, err, "[TEST SPEC FAIL] Failed to encode image")

			reader := bytes.NewReader(buf.Bytes())
			actualImg, err := plugin.ReadImage(reader)

			assert.NoError(t, err, "ReadImage failed with valid image")
			assert.NotNil(t, actualImg, "ReadImage should return a non-nil image with valid image")

			assert.Equal(t, img.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with valid GIF")
		})

		t.Run("Test with valid WEBP", func(t *testing.T) {
			var buf bytes.Buffer
			imageEncoder := &ImageEncoder{}
			_, err := imageEncoder.EncodeWebP(img, &buf, 90)
			assert.NoError(t, err, "[TEST SPEC FAIL] Failed to encode image")

			reader := bytes.NewReader(buf.Bytes())
			actualImg, err := plugin.ReadImage(reader)

			assert.NoError(t, err, "ReadImage failed with valid image")
			assert.NotNil(t, actualImg, "ReadImage should return a non-nil image with valid image")

			assert.Equal(t, img.Bounds().Dx(), actualImg.Bounds().Dx(), "Expected image width does not match the actual image width with valid WEBP")
		})
	})

	t.Run("Test with invalid image", func(t *testing.T) {
		reader := bytes.NewReader([]byte("invalid image"))
		_, err := plugin.ReadImage(reader)

		assert.Error(t, err, "ReadImage should fail with invalid image")
	})

	t.Run("Test with nil reader", func(t *testing.T) {
		_, err := plugin.ReadImage(nil)

		assert.Error(t, err, "ReadImage should fail with nil reader")
	})
}
