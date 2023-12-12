# 🖼️ Dummy Image

Dummy Image allows you to dynamically create and return images of various sizes, colors and styles, providing flexible customization of image generation through a simple and user-friendly URL interface.

> still in development

## 📃 URL

/ `size` / `background_color` / `text_color` / `format` & `text`

### Formats

- Size: `100x100` `100` `hd` ...
- Color: `fff` `ffffff` `255,255,255`
- Image: `png` `jpeg` `svg`

### Limits

- Min size: `10x10`
- Max size: `2000x2000`

Examples

```
http://127.0.0.1:8080/800x800/fff/000/png&text="Hello world"
```

> size, background_color, text_color, format, text

```
http://127.0.0.1:8080/hd/f1e2d3/png&text="Hello world"
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
