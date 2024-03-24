package main

import (
	"fmt"
	"image"
	"io"
	"strings"
	"sync"

	"github.com/adrium/goheif"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"

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

func (p *Plugin) ResizeImageOrPassThrough(original_image image.Image) image.Image {
	if p.configuration.DoImageResize {
		interpolationFunctionName := p.configuration.ImageInterpolationFunction

		resizer := &ImageResizer{}

		image_interpolation_function := resizer.GetInterpolationFunctionByName(interpolationFunctionName)

		new_image_max_dimension := p.configuration.ImageMaxDimension
		if new_image_max_dimension < 1 {
			p.API.LogWarn("Invalid image max dimension value, use 1280.")
			new_image_max_dimension = 1280
		}

		if (original_image.Bounds().Dx() <= new_image_max_dimension) && (original_image.Bounds().Dy() <= new_image_max_dimension) {
			p.API.LogDebug("Image is smaller than max dimension, passing through.")
			return original_image
		}

		new_img, err := resizer.ResizeImageToOneDimension(original_image, new_image_max_dimension, image_interpolation_function)
		if err != nil {
			p.API.LogError("Failed to resize image", "error", err.Error())
			return original_image
		}

		p.API.LogDebug("Successfully resized image", "new_width", new_img.Bounds().Dx(), "new_height", new_img.Bounds().Dy())

		return new_img
	} else {
		p.API.LogDebug("Image resize is disabled, passing through.")
		return original_image
	}
}

func (p *Plugin) ExportImageAsWebp(info *model.FileInfo, image image.Image, output io.Writer) (*model.FileInfo, error) {
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

func (p *Plugin) ExportImageAsJpeg(info *model.FileInfo, image image.Image, output io.Writer) (*model.FileInfo, error) {
	if (info == nil) || (image == nil) || (output == nil) {
		p.API.LogError("Invalid arguments to ExportImageAsJpeg. THIS SHOULD NEVER HAPPEN.")
		return nil, fmt.Errorf("Invalid arguments to ExportImageAsJpeg. THIS SHOULD NEVER HAPPEN.")
	}

	image_quality := p.configuration.ImageQuality

	if (image_quality < 0) || (image_quality > 100) {
		p.API.LogWarn("Invalid image quality value, use 10.")
		image_quality = 10
	}

	ImageEncoder := &ImageEncoder{}

	output_size, err := ImageEncoder.EncodeJpeg(image, output, image_quality)
	if err != nil {
		p.API.LogError("Failed to encode image as JPEG", "error", err.Error())
		return nil, err
	}

	p.API.LogDebug("Successfully encoded image as JPEG", "size", output_size)

	info.Extension = "jpg"
	info.Name = info.Name + ".jpg"
	info.MimeType = "image/jpeg"
	info.Size = int64(output_size)

	return info, nil
}

func (p *Plugin) ReadImage(file io.Reader) (image.Image, error) {
	if file == nil {
		p.API.LogError("Invalid arguments to ReadImage. THIS SHOULD NEVER HAPPEN.")
		return nil, fmt.Errorf("Invalid arguments to ReadImage. THIS SHOULD NEVER HAPPEN.")
	}

	image, image_type, err := image.Decode(file)
	if err != nil {
		p.API.LogError("Failed to decode image", "error", err.Error())
		return nil, err
	}
	p.API.LogDebug("Input image type", "type", image_type)
	return image, nil
}

func (p *Plugin) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {

	var original_image image.Image
	var err error

	switch strings.ToLower(info.Extension) {
	case "jpg", "jpeg", "png", "gif", "webp":
		p.API.LogDebug("Processing image via Image API", "name", info.Name, "extension", info.Extension)
		original_image, err = p.ReadImage(file)
		if err != nil {
			return info, ""
		}
	case "heif", "heic":
		p.API.LogDebug("Processing image via goHeif", "name", info.Name, "extension", info.Extension)
		original_image, err = goheif.Decode(file)
		if err != nil {
			return info, ""
		}
	default:
		return info, ""
	}

	resized_image := p.ResizeImageOrPassThrough(original_image)

	output_format := p.configuration.OutputImageFormat

	var new_file_info *model.FileInfo

	p.API.LogDebug("Trying to export image", "format", output_format, "original_size", info.Size)

	switch output_format {
	case "webp":
		new_file_info, err = p.ExportImageAsWebp(info, resized_image, output)
	case "jpeg":
		new_file_info, err = p.ExportImageAsJpeg(info, resized_image, output)
	default:
		p.API.LogWarn("Output format is not valid value.")
		return info, ""
	}

	new_file_info.Width = resized_image.Bounds().Dx()
	new_file_info.Height = resized_image.Bounds().Dy()

	if err != nil {
		return nil, "Unable to export compressed image. Contact your system administrator."
	}

	return new_file_info, ""
}
