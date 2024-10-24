package utils

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
)

func GetImageDimensions(file multipart.File) (int, int, error) {
	file.Seek(0, 0)
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	width := img.Width
	height := img.Height

	return width, height, nil
}
