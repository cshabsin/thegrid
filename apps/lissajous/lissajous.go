package main

import (
	"math"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/canvas"
)

const (
	size    = 200 // image canvas size
	cycles  = 5   // number of revolutions
	resol   = 0.001 // angular resolution
	freq    = 3.0 // frequency of y oscillator
	speed   = 0.001 // animation speed
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

	var startTime float64

	var renderFrame func(timestamp float64)
	renderFrame = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime
		phase := elapsed * speed

		// Clear the entire canvas (the back buffer).
		ctx.ClearRect(-width/2, -height/2, width, height)

		ctx.BeginPath()
		// Draw one complete, closed curve on each frame.
		// The animation comes from varying the phase, not from drawing more of the curve.
		for t := 0.0; t < cycles*2*math.Pi; t += resol {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			screenX := x * size
			screenY := y * size
			if t == 0 {
				ctx.MoveTo(screenX, screenY)
			} else {
				ctx.LineTo(screenX, screenY)
			}
		}
		ctx.Stroke()

		// Schedule the next frame.
		js.RequestAnimationFrame(renderFrame)
	}

	// Start the animation loop.
	js.RequestAnimationFrame(renderFrame)

	// Prevent the main function from exiting.
	select {}
}
