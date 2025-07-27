package main

import (
	"fmt"
	"math"

	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	canvas := document.GetElementByID("map-canvas")
	ctx := canvas.GetContext("2d")

	ctx.Call("scale", 0.1, 0.1)

	width := canvas.Get("width").Float() * 10
	height := canvas.Get("height").Float() * 10

	var startTime float64
	var animate func(timestamp float64)
	animate = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime

		ctx.Call("clearRect", 0, 0, width, height)

		for i := 0; i < 10000; i += 70 {
			i64 := math.Mod(float64(i*70)+elapsed*8.5, 10000)
			ctx.Set("strokeStyle", fmt.Sprintf("hsl(%d, 100%%, 50%%)", i*360/10000))
			ctx.Call("beginPath")
			ctx.Call("moveTo", i64, 0)
			ctx.Call("lineTo", 0, 10000-i64)
			ctx.Call("stroke")

			ctx.Call("beginPath")
			ctx.Call("moveTo", 10000-i64, 10000)
			ctx.Call("lineTo", 0, 10000-i64)
			ctx.Call("stroke")

			ctx.Call("beginPath")
			ctx.Call("moveTo", 10000-i64, 10000)
			ctx.Call("lineTo", 10000, i64)
			ctx.Call("stroke")

			ctx.Call("beginPath")
			ctx.Call("moveTo", i64, 0)
			ctx.Call("lineTo", 10000, i64)
			ctx.Call("stroke")
		}

		js.RequestAnimationFrame(animate)
	}
	js.RequestAnimationFrame(animate)

	select {}
}
