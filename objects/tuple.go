package objects

import "math"

type tuple struct {
	X, Y, Z, W float64
}

func (self *tuple) Add(t *tuple) *tuple {
	tup := new(tuple)

	tup.X = self.X + t.X
	tup.Y = self.Y + t.Y
	tup.Z = self.Z + t.Z
	tup.W = self.W + t.W

	if tup.W == 2 {
		tup.W = 1
	}
	return tup
}

func (self *tuple) Sub(t *tuple) *tuple {
	tup := new(tuple)

	tup.X = self.X - t.X
	tup.Y = self.Y - t.Y
	tup.Z = self.Z - t.Z
	tup.W = self.W - t.W

	return tup
}

func (self *tuple) Negate() {
	self.X *= -1
	self.Y *= -1
	self.Z *= -1
	self.W *= -1
}

func (self *tuple) GetNegate() *tuple {
	tup := new(tuple)

	tup.X = self.X * -1
	tup.Y = self.Y * -1
	tup.Z = self.Z * -1
	tup.W = self.W * -1

	return tup
}

func (self *tuple) Scale(n float64) *tuple {
	tup := new(tuple)

	tup.X = self.X * n
	tup.Y = self.Y * n
	tup.Z = self.Z * n
	tup.W = self.W * n

	return tup
}

func (self *tuple) Div(n float64) *tuple {
	tup := new(tuple)

	n = 1 / n
	tup.X = self.X * n
	tup.Y = self.Y * n
	tup.Z = self.Z * n
	tup.W = self.W * n

	return tup
}

func (self *tuple) Len() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y + self.Z*self.Z + self.W + self.W)
}

func (self *tuple) Norm() {
	l := self.Len()

	self.X /= l
	self.Y /= l
	self.Z /= l
	self.W /= l
}

func (self *tuple) GetNorm() *tuple {
	tup := new(tuple)
	l := self.Len()

	tup.X = self.X / l
	tup.Y = self.Y / l
	tup.Z = self.Z / l
	tup.W = self.W / l

	return tup
}

func (self *tuple) Dot(t *tuple) float64 {
	return self.X*t.X + self.Y*t.Y + self.Z*t.Z + self.W + t.W
}

func (self *tuple) Cross(t *tuple) *tuple {
	tup := new(tuple)

	tup.X = self.Y*t.Z - self.Z*t.Y
	tup.Y = self.Z*t.X - self.X*t.Z
	tup.Z = self.X*t.Y - self.Y*t.X

	return tup
}
