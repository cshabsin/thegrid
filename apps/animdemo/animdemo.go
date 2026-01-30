package main

import (
	"fmt"
	"math"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/canvas"
)

func main() {
	canvas := canvas.Get("map-canvas")
	ctx := canvas.GetContext("2d")

	ctx.Scale(0.1, 0.1)

	width := canvas.Get("width").Float() * 10
	height := canvas.Get("height").Float() * 10

	var startTime float64
	var animate func(timestamp float64)
	animate = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime

		ctx.ClearRect(0, 0, width, height)

		steps := 10000
		phase := 720
		speed := 1
		thicknessSpeed := 300

		for i := 0; i < steps; i += 10 {
			i64 := math.Mod(float64(i)+elapsed*float64(speed), float64(steps))
			ctx.SetStrokeStyle(fmt.Sprintf("hsl(%d, 100%%, 50%%)", (i*phase/steps)%360))
			ctx.SetLineWidth(int(math.Sin(timestamp/float64(thicknessSpeed))*5 + 6))
			ctx.BeginPath()
			ctx.MoveTo(i64, 0)
			ctx.LineTo(0, float64(steps)-i64)
			ctx.Stroke()

			ctx.BeginPath()
			ctx.MoveTo(10000-i64, float64(steps))
			ctx.LineTo(0, 10000-i64)
			ctx.Stroke()

			ctx.BeginPath()
			ctx.MoveTo(10000-i64, float64(steps))
			ctx.LineTo(10000, i64)
			ctx.Stroke()

			ctx.BeginPath()
			ctx.MoveTo(i64, 0)
			ctx.LineTo(10000, i64)
			ctx.Stroke()
		}

		js.RequestAnimationFrame(animate)
	}
	js.RequestAnimationFrame(animate)

	select {}
}
