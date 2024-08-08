package main

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"

	"github.com/davidbyttow/govips/v2/vips"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

func (p *Plugin) ResizeImageOrPassThrough(original_image *vips.ImageRef) {
	if p.configuration.DoImageResize {
		new_image_max_dimension := p.configuration.ImageMaxDimension

		if new_image_max_dimension < 1 {
			p.API.LogWarn("Invalid image max dimension value, use 1280.")
			new_image_max_dimension = 1280
		}

		img_width := original_image.Width()
		img_height := original_image.Height()

		if img_width <= new_image_max_dimension && img_height <= new_image_max_dimension {
			p.API.LogDebug("Image is smaller than max dimension, passing through.")
			return
		}

		resizer := &ImageResizer{}

		interpolationFunctionName := p.configuration.ImageInterpolationFunction
		image_interpolation_function := resizer.GetInterpolationFunctionByName(interpolationFunctionName)

		err := resizer.ResizeImageToOneDimension(original_image, new_image_max_dimension, image_interpolation_function)
		if err != nil {
			p.API.LogError("Failed to resize image", "error", err.Error())
			return
		}

		img_width = original_image.Width()
		img_height = original_image.Height()

		p.API.LogDebug("Successfully resized image", "new_width", img_width, "new_height", img_height)

		return
	} else {
		p.API.LogDebug("Image resize is disabled, passing through.")
		return
	}
}

func (p *Plugin) ExportImageAsWebp(info *model.FileInfo, image *vips.ImageRef, output io.Writer) (*model.FileInfo, error) {
	if (info == nil) || (image == nil) || (output == nil) {
		p.API.LogError("Invalid arguments to ExportImageAsWebp. THIS SHOULD NEVER HAPPEN.")
		return nil, fmt.Errorf("Invalid arguments to ExportImageAsWebp. THIS SHOULD NEVER HAPPEN.")
	}

	ImageEncoder := &ImageEncoder{}

	image_quality := p.configuration.ImageQuality
	if (image_quality < 0) || (image_quality > 100) {
		p.API.LogWarn("Invalid image quality value, use 10.")
		image_quality = 10
	}

	output_size, err := ImageEncoder.EncodeWebP(image, output, image_quality)
	if err != nil {
		p.API.LogError("Failed to encode image as webp", "error", err.Error())
		return nil, err
	}

	p.API.LogDebug("Successfully encoded image as webp", "size", output_size)

	info.Extension = "webp"
	info.Name = info.Name + ".webp"
	info.MimeType = "image/webp"
	info.Size = int64(output_size)

	return info, nil
}

func (p *Plugin) ReadImage(file io.Reader) (*vips.ImageRef, error) {
	if file == nil {
		return nil, fmt.Errorf("Input reader is nil.")
	}

	img, err := vips.NewImageFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to read image.")
	}

	return img, nil
}

func (p *Plugin) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {

	var img *vips.ImageRef
	var err error

	switch strings.ToLower(info.Extension) {
	case "jpg", "jpeg", "png", "gif", "webp", "heif", "heic":
		p.API.LogDebug("Processing image via Image API", "name", info.Name, "extension", info.Extension)
		img, err = p.ReadImage(file)
		if err != nil {
			return info, ""
		}
	default:
		// When file extension is not to process.
		return info, ""
	}

	p.ResizeImageOrPassThrough(img)

	output_format := p.configuration.OutputImageFormat

	var new_file_info *model.FileInfo

	p.API.LogDebug("Trying to export image", "format", output_format, "original_size", info.Size)

	switch output_format {
	case "webp":
		new_file_info, err = p.ExportImageAsWebp(info, img, output)
	default:
		p.API.LogWarn("Output format is not valid value.")
		new_file_info, err = p.ExportImageAsWebp(info, img, output)
	}

	new_file_info.Width = img.Width()
	new_file_info.Height = img.Height()

	if err != nil {
		return nil, "Unable to export compressed image. Contact your system administrator."
	}

	return new_file_info, ""
}
