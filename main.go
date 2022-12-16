package main

const WIDTH = 800
const HEIGHT = 600

const COLS = 8 * 2
const ROWS = 6 * 2
const CellWidth = WIDTH / COLS
const CellHeight = HEIGHT / ROWS

const BackgroundColor = 0xFF202020

//goland:noinspection GoUnusedConst
const ForegroundColor = 0xFF2020FF

func swapInt(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func drawLine(pixels *[]uint32, pixelsWidth, pixelsHeight sizeT, x1, y1, x2, y2 int, color uint32) {
	dx := x2 - x1
	dy := y2 - y1

	if dx != 0 {
		c := y1 - dy*x1/dx

		if x1 > x2 {
			swapInt(&x1, &x2)
		}

		for x := x1; x <= x2; x++ {
			if 0 <= x && x < int(pixelsWidth) {
				sy1 := dy*x/dx + c
				sy2 := dy*(x+1)/dx + c
				if sy1 > sy2 {
					swapInt(&sy1, &sy2)
				}

				for y := sy1; y <= sy2; y++ {
					if 0 <= y && y < int(pixelsHeight) {
						(*pixels)[y*int(pixelsWidth)+x] = color
					}
				}
			}
		}
	} else {
		x := x1
		if 0 <= x && x < int(pixelsWidth) {
			if y1 > y2 {
				swapInt(&y1, &y2)
			}
			for y := y1; y <= y2; y++ {
				if 0 <= y && y < int(pixelsHeight) {
					(*pixels)[y*int(pixelsWidth)+x] = color
				}
			}
		}
	}
}

func checkerExample(pixels *[WIDTH * HEIGHT]uint32) bool {
	fill(pixels, WIDTH, HEIGHT, BackgroundColor)

	for y := 0; y < ROWS; y++ {
		for x := 0; x < COLS; x++ {
			color := uint32(BackgroundColor)
			if (x+y)%2 == 0 {
				color = 0xFF2020FF
			}
			fillRect(pixels, WIDTH, HEIGHT, x*CellWidth, y*CellHeight, CellWidth, CellHeight, color)
		}
	}

	filePath := "checker.ppm"
	err := saveToPpmFile(pixels, WIDTH, HEIGHT, filePath)
	if err != 0 {
		println("Could not save to file: " + filePath)
		return false
	}

	return true
}

func main() {
	pixels := [WIDTH * HEIGHT]uint32{}

	for i := 0; i < len(pixels); i++ {
		pixels[i] = 0
	}

	checkerExample(&pixels)
}
