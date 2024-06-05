package asciigl2d

import (
	"fmt"
)

// Vertices defines the area where drawing is allowed, represented by a 2D array of coordinates.
type Vertices [4][2]int

// Pixel represents a single unit of rendering information.
// It contains the X and Y coordinates, color code, and character as integers.
type Pixel [4]int

// Clears pixel
func(p *Pixel) Clear(){
	p[2] = int(ColorNone)
	p[3] = '-'
} 

// ColorCode represents the ANSI color codes for text coloring.
type ColorCode int

// Main vertices object instance.
var mVertices Vertices

// Current buffer holds the current state of pixels within the drawable area.
var currentBuffer []Pixel

// Previous buffer contains the previous state of pixels for comparison and rendering updates.
var previousBuffer []Pixel

// Available ANSI color codes.
const (
	ColorNone    ColorCode = 99
	ColorBlack   ColorCode = 30
	ColorRed     ColorCode = 31
	ColorGreen   ColorCode = 32
	ColorYellow  ColorCode = 33
	ColorBlue    ColorCode = 34
	ColorMagenta ColorCode = 35
	ColorCyan    ColorCode = 36
	ColorWhite   ColorCode = 37
)

// SetVertices initializes the main vertices object and renders the first frame.
func SetVertices(v Vertices) {
	mVertices = v

	// Initialize the current buffer with pixels set to default values.
	currentBuffer = make([]Pixel, mVertices[3][1]*mVertices[1][0])
	for i := 0; i < mVertices[3][1]; i++ {
		for k := 0; k < mVertices[1][0]; k++ {
			// Set default values for each pixel.
			currentBuffer[mVertices[1][0]*i+k] = Pixel{i, k, int(ColorNone), '-'}
		}
	}

	// Create a copy of the current buffer for comparison.
	previousBuffer = make([]Pixel, len(currentBuffer))
	copy(previousBuffer, currentBuffer)

	// Render the first frame.
	for _, p := range currentBuffer {
		asciiSyncPixel(p) // Render each pixel.
	}
}

// rayCast performs ray casting to determine which pixels are inside a polygon defined by the vertices.
func rayCast(vertices [][2]int) []*Pixel {
	pixels := []*Pixel{}

	for i, p := range currentBuffer {

		n := len(vertices)
		inside := false
		for i, j := 0, n-1; i < n; i++ {
			if (vertices[i][1] > p[1]) != (vertices[j][1] > p[1]) &&
				p[0] < (vertices[j][0]-vertices[i][0])*(p[1]-vertices[i][1])/(vertices[j][1]-vertices[i][1])+vertices[i][0] {
				inside = !inside
			}
			j = i
		}
		if inside {
			pixels = append(pixels, &currentBuffer[i])
		}
	}

	return pixels
}

// GetAreaPixels returns the pixels within a polygon defined by the vertices.
func GetAreaPixels(vertices [][2]int) []*Pixel {
	return rayCast(vertices)
}

// FillArea fills a polygon defined by the vertices with a specified color and character.
func FillArea(areaVertices [][2]int, color ColorCode, char int) {
	pixels := rayCast(areaVertices)

	for _, p := range pixels {
		p[2] = int(color) // Set color code.
		p[3] = char       // Set character.
	}
}

// PixelPointer returns a pointer to a pixel at specified coordinates (x, y).
func PixelPointer(x int, y int) *Pixel {
	return &currentBuffer[y*mVertices[1][0]+x]
}

// PixelColorBuff returns a pointer to the color buffer of the pixel at coordinates (x, y).
func PixelColorBuff(x int, y int) *int {
	return &currentBuffer[y*mVertices[1][0]+x][2]
}

// PixelCharBuff returns a pointer to the character buffer of the pixel at coordinates (x, y).
func PixelCharBuff(x int, y int) *int {
	return &currentBuffer[y*mVertices[1][0]+x][3]
}

// GenerateFrame checks for differences between the current and previous buffer and updates individual pixels accordingly.
func GenerateFrame() {
	for i := range currentBuffer {
		if currentBuffer[i][2] != previousBuffer[i][2] || currentBuffer[i][3] != previousBuffer[i][3] {
			// Asynchronously update individual pixel.
			go asciiSyncPixel(currentBuffer[i])
		}
	}
	copy(previousBuffer, currentBuffer)
}

// asciiSyncPixel synchronizes ASCII terminal output with the pixel information.
func asciiSyncPixel(p Pixel) {
	setCursorPosition := fmt.Sprintf("\033[%d;%dH", p[0], p[1]) // Set cursor position.
	colorCode := fmt.Sprintf("\033[%dm", p[2])                  // Get ANSI color code.
	if p[2] != int(ColorNone) {
		fmt.Printf("%s%s%c\033[0m", setCursorPosition, colorCode, rune(p[3])) // Print pixel information with proper formatting.
	}
}
