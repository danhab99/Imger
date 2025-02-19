package histogram

import (
	"github.com/danhab99/Imger/utils"
	"image"
	"image/color"
)

const hsize = 256
const channels = 3

// HistogramGray computes the histogram for a grayscale image. Returns an array of 256 uint64 values containing
// distribution of the pixel values.
func HistogramGray(img *image.Gray) [hsize]uint64 {
	var res [hsize]uint64
	utils.ForEachGrayPixel(img, func(pixel color.Gray) {
		res[pixel.Y]++
	})
	return res
}

// HistogramRGBARed computes the histogram for red channel from an RGBA image. Returns an array of 256 uint64 values
// containing distribution of the pixel values.
func HistogramRGBARed(img *image.RGBA) [hsize]uint64 {
	var res [hsize]uint64
	utils.ForEachRGBARedPixel(img, func(red uint8) {
		res[red]++
	})
	return res
}

// HistogramRGBAGreen computes the histogram for green channel from an RGBA image. Returns an array of 256 uint64 values
// containing distribution of the pixel values.
func HistogramRGBAGreen(img *image.RGBA) [hsize]uint64 {
	var res [hsize]uint64
	utils.ForEachRGBAGreenPixel(img, func(green uint8) {
		res[green]++
	})
	return res
}

// HistogramRGBABlue computes the histogram for blue channel from an RGBA image. Returns an array of 256 uint64 values
// containing distribution of the pixel values.
func HistogramRGBABlue(img *image.RGBA) [hsize]uint64 {
	var res [hsize]uint64
	utils.ForEachRGBABluePixel(img, func(blue uint8) {
		res[blue]++
	})
	return res
}

// HistogramRGBA computes the histogram for a RGBA image. Returns an 2D (shape: [3][256]) array of uint64 values
// containing distribution of color values from each RGBA channel.
func HistogramRGBA(img *image.RGBA) [channels][hsize]uint64 {
	var res [channels][hsize]uint64
	utils.ForEachRGBAPixel(img, func(pixel color.RGBA) {
		res[0][pixel.R]++
		res[1][pixel.G]++
		res[2][pixel.B]++
	})
	return res
}

// DrawHistogramGray computes and draws the histogram of a grayscale image. The size of the image is 256*scale width
// and 256*scale height.
func DrawHistogramGray(img *image.Gray, size image.Point) *image.Gray {
	h := HistogramGray(img)
	normHist := normalizeHistogram(h, uint64(size.Y))
	res := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	drawerFunc(size, func(i int) uint64 {
		return normHist[i]
	}, func(x, y int) {
		res.SetGray(x, y, color.Gray{Y: utils.MaxUint8})
	})
	return res
}

// DrawHistogramRGBA computes and draws the histogram of a RGBA image. The size of the image is 256*scale width and
// 256*scale height.
func DrawHistogramRGBA(img *image.RGBA, size image.Point) *image.RGBA {
	h := HistogramRGBA(img)
	normRHist := normalizeHistogram(h[0], uint64(size.Y))
	normGHist := normalizeHistogram(h[1], uint64(size.Y))
	normBHist := normalizeHistogram(h[2], uint64(size.Y))
	res := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	drawerFunc(size, func(i int) uint64 {
		return normRHist[i]
	}, func(x, y int) {
		pix := res.RGBAAt(x, y)
		pix.R = utils.MaxUint8
		res.SetRGBA(x, y, pix)
	})
	drawerFunc(size, func(i int) uint64 {
		return normGHist[i]
	}, func(x, y int) {
		pix := res.RGBAAt(x, y)
		pix.G = utils.MaxUint8
		res.SetRGBA(x, y, pix)
	})
	drawerFunc(size, func(i int) uint64 {
		return normBHist[i]
	}, func(x, y int) {
		pix := res.RGBAAt(x, y)
		pix.B = utils.MaxUint8
		res.SetRGBA(x, y, pix)
	})
	return res
}

//---------------------------------------------------------------------------------------------
func drawerFunc(size image.Point, getNormAt func(i int) uint64, setPixel func(x, y int)) {
	scaleX := float64(size.X) / float64(hsize)
	for i := 0; i < hsize; i++ {
		for width := int(float64(i) * scaleX); width < int((float64(i)+1.0)*scaleX); width++ {
			for height := size.Y; height >= size.Y-int(getNormAt(i)); height-- {
				setPixel(width, height)
			}
		}
	}
}

func normalizeHistogram(v [hsize]uint64, maxHeight uint64) [hsize]uint64 {
	max := utils.GetMax(v[:])
	var norm [hsize]uint64
	for i := 0; i < len(v); i++ {
		norm[i] = v[i] * maxHeight / max
	}
	return norm
}
