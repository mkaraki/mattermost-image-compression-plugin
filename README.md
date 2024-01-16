# Mattermost Image Compression Plugin

[![Build Status](https://github.com/mkaraki/mattermost-image-compression-plugin/actions/workflows/build.yml/badge.svg)](https://github.com/mkaraki/mattermost-image-compression-plugin/actions/workflows/build.yml)

Resize uploaded image and compress them with [WebP](https://developers.google.com/speed/webp).

## Configuration

You can configure following settings.

|Settings|Default|Description|
|---|---|---|
|`OutputImageFormat`|`webp`|Compress type|
|`ImageQuality`|`10`|Image encoding quality|
|`DoImageResize`|`true`|Do image resize for decrease file size|
|`ImageMaxDimension`|`1280`|Max image dimension|
|`ImageInterpolationFunction`|`Lanczos (a=3)`|See [here](https://github.com/OneOfOne/resize/blob/master/README.md#usage) for more information|
