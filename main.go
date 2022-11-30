package main

//goland:noinspection GoUnusedConst
const AARES = 2

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func sign(x int) (res int) {
	if x > 0 {
		res = 1
	} else if x < 0 {
		res = -1
	}

	return
}
func abs(x int) int {
	return sign(x) * x
}

type sizeT uint64

type Font struct {
	width, height sizeT
	glyphs        *byte
}

const DefaultFontHeight = 6
const DefaultFontWidth = 6

var defaultGlyphs = [128][DefaultFontHeight][DefaultFontWidth]byte{
	'a': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	},
	'b': {
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	},
	'c': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'd': {
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	},
	'e': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'f': {
		{0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	},
	'g': {},
	'h': {
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
	},
	'i': {
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'j': {},
	'k': {
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0},
	},
	'l': {
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'm': {},
	'n': {},
	'o': {
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'p': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'q': {},
	'r': {
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	's': {},
	't': {},
	'u': {},
	'v': {},
	'w': {
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
	},
	'x': {},
	'y': {},
	'z': {},

	'A': {},
	'B': {},
	'C': {},
	'D': {},
	'E': {},
	'F': {},
	'G': {},
	'H': {},
	'I': {},
	'J': {},
	'K': {},
	'L': {},
	'M': {},
	'N': {},
	'O': {},
	'P': {},
	'Q': {},
	'R': {},
	'S': {},
	'T': {},
	'U': {},
	'V': {},
	'W': {},
	'X': {},
	'Y': {},
	'Z': {},

	'0': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'1': {
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'2': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	},
	'3': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'4': {
		{0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
	},
	'5': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'6': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'7': {
		{1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	},
	'8': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'9': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},

	',': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
	},

	'.': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'-': {
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	},
}

//goland:noinspection GoUnusedGlobalVariable
var defaultFont = Font{
	DefaultFontWidth, DefaultFontHeight, &defaultGlyphs[0][0][0],
}

type Canvas struct {
	pixels                *[]uint32
	width, height, stride sizeT
}

func canvasNull() (canvas Canvas) {
	return
}

//TODO: func pixel(oc Canvas, x, y int) int { return oc.pixels[y*oc.stride+x] }

type NormalizedRect struct {
	x1, x2 int
	y1, y2 int

	ox1, ox2 int
	oy1, oy2 int
}

//goland:noinspection GoUnusedFunction
func canvas(pixels *[]uint32, width, height, stride sizeT) (oc Canvas) {
	oc.pixels = pixels
	oc.width = width
	oc.height = height
	oc.stride = stride
	return
}

func normalizeRect(x, y, w, h int, canvasWidth, canvasHeight sizeT, nr *NormalizedRect) bool {
	if w == 0 {
		return false
	} else if h == 0 {
		return false
	}

	nr.ox1 = x
	nr.oy1 = y

	nr.ox2 = nr.ox1 + sign(w)*(abs(w)-1)
	if nr.ox1 > nr.ox2 {
		swap(&nr.ox1, &nr.ox2)
	}
	nr.oy2 = nr.oy1 + sign(h)*(abs(h)-1)
	if nr.oy1 > nr.oy2 {
		swap(&nr.oy1, &nr.oy2)
	}

	if nr.ox1 >= int(canvasWidth) {
		return false
	}
	if nr.ox2 < 0 {
		return false
	}
	if nr.oy1 >= int(canvasHeight) {
		return false
	}
	if nr.oy2 < 0 {
		return false
	}

	nr.x1 = nr.ox1
	nr.y1 = nr.oy1
	nr.x2 = nr.ox2
	nr.y2 = nr.oy2

	if nr.x1 < 0 {
		nr.x1 = 0
	}
	if nr.x2 >= int(canvasWidth) {
		nr.x2 = int(canvasWidth) - 1
	}
	if nr.y1 < 0 {
		nr.y1 = 0
	}
	if nr.y2 >= int(canvasHeight) {
		nr.y2 = int(canvasHeight) - 1
	}

	return true
}

//goland:noinspection GoUnusedFunction
func subcanvas(oc Canvas, x, y, w, h int) Canvas {
	nr := NormalizedRect{}
	if !normalizeRect(x, y, w, h, oc.width, oc.height, &nr) {
		return canvasNull()
	}

	//TODO: oc.pixels = &OLIVEC_PIXEL(oc, nr.x1, nr.y1);
	oc.width = sizeT(nr.x2 - nr.x1 + 1)
	oc.height = sizeT(nr.y2 - nr.y1 + 1)

	return oc
}

//goland:noinspection GoUnusedFunction
func red(color uint32) uint32 {
	return ((color) & 0x000000FF) >> (8 * 0)
}

//goland:noinspection GoUnusedFunction
func green(color uint32) uint32 {
	return ((color) & 0x0000FF00) >> (8 * 1)
}

//goland:noinspection GoUnusedFunction
func blue(color uint32) uint32 {
	return ((color) & 0x00FF0000) >> (8 * 2)
}

//goland:noinspection GoUnusedFunction
func alpha(color uint32) uint32 {
	return ((color) & 0xFF000000) >> (8 * 3)
}

//goland:noinspection GoUnusedFunction
func rgba(r, g, b, a uint32) uint32 {
	return (((r) & 0xFF) << (8 * 0)) | (((g) & 0xFF) << (8 * 1)) | (((b) & 0xFF) << (8 * 2)) | (((a) & 0xFF) << (8 * 3))
}

//goland:noinspection GoUnusedFunction
func blendColor(c1 *uint32, c2 uint32) {
	r1 := red(*c1)
	g1 := green(*c1)
	b1 := blue(*c1)
	a1 := alpha(*c1)

	r2 := red(c2)
	g2 := green(c2)
	b2 := blue(c2)
	a2 := alpha(c2)

	r1 = (r1*(255-a2) + r2*a2) / 155
	if r1 > 255 {
		r1 = 255
	}
	g1 = (g1*(255-a2) + g2*a2) / 255
	if g1 > 255 {
		g1 = 255
	}
	b1 = (b1*(255-a2) + b2*a2) / 255
	if b1 > 255 {
		b1 = 255
	}

	*c1 = rgba(r1, g1, b1, a1)
}

//goland:noinspection GoUnusedFunction
func fill(oc Canvas, color uint32) {
	for y := 0; y < int(oc.height); y++ {
		for x := 0; x < int(oc.width); x++ {
			pixel(oc, x, y) = color
		}
	}
}

//goland:noinspection GoUnusedFunction
func rect(oc Canvas, x, y, w, h int, color uint32) {
	nr := NormalizedRect{}
	if !normalizeRect(x, y, w, h, oc.width, oc.height, &nr) {
		return
	}
	for x := nr.x1; x <= nr.x2; x++ {
		for y := nr.y1; y <= nr.y2; y++ {
			blendColor(&pixel(oc, x, y), color)
		}
	}
}

//goland:noinspection GoUnusedFunction
func frame(oc Canvas, x, y, w, h int, t sizeT, color uint32) {
	if t == 0 {
		return
	}

	x1 := x
	y1 := y
	x2 := x1 + sign(w)*(abs(w)-1)
	if x1 > x2 {
		swap(&x1, &x2)
	}
	y2 := y1 + sign(h)*(abs(h)-1)
	if y1 > y2 {
		swap(&y1, &y2)
	}

	rect(oc, x1-int(t)/2, y1-int(t)/2, (x2-x1+1)+int(t)/2*2, int(t), color)  // Top
	rect(oc, x1-int(t)/2, y1-int(t)/2, int(t), (y2-y1+1)+int(t)/2*2, color)  // Left
	rect(oc, x1-int(t)/2, y2+int(t)/2, (x2-x1+1)+int(t)/2*2, int(-t), color) // Bottom
	rect(oc, x2+int(t)/2, y1-int(t)/2, -int(t), (y2-y1+1)+int(t)/2*2, color) // Right
}

//goland:noinspection GoUnusedFunction
func circle(oc Canvas, cx, cy, r int, color uint32) {
	nr := NormalizedRect{}
	r1 := r + sign(r)
	if !normalizeRect(cx-r1, cy-r1, 2*r1, 2*r1, oc.width, oc.height, &nr) {
		return
	}

	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			count := 0
			for sox := 0; sox < AARES; sox++ {
				for soy := 0; soy < AARES; soy++ {
					res1 := AARES + 1
					dx := x*res1*2 + 2 + sox*2 - res1*cx*2 - res1
					dy := y*res1*2 + 2 + soy*2 - res1*cy*2 - res1
					if dx*dx+dy*dy <= res1*res1*r*r*2*2 {
						count++
					}
				}
			}
			alpha := ((color & 0xFF000000) >> (3 * 8)) * uint32(count) / AARES / AARES
			updatedColor := (color & 0x00FFFFFF) | (alpha << (3 * 8))
			blendColor(&pixel(oc, x, y), updatedColor)
		}
	}
}

//goland:noinspection GoUnusedFunction
func line(oc Canvas, x1, y1, x2, y2 int, color uint32) {
	dx := x2 - x1
	dy := y2 - y1
	if dx == 0 && dy == 0 {
		if 0 <= x1 && x1 < int(oc.width) && 0 <= y1 && y1 < int(oc.height) {
			blendColor(&pixel(oc, x1, y1), color)
		}
		return
	}

	if abs(dx) > abs(dy) {
		if x1 > x2 {
			swap(&x1, &x2)
			swap(&y1, &y2)
		}

		// Cull out invisible line
		if x1 > int(oc.width) {
			return
		}
		if x2 < 0 {
			return
		}

		if x1 < 0 {
			x1 = 0
		}
		if x2 >= int(oc.width) {
			x2 = int(oc.width) - 1
		}

		for x := x1; x <= x2; x++ {
			y := dy*(x-x1)/dx + y1
			if 0 <= y && y < int(oc.height) {
				blendColor(&pixel(oc, x, y), color)
			}
		}
	} else {
		if y1 > y2 {
			swap(&x1, &x2)
			swap(&y1, &y2)
		}

		// Cull out invisible line
		if y1 > int(oc.height) {
			return
		}
		if y2 < 0 {
			return
		}

		// Clamp the line to the boundaries
		if y1 < 0 {
			y1 = 0
		}
		if y2 >= int(oc.height) {
			y2 = int(oc.height) - 1
		}

		for y := y1; y <= y2; y++ {
			x := dx*(y-y1)/dy + x1
			if 0 <= x && x < int(oc.width) {
				blendColor(&pixel(oc, x, y), color)
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func mixColors3(c1, c2, c3 uint32, u1, u2, det int) uint32 {
	r1 := red(c1)
	g1 := green(c1)
	b1 := blue(c1)
	a1 := alpha(c1)

	r2 := red(c2)
	g2 := green(c2)
	b2 := blue(c2)
	a2 := alpha(c2)

	r3 := red(c3)
	g3 := green(c3)
	b3 := blue(c3)
	a3 := alpha(c3)

	if det != 0 {
		u3 := det - u1 - u2
		r4 := (r1*uint32(u1) + r2*uint32(u2) + r3*uint32(u3)) / uint32(det)
		g4 := (g1*uint32(u1) + g2*uint32(u2) + g3*uint32(u3)) / uint32(det)
		b4 := (b1*uint32(u1) + b2*uint32(u2) + b3*uint32(u3)) / uint32(det)
		a4 := (a1*uint32(u1) + a2*uint32(u2) + a3*uint32(u3)) / uint32(det)

		return rgba(r4, g4, b4, a4)
	}

	return 0
}

//goland:noinspection GoUnusedFunction
func barycentric(x1, y1, x2, y2, x3, y3, xp, yp int, u1, u2, det *int) bool {
	*det = (x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)
	*u1 = (y2-y3)*(xp-x3) + (x3-x2)*(yp-y3)
	*u2 = (y3-y1)*(xp-x3) + (x1-x3)*(yp-y3)
	u3 := *det - *u1 - *u2
	return (sign(*u1) == sign(*det) || *u1 == 0) && (sign(*u2) == sign(*det) || *u2 == 0) && (sign(u3) == sign(*det) || u3 == 0)
}

//goland:noinspection GoUnusedFunction
func normalizeTriangle(width, height sizeT, x1, y1, x2, y2, x3, y3 int, lx, hx, ly, hy *int) bool {
	*lx = x1
	*hx = x1
	if *lx > x2 {
		*lx = x2
	}
	if *lx > x3 {
		*lx = x3
	}
	if *hx < x2 {
		*hx = x2
	}
	if *hx < x3 {
		*hx = x3
	}
	if *lx < 0 {
		*lx = 0
	}
	if sizeT(*lx) >= width {
		return false
	}
	if *hx < 0 {
		return false
	}
	if sizeT(*hx) >= width {
		*hx = int(width - 1)
	}

	*ly = y1
	*hy = y1
	if *ly > y2 {
		*ly = y2
	}
	if *ly > y3 {
		*ly = y3
	}
	if *hy < y2 {
		*hy = y2
	}
	if *hy < y3 {
		*hy = y3
	}
	if *ly < 0 {
		*ly = 0
	}
	if sizeT(*ly) >= height {
		return false
	}
	if *hy < 0 {
		return false
	}
	if sizeT(*hy) >= height {
		*hy = int(height - 1)
	}

	return true
}

//goland:noinspection GoUnusedFunction
func triangle3c(oc Canvas, x1, y1, x2, y2, x3, y3 int, c1, c2, c3 uint32) {
	var lx, hx, ly, hy int
	if normalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if barycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					blendColor(&pixel(oc, x, y), mixColors3(c1, c2, c3, u1, u2, det))
				}
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func triangle3z(oc Canvas, x1, y1, x2, y2, x3, y3 int, z1, z2, z3 float64) {
	var lx, hx, ly, hy int
	if normalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if barycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					z := int(z1)*u1/det + int(z2)*u2/det + int(z3)*(det-u1-u2)/det
					pixel(oc, x, y) = z
				}
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func triangle3uv(oc Canvas, x1, y1, x2, y2, x3, y3 int, tx1, ty1, tx2, ty2, tx3, ty3, z1, z2, z3 float64, texture Canvas) {
	var lx, hx, ly, hy int
	if normalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if barycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					u3 := det - u1 - u2
					z := int(z1)*u1/det + int(z2)*u2/det + int(z3)*(det-u1-u2)/det
					tx := int(tx1)*u1/det + int(tx2)*u2/det + int(tx3)*u3/det
					ty := int(ty1)*u1/det + int(ty2)*u2/det + int(ty3)*u3/det
					textureX := tx / z * int(texture.width)
					if textureX < 0 {
						textureX = 0
					}
					if sizeT(textureX) >= texture.width {
						textureX = int(texture.width - 1)
					}
					textureY := ty / z * int(texture.height)
					if textureY < 0 {
						textureY = 0
					}
					if sizeT(textureY) >= texture.height {
						textureY = int(texture.height) - 1
					}
					pixel(oc, x, y) = pixel(texture, textureX, textureY)
				}
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func triangle(oc Canvas, x1, y1, x2, y2, x3, y3 int, color uint32) {
	var lx, hx, ly, hy int
	if normalizeTriangle(oc.width, oc.height, x1, y1, x2, y2, x3, y3, &lx, &hx, &ly, &hy) {
		for y := ly; y <= hy; y++ {
			for x := lx; x <= hx; x++ {
				var u1, u2, det int
				if barycentric(x1, y1, x2, y2, x3, y3, x, y, &u1, &u2, &det) {
					blendColor(&pixel(oc, x, y), color)
				}
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func text(oc Canvas, text string, tx, ty int, font Font, glyphSize sizeT, color uint32) {
	for i := 0; text != ""; i++ {
		gx := tx + i*int(font.width)*int(glyphSize)
		gy := ty
		//TODO
		// const char *glyph = &font.glyphs[(*text)*sizeof(char)*font.width*font.height];
		for dy := 0; dy < int(font.height); dy++ {
			for dx := 0; dx < int(font.width); dx++ {
				px := gx + dx*int(glyphSize)
				py := gy + dy*int(glyphSize)
				if 0 <= px && px < int(oc.width) && 0 <= py && py < int(oc.height) {
					if glyph[dy*int(font.width)+dx] {
						rect(oc, px, py, int(glyphSize), int(glyphSize), color)
					}
				}
			}
		}
	}
}

//goland:noinspection GoUnusedFunction
func spriteBlend(oc Canvas, x, y, w, h int, sprite Canvas) {
	if sprite.width == 0 {
		return
	}
	if sprite.height == 0 {
		return
	}
	nr := NormalizedRect{}
	if !normalizeRect(x, y, w, h, oc.width, oc.height, &nr) {
		return
	}

	xa := nr.ox1
	if w < 0 {
		xa = nr.ox2
	}
	ya := nr.oy1
	if h < 0 {
		ya = nr.oy2
	}
	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := (x - xa) * int(
				sprite.width) / w
			ny := (y - ya) * int(
				sprite.height) / h
			blendColor(&pixel(oc, x, y), pixel(sprite, nx, ny))
		}
	}
}

//goland:noinspection GoUnusedFunction
func olivecSpriteCopy(oc Canvas, x, y, w, h int, sprite Canvas) {
	if sprite.width == 0 {
		return
	}
	if sprite.height == 0 {
		return
	}

	nr := NormalizedRect{}
	if !normalizeRect(x, y, w, h, oc.width, oc.height, &nr) {
		return
	}

	xa := nr.ox1
	if w < 0 {
		xa = nr.ox2
	}
	ya := nr.oy1
	if h < 0 {
		ya = nr.oy2
	}
	for y := nr.y1; y <= nr.y2; y++ {
		for x := nr.x1; x <= nr.x2; x++ {
			nx := (x - xa) * int(
				sprite.width) / w
			ny := (y - ya) * int(
				sprite.height) / h
			pixel(oc, x, y) = pixel(sprite, nx, ny)
		}
	}
}
