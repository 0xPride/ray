package core

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func TupleInit(x float64, y float64, z float64, w float64) *Tuple {
	tup := &Tuple{X: x, Y: y, Z: z, W: w}

	return tup
}

func VectorInit(x float64, y float64, z float64) *Tuple {
	return TupleInit(x, y, z, 0)
}

func PointInit(x float64, y float64, z float64) *Tuple {
	return TupleInit(x, y, z, 1)
}

func (self *Tuple) Add(t *Tuple) *Tuple {
	tup := new(Tuple)

	tup.X = self.X + t.X
	tup.Y = self.Y + t.Y
	tup.Z = self.Z + t.Z
	tup.W = self.W + t.W

	if tup.W == 2 {
		tup.W = 1
	}
	return tup
}

func (self *Tuple) Sub(t *Tuple) *Tuple {
	tup := new(Tuple)

	tup.X = self.X - t.X
	tup.Y = self.Y - t.Y
	tup.Z = self.Z - t.Z
	tup.W = self.W - t.W

	return tup
}

func (self *Tuple) Negate() {
	self.X *= -1
	self.Y *= -1
	self.Z *= -1
	self.W *= -1
}

func (self *Tuple) GetNegate() *Tuple {
	tup := new(Tuple)

	tup.X = self.X * -1
	tup.Y = self.Y * -1
	tup.Z = self.Z * -1
	tup.W = self.W * -1

	return tup
}

func (self *Tuple) Scale(n float64) *Tuple {
	tup := new(Tuple)

	tup.X = self.X * n
	tup.Y = self.Y * n
	tup.Z = self.Z * n
	tup.W = self.W * n

	return tup
}

func (self *Tuple) Div(n float64) *Tuple {
	tup := new(Tuple)

	n = 1 / n
	tup.X = self.X * n
	tup.Y = self.Y * n
	tup.Z = self.Z * n
	tup.W = self.W * n

	return tup
}

func (self *Tuple) Len() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y + self.Z*self.Z + self.W + self.W)
}

func (self *Tuple) Norm() {
	l := self.Len()

	self.X /= l
	self.Y /= l
	self.Z /= l
	self.W /= l
}

func (self *Tuple) GetNorm() *Tuple {
	tup := new(Tuple)
	l := self.Len()

	tup.X = self.X / l
	tup.Y = self.Y / l
	tup.Z = self.Z / l
	tup.W = self.W / l

	return tup
}

func (self *Tuple) Dot(t *Tuple) float64 {
	return self.X*t.X + self.Y*t.Y + self.Z*t.Z + self.W*t.W
}

func (self *Tuple) Cross(t *Tuple) *Tuple {
	tup := new(Tuple)

	tup.X = self.Y*t.Z - self.Z*t.Y
	tup.Y = self.Z*t.X - self.X*t.Z
	tup.Z = self.X*t.Y - self.Y*t.X

	return tup
}
