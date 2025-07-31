package main

import (
	"math"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/canvas"
)

const (
	size = 200  // image canvas size
	freq = 3.0  // relative frequency of y oscillator
)

func main() {
	doc := js.Document()
	canvasEl := canvas.Get("canvas")
	canvasEl.FullSize()
	ctx := canvasEl.GetContext("2d")

	width := doc.Body().Get("clientWidth").Float()
	height := doc.Body().Get("clientHeight").Float()
	ctx.SetLineWidth(1)
	ctx.SetStrokeStyle("green")

	// Center the origin
	ctx.Call("translate", width/2, height/2)

	var t float64

	// renderFrame is the core animation loop.
	// We need to declare it before we use it for the recursive call.
	var renderFrame func(timestamp float64)

	renderFrame = func(timestamp float64) {
		// Clear the entire canvas (the back buffer).
		ctx.ClearRect(-width/2, -height/2, width, height)

		ctx.BeginPath()
		// We draw the curve from the beginning up to the current time `t` on each frame.
		for i := 0.0; i < t; i += 0.01 {
			x := math.Sin(i)
			y := math.Sin(i*freq + t)
			screenX := x * size
			screenY := y * size
			if i == 0 {
				ctx.MoveTo(screenX, screenY)
			} else {
				ctx.LineTo(screenX, screenY)
			}
		}
		ctx.Stroke()

		// Increment time for the next frame.
		t += 0.02

		// Schedule the next frame.
		js.RequestAnimationFrame(renderFrame)
	}

	// Start the animation loop.
	js.RequestAnimationFrame(renderFrame)

	// Prevent the main function from exiting.
	select {}
}