package image

import (
	"errors"
	"os"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"strconv"
	"runtime"
	"math"
	"github.com/disintegration/imaging"
	"github.com/RomaBiliak/go_api_admin_blog/library/file"
)
var originalImage image.Image
var watermarkImage image.Image

func AddWatermark(path string) (string, error){
	runtime.GOMAXPROCS(runtime.NumCPU())
	original, err := os.Open(path)
	defer original.Close()
	if err != nil {
		return path, err
	}
	
	config, formatOriginal, err := image.DecodeConfig(original)
	x := config.Width
	y := config.Height
	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	original.Seek(0, 0)
	if err != nil {
		return path, err
	}

	if formatOriginal == "jpeg"{
		originalImage, err = jpeg.Decode(original)
	}else if formatOriginal == "png"{
		originalImage, err = png.Decode(original)
	}else{
		return path, errors.New("incorrect file format")
	}
	if err != nil {
		return path, err
	}
	
	watermark, err := os.Open(os.Getenv("WATERMARK"))
	defer watermark.Close()
	if err != nil {
		return os.Getenv("WATERMARK"), err
	}

	config, _, err = image.DecodeConfig(watermark)
	x = x - config.Width
	y = y - config.Height
	watermark.Seek(0, 0)
	watermarkImage, err = png.Decode(watermark)
	if err != nil {
		return os.Getenv("WATERMARK"), err
	}
	
	// Watermark offset. left top corner in this example
	offset := image.Pt(x, y)
	// Use same size as source image has
	b := originalImage.Bounds()
	m := image.NewRGBA(b)
	// Draw source
	draw.Draw(m, b, originalImage, image.ZP, draw.Src)
	// Draw watermark
	draw.Draw(m, watermarkImage.Bounds().Add(offset), watermarkImage, image.ZP, draw.Over)

	// Save final JPG
	file.MkdirIfNotExist(os.Getenv("POST_PRODUCTION_DIR"))
	new_name := os.Getenv("POST_PRODUCTION_DIR") + "/" + file.GetName(path) + ".jpg"
	out, err := os.Create(new_name)
	if err != nil {
		return new_name, err
	}
	defer out.Close()
	jpeg.Encode(out, m, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})

	return new_name, err
}

func Resize(path string) (string, error){
	runtime.GOMAXPROCS(runtime.NumCPU())

	folder := os.Getenv("POST_PRODUCTION_DIR") + "/" + os.Getenv("CROP_IMAGE_WIGHT")
	file.MkdirIfNotExist(folder)

	img, err := imaging.Open(path)

	if err != nil {
		return path, err
	}

	wight, _ := strconv.ParseInt(os.Getenv("CROP_IMAGE_WIGHT"), 10, 32)
	newImage := imaging.Resize(img, int(wight), 0, imaging.Lanczos)
	new_name := folder + "/" + file.GetName(path) + ".jpg"
	err = imaging.Save(newImage, new_name)
	return new_name, err

}

func Split(path string) (map[int]string, error){
	runtime.GOMAXPROCS(runtime.NumCPU())
	var x_c_0 = 0
	var y_r_0 = 0
	var x_c_1 = 0
	var y_r_1 = 0
	var n = 0
	var split_name = ""
	paths := make(map[int]string)
	var name = file.GetName(path)

	file_open, err := os.Open(path)
	if err != nil {
		return paths, err
	}
	
	config, _, err := image.DecodeConfig(file_open)
	x := int(config.Width)
	y := int(config.Height)
	dx := int(math.Round(float64(x)/3))
	dy := int(math.Round(float64(y)/3))
	file_open.Seek(0, 0)
	if err != nil {
		return paths, err
	}

	img, err := jpeg.Decode(file_open)
	folder := os.Getenv("POST_PRODUCTION_DIR") + "/" + "split"
	file.MkdirIfNotExist(folder)
	 
	for r := 0; r < 3; r++{
		for c := 0; c < 3; c++{
			n++
			x_c_0 = c*dx
			y_r_0 = r*dy
			x_c_1 = x - (3-c-1) * dx
			y_r_1 = y - (3-r-1) * dy
			rectcropimg := imaging.Crop(img, image.Rect(x_c_0, y_r_0, x_c_1, y_r_1))
			split_name = folder + "/" + strconv.Itoa(n) + "_"  + name +".jpg"
			err = imaging.Save(rectcropimg, split_name)
			if err == nil{
				paths[n] = split_name
			}
		}
	}

	return paths, err
}
