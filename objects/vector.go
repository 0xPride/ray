package objects

import "math"

type Vector3d struct {
	X, Y, Z float64
}

func (self *Vector3d) Add(v *Vector3d) *Vector3d {
	vec := new(Vector3d)

	vec.X = self.X + v.X
	vec.Y = self.Y + v.Y
	vec.Z = self.Z + v.Z

	return vec
}

func (self *Vector3d) Sub(v *Vector3d) *Vector3d {
	vec := new(Vector3d)

	vec.X = self.X - v.X
	vec.Y = self.Y - v.Y
	vec.Z = self.Z - v.Z

	return vec
}

func (self *Vector3d) Scale(s float64) *Vector3d {
	vec := new(Vector3d)

	vec.X = self.X * s
	vec.Y = self.Y * s
	vec.Z = self.Z * s

	return vec
}

func (self *Vector3d) Div(s float64) *Vector3d {
	vec := new(Vector3d)

	s = 1 / s
	vec.X = self.X * s
	vec.Y = self.Y * s
	vec.Z = self.Z * s

	return vec
}

func (self *Vector3d) Len() float64 {
	return math.Sqrt(self.LenPow2())
}

func (self *Vector3d) LenPow2() float64 {
	return math.Pow(self.X, 2) + math.Pow(self.Y, 2) + math.Pow(self.Z, 2)
}

func (self *Vector3d) Negate() {
	self.X *= -1
	self.Y *= -1
	self.Z *= -1
}

func (self *Vector3d) GetNegate() *Vector3d {
	vec := new(Vector3d)

	vec.X = -self.X
	vec.Y = -self.Y
	vec.Z = -self.Z

	return vec
}

func (self *Vector3d) Dot(v *Vector3d) float64 {
	return self.X*v.X + self.Y*v.Y + self.Z*v.Z
}

func (self *Vector3d) Cross(v *Vector3d) *Vector3d {
	vec := new(Vector3d)

	vec.X = self.Y*v.Z - self.Z*v.Y
	vec.Y = self.Z*v.X - self.X*v.Z
	vec.Z = self.X*v.Y - self.Y*v.X

	return vec
}

func (self *Vector3d) Norm() {
	len := self.Len()

	self.X /= len
	self.Y /= len
	self.Z /= len
}

func (self *Vector3d) GetNorm() *Vector3d {
	vec := new(Vector3d)
	len := self.Len()

	vec.X = self.X / len
	vec.Y = self.Y / len
	vec.Z = self.Z / len

	return vec
}
