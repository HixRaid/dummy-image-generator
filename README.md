# 🖼️ Dummy Image

Dummy Image allows you to dynamically create and return images of various sizes, colors and styles, providing flexible customization of image generation through a simple and user-friendly URL interface.

> Still in development

## 📃 URL

/ `size` / `background_color` / `text_color` / `format` & `text`

> Size required

### Formats

- Size: `100x100` `100` `HD` [and other...](RESOLUTIONS.md)
- Color: `FFF` `FFFFFF` `255,255,255`
- Image: `PNG` `JPEG` `SVG`

> Default PNG

### Limits

- Min size: `10x10`
- Max size: `2000x2000`

Examples

```
http://127.0.0.1:8080/800x800/FFF/000/png&text="Hello world"
```

> size, background_color, text_color, format, text

```
http://127.0.0.1:8080/HD/F1E2D3/png&text="Hello world"
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
