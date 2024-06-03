# asciigl2d
AsciiGL2D is a lightweight ASCII graphics library supporting animations, colors, and precise character rendering for creating dynamic visualizations.

This library provides a convenient abstraction for generating 2D ASCII graphics. By defining object vertices, you can effortlessly draw shapes. Each "pixel," though technically a character within the ASCII terminal, functions akin to a pixel in real displays, undergoing updates in position, color, or character directly affecting the frame. Additionally, frame generation optimizes efficiency by rendering only those individual pixels that have changed.

Furthermore, this library grants granular control over pixels, facilitating direct manipulation of their position, color, and character, while also offering a diverse set of built-in methods for enhanced functionality.

## Installation
```bash
go get github.com/burakturkerdev/asciigl2d
```

## Drawing triangle and some effects

![triangle](https://github.com/burakturkerdev/asciigl2d/assets/166562458/21dbdf09-2a18-435d-b8a0-fc84c527d187)


```go
func main() {
	// Define the vertices for the main drawable area.
	vertices := asciigl2d.Vertices{{0, 0}, {100, 0}, {100, 100}, {0, 100}}

	// Initialize the drawable area with the specified vertices.
	asciigl2d.SetVertices(vertices)

	// Define a function to continuously generate frames.
	frameLoop := func() {
		for {
			// Generate and render the next frame at regular intervals.
			asciigl2d.GenerateFrame()
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Launch the frame generation process concurrently.
	go frameLoop()

	// Define a structure representing a triangle.
	type triangle struct {
		corner1 [2]int
		corner2 [2]int
		corner3 [2]int
	}

	// Specify the coordinates for the corners of the triangle.
	tangle := triangle{[2]int{0, 18}, [2]int{16, 2}, [2]int{32, 18}}

	// Retrieve all pixels within the triangle.
	pixels := asciigl2d.GetAreaPixels([][2]int{tangle.corner1, tangle.corner2, tangle.corner3})

	// Example Effect 1: Character and Color Transition

	// Gradually transition each pixel's character inside the triangle to '-' while modifying its color.
	for _, p := range pixels {
		p[2] = int(asciigl2d.ColorRed) // Set color to red
		p[3] = '-'                     // Change character to '-'
		time.Sleep(5 * time.Millisecond)
	}

	// Reverse the transition, changing each pixel's character inside the triangle to 'n' while modifying its color.
	for i := len(pixels) - 1; i >= 0; i-- {
		pixels[i][2] = int(asciigl2d.ColorBlue) // Set color to blue
		pixels[i][3] = 'n'                      // Change character to 'n'
		time.Sleep(5 * time.Millisecond)
	}

	// Transition each pixel's character inside the triangle to 'x' while modifying its color.
	for _, p := range pixels {
		p[2] = int(asciigl2d.ColorGreen) // Set color to green
		p[3] = 'x'                       // Change character to 'x'
		time.Sleep(5 * time.Millisecond)
	}

	// Reverse the transition, changing each pixel's character inside the triangle to 'f' while modifying its color.
	for i := len(pixels) - 1; i >= 0; i-- {
		pixels[i][2] = int(asciigl2d.ColorCyan) // Set color to cyan
		pixels[i][3] = 'f'                      // Change character to 'f'
		time.Sleep(5 * time.Millisecond)
	}

	// Example Effect 2: Color Cycling

	for {
		// Cycle through colors for all pixels inside the triangle.
		for _, p := range pixels {
			p[2] = int(asciigl2d.ColorRed) // Set color to red
		}
		for _, p := range pixels {
			p[2] = int(asciigl2d.ColorGreen) // Set color to green
		}
		for _, p := range pixels {
			p[2] = int(asciigl2d.ColorBlue) // Set color to blue
		}
	}
}```
