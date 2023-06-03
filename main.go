package main

import "github.com/0xpride/ray/canvas"

func main() {
	var c canvas.Canvas
	var col canvas.Color

	*col.Red() = 1
	*col.Green() = 1
	*col.Blue() = 1
	c.Init(4, 8, "output")
	c.Write(3, 7, &col)
	c.SaveToDisk()
}
