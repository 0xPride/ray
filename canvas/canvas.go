package canvas

import (
	"fmt"
	"os"
)

type Canvas struct {
	Width, Height uint32
	Data          [][]Color
	Filename      string
}

func (self *Canvas) Init(width uint32, height uint32, filename string) {
	self.Width = width
	self.Height = height
	self.Filename = filename
	self.Data = make([][]Color, height)
	var i uint32
	for i = 0; i < height; i++ {
		self.Data[i] = make([]Color, width)
	}
}

func (self *Canvas) Write(width uint32, height uint32, c *Color) {
	self.Data[height][width] = *c
}

func (self *Canvas) SaveToDisk() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/" + self.Filename + ".ppm"
	header := "P3\n" + fmt.Sprint(self.Width, " ", fmt.Sprint(self.Height), "\n")
	var body string
	var i uint32
	var j uint32
	for ; i < self.Height; i++ {
		j = 0
		for ; j < self.Width; j++ {
			body += fmt.Sprint(uint32(*self.Data[i][j].Red()*255)) +
				" " + fmt.Sprint(uint32(*self.Data[i][j].Green()*255)) +
				" " + fmt.Sprint(uint32(*self.Data[i][j].Blue()*255))
			if j != self.Width-1 {
				body += " "
			}
		}
		body += "\n"
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(header)
	f.WriteString(body)
}
