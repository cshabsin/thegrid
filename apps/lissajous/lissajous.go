package main

import (
	"math"
	"strconv"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/canvas"
)

type params struct {
	size    float64
	cycles  float64
	resol   float64
	freq    float64
	speed   float64
}

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

	p := &params{
		size:    200,
		cycles:  5,
		resol:   0.001,
		freq:    3.0,
		speed:   0.001,
	}

	setupControls(doc, p)

	var startTime float64

	var renderFrame func(timestamp float64)
	renderFrame = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime
		phase := elapsed * p.speed

		// Clear the entire canvas (the back buffer).
		ctx.ClearRect(-width/2, -height/2, width, height)

		ctx.BeginPath()
		// Draw one complete, closed curve on each frame.
		// The animation comes from varying the phase, not from drawing more of the curve.
		for t := 0.0; t < p.cycles*2*math.Pi; t += p.resol {
			x := math.Sin(t)
			y := math.Sin(t*p.freq + phase)
			screenX := x * p.size
			screenY := y * p.size
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

func setupControls(doc js.DOMDocument, p *params) {
	addSlider := func(id string, setter func(float64)) {
		slider := doc.GetElementByID(id)
		valueSpan := doc.GetElementByID(id + "-value")
		slider.AddEventListener("input", func(el js.DOMElement, ev js.DOMEvent) {
			valStr := el.Get("value").String()
			valueSpan.SetText(valStr)
			val, _ := strconv.ParseFloat(valStr, 64)
			setter(val)
		})
	}

	addSlider("size", func(v float64) { p.size = v })
	addSlider("cycles", func(v float64) { p.cycles = v })
	addSlider("resol", func(v float64) { p.resol = v })
	addSlider("freq", func(v float64) { p.freq = v })
	addSlider("speed", func(v float64) { p.speed = v })
}