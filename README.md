# 🖼️ Dummy Image Generator

<p align="center">
  <img src="static/img/logo.svg" width=300>
</p>

## What is this

Dummy Image Generator allows you to dynamically create and return images of various sizes, colors and styles, providing flexible customization of image generation through a simple and user-friendly URL interface.

> Still in development

## 🗂️ Project

### Structure

- Executable file
- Config `config.yaml`
- Fonts `arial.ttf`
- Resources `favicon.ico`
- Static `css` `img` `js`
- Views `index.html`

### Config

- server
  - addr `host` : `port`
- image
  - default_format `image format`
  - size
    - min_size `size format`
    - max_size `size format`
    - size_clamp `bool`
  - color
    - default_background_color `color format`
    - default_text_color `color format`

### Build and run

```
go mod download
go build -ldflags="-s -w" -o http_server.exe cmd\http_server\main.go
.\http_server.exe
```

> Example for Windows

## 📃 URL

/ `size` / `background_color` / `text_color` / `format` ? `text`

### Formats

- Size: `100x100` `100` `HD` [and other...](RESOLUTIONS.md)

  > Required

- Color: `F` `FF` `FFF` `FFFFFF` `255,255,255` `White` [and other...](COLORS.md)

  > Default: background is `Black` text is `White`

- Text: `Hello world` `no_text`

  > Default: `width x height`

- Image: `PNG` `JPEG` `WEBP` `GIF` `SVG`

  > Default: `PNG`

### Limits

- Min size: `10x10`
- Max size: `2000x2000`

> Automatic size clamp

Examples

```
http://127.0.0.1:8080/800x800/FFF/000/gif?text=Hello%20world
```

> size, background_color, text_color, format, text

```
http://127.0.0.1:8080/HD/F1E2D3/webp?text=Hello%20world
```

> size, background_color, format, text

```
http://127.0.0.1:8080/400
```

> size

## 📦 Libraries

| Name   | Module                                                           | Version  |
| :----- | :--------------------------------------------------------------- | :------- |
| gin    | [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)     | `1.9.1`  |
| viper  | [github.com/spf13/viper](https://github.com/spf13/viper)         | `1.17.0` |
| logrus | [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) | `1.9.3`  |
| gg     | [github.com/fogleman/gg](https://github.com/fogleman/gg)         | `1.3.0`  |
| webp   | [github.com/chai2010/webp](https://github.com/chai2010/webp)     | `1.1.1`  |
| svgo   | [github.com/ajstarks/svgo](https://github.com/ajstarks/svgo)     | `0.0.0`  |
