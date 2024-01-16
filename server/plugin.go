package main

import (
	"image"
	"io"
	"strings"
	"sync"

	"github.com/chai2010/webp"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"go.oneofone.dev/resize"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
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

func (p *Plugin) SelectImageInterpolationFunction() resize.InterpolationFunction {
	switch p.configuration.ImageInterpolationFunction {
	default:
		p.API.LogWarn("Invalid image interpolation function, use nearest.")
		fallthrough
	case "nearest":
		p.API.LogDebug("Decided to use interpolation function:", "function", "nearest")
		return resize.NearestNeighbor
	case "bilinear":
		p.API.LogDebug("Decided to use interpolation function:", "function", "bilinear")
		return resize.Bilinear
	case "bicubic":
		p.API.LogDebug("Decided to use interpolation function:", "function", "bicubic")
		return resize.Bicubic
	case "mitchell-netravali":
		p.API.LogDebug("Decided to use interpolation function:", "function", "mitchell_netravali")
		return resize.MitchellNetravali
	case "lanczos2":
		p.API.LogDebug("Decided to use interpolation function:", "function", "lanczos (a=2)")
		return resize.Lanczos2
	case "lanczos3":
		p.API.LogDebug("Decided to use interpolation function:", "function", "lanczos (a=3)")
		return resize.Lanczos3
	}
}

func (p *Plugin) ResizeImageOrPassThrough(original_image image.Image) image.Image {
	image_interpolation_function := p.SelectImageInterpolationFunction()

	new_image_max_dimension := p.configuration.ImageMaxDimension
	if new_image_max_dimension < 1 {
		p.API.LogWarn("Invalid image max dimension value, use 1280.")
		new_image_max_dimension = 1280
	}

	if p.configuration.DoImageResize {
		if (original_image.Bounds().Dx() <= new_image_max_dimension) && (original_image.Bounds().Dy() <= new_image_max_dimension) {
			p.API.LogDebug("Image is smaller than max dimension, passing through.")
			return original_image
		}

		return resize.Thumbnail(uint(p.configuration.ImageMaxDimension), uint(p.configuration.ImageMaxDimension), original_image, image_interpolation_function)
	} else {
		p.API.LogDebug("Image resize is disabled, passing through.")
		return original_image
	}
}

func (p *Plugin) ExportImageAsWebp(info *model.FileInfo, image image.Image, output io.Writer) (*model.FileInfo, error) {
	image_quality := p.configuration.ImageQuality

	if (image_quality < 0) || (image_quality > 100) {
		p.API.LogWarn("Invalid image quality value, use 10.")
		image_quality = 10
	}

	webp_data, err := webp.EncodeRGB(image, float32(image_quality))
	if err != nil {
		p.API.LogError("Failed to encode image as webp", "error", err.Error())
		return nil, err
	}

	output_size := len(webp_data)

	p.API.LogDebug("Successfully encoded image as webp", "size", output_size)

	_, err = output.Write(webp_data)
	info.Extension = "webp"
	info.Name = info.Name + ".webp"
	info.MimeType = "image/webp"
	info.Size = int64(output_size)

	return info, nil
}

func (p *Plugin) ReadImage(file io.Reader) (image.Image, error) {
	image, image_type, err := image.Decode(file)
	if err != nil {
		p.API.LogError("Failed to decode image", "error", err.Error())
		return nil, err
	}
	p.API.LogDebug("Input image type", "type", image_type)
	return image, nil
}

func (p *Plugin) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {

	switch strings.ToLower(info.Extension) {
	case "jpg", "jpeg", "png", "gif", "webp":
	default:
		return info, ""
	}

	p.API.LogDebug("Processing image", "name", info.Name, "extension", info.Extension)

	original_image, err := p.ReadImage(file)
	if err != nil {
		return info, ""
	}

	resized_image := p.ResizeImageOrPassThrough(original_image)

	new_file_info, err := p.ExportImageAsWebp(info, resized_image, output)
	if err != nil {
		return nil, "Unable to export compressed image. Contact your system administrator."
	}

	return new_file_info, ""
}
