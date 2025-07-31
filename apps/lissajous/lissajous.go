package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/canvas"
)

type params struct {
	size    float64
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
		freq:    3.0,
		speed:   0.001,
	}

	setupControls(doc, p)

	var startTime, lastFrameTime float64
	resol := 0.001 // Initial resolution
	fpsDisplay := doc.GetElementByID("fps")
	var frameCount int

	var renderFrame func(timestamp float64)
	renderFrame = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
			lastFrameTime = timestamp
		}
		delta := timestamp - lastFrameTime
		lastFrameTime = timestamp
		fps := 1000 / delta

		// Adjust resolution based on frame rate
		if fps < 58 && resol < 0.01 {
			resol *= 1.1 // Decrease resolution
		} else if fps > 60 && resol > 0.0001 {
			resol *= 0.9 // Increase resolution
		}

		// Update FPS display every 10 frames
		frameCount++
		if frameCount%10 == 0 {
			fpsDisplay.SetText(fmt.Sprintf("%.1f FPS", fps))
		}

		elapsed := timestamp - startTime
		phase := elapsed * p.speed

		// Clear the entire canvas (the back buffer).
		ctx.ClearRect(-width/2, -height/2, width, height)

		// Calculate cycles needed to close the curve.
		num := int(p.freq * 1000)
		den := 1000
		g := gcd(num, den)
		cycles := float64(num / g)

		ctx.BeginPath()
		// Draw one complete, closed curve on each frame.
		// The animation comes from varying the phase, not from drawing more of the curve.
		for t := 0.0; t < cycles*2*math.Pi; t += resol {
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
	addSlider("freq", func(v float64) { p.freq = v })
	addSlider("speed", func(v float64) { p.speed = v })
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
