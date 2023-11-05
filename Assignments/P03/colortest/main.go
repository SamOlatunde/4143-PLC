package main

import (
	"github.com/SamOlatunde/img_mod/Colors"
	"github.com/SamOlatunde/img_mod/GetPic"
	"github.com/SamOlatunde/img_mod/Grayscale"
	"github.com/SamOlatunde/img_mod/Text"
)

func main() {
	GetPic.GetPic()
	Colors.Colors("downloaded_image.jpg")
	Grayscale.Grayscale("downloaded_image.jpg")
	Text.Text()
}
