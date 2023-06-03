package canvas

import "github.com/0xpride/ray/objects"

type Color struct {
	objects.Vec3d
}

func (self *Color) Red() *float64 {
	return &self.X
}

func (self *Color) Green() *float64 {
	return &self.Y
}

func (self *Color) Blue() *float64 {
	return &self.Z
}

func (self *Color) Mult(c *Color) *Color {
	color := new(Color)

	*color.Red() = *self.Red() * *c.Red()
	*color.Green() = *self.Green() * *c.Green()
	*color.Blue() = *self.Blue() * *c.Blue()

	return color
}
