package github.com/eclypse13/mandelbrot

import (
	"math"

	"github.com/fogleman/gg"
)

// Config contains all the parameters for generating the Mandelbrot set.
type Config struct {
	Width, Height                      int     // Image size in pixels
	MinReal, MaxReal, MinImag, MaxImag float64 // Boundaries of the domain of the complex plane
	MaxIter                            int     // Maximum number of iterations for eact point
}

// GenerateMandelbrot generates an image of a Mandelbrot set and saves it to the file "mandelbrot.png".
func GenerateMandelbrot() {
	cfg := Config{
		Width:   9600,
		Height:  5400,
		MinReal: -2.0,
		MaxReal: 1.0,
		MinImag: -1.0,
		MaxImag: 1.0,
		MaxIter: 1000,
	}

	dc := gg.NewContext(cfg.Width, cfg.Height)

	for py := 0; py < cfg.Height; py++ {
		for px := 0; px < cfg.Width; px++ {
			// Converting pixel coordinates into a complex number c
			cReal := cfg.MinReal + (cfg.MaxReal-cfg.MinReal)*float64(px)/float64(cfg.Width-1)
			cImag := cfg.MinImag + (cfg.MaxImag-cfg.MinImag)*float64(py)/float64(cfg.Height-1)

			// Iteratively calculate z = z * z + c, starting from z = 0
			zReal, zImag := 0.0, 0.0
			iter := 0

			for iter < cfg.MaxIter {
				newZReal := zReal*zReal - zImag*zImag + cReal
				newZImag := 2*zReal*zImag + cImag

				// If |z| > 2 then z does not belong to the Mandelbrot set and can be discarded
				if newZReal*newZReal+newZImag*newZImag > 4 {
					break
				}

				zReal, zImag = newZReal, newZImag
				iter++
			}

			// Selecting a pixel color based on the number of variations
			var r, g, b float64

			if iter == cfg.MaxIter {
				// The point belongs to Mandelbrot set - paint it black
				r, g, b = 0, 0, 0
			} else {
				// Gradient: the faster a point leaves a Mandelbrot set, the brigther the color
				t := float64(iter) / float64(cfg.MaxIter)

				r = 0.5 + 0.5*math.Sin(2*math.Pi*(t+0.00))
				g = 0.5 + 0.5*math.Sin(2*math.Pi*(t+0.33))
				b = 0.5 + 0.5*math.Sin(2*math.Pi*(t+0.67))
			}

			// Fill the current pixel
			dc.SetRGB(r, g, b)
			dc.SetPixel(px, py)
		}
	}
	dc.SavePNG("mandelbrot.png")
}
