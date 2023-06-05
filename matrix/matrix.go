package matrix

import (
	"math"

	"github.com/0xpride/ray/core"
	"github.com/0xpride/ray/utils"
)

type Matrix [4][4]float64
type Matrix3 [3][3]float64
type Matrix2 [2][2]float64
type HasDeterminante interface {
	Determinant() float64
}

func (self *Matrix) Mult(m *Matrix) *Matrix {
	matrix := Matrix{}

	for row := range matrix {
		for col := range matrix {
			matrix[row][col] = self.GetRow(uint32(row)).Dot(m.GetColum(uint32(col)))
		}
	}
	return &matrix
}

func (self *Matrix) Scale(num float64) {
	for i := range self {
		for j := range self {
			self[i][j] *= num
		}
	}
}

func (self *Matrix) GetColum(col uint32) *core.Tuple {
	tup := new(core.Tuple)

	tup.X = self[0][col]
	tup.Y = self[1][col]
	tup.Z = self[2][col]
	tup.W = self[3][col]

	return tup
}

func (self *Matrix) GetRow(row uint32) *core.Tuple {
	tup := new(core.Tuple)

	tup.X = self[row][0]
	tup.Y = self[row][1]
	tup.Z = self[row][2]
	tup.W = self[row][3]

	return tup
}

func (self *Matrix) MultTuple(t *core.Tuple) *core.Tuple {
	tup := new(core.Tuple)

	// tup.X = self.GetRow(0).Dot(t)
	// tup.Y = self.GetRow(1).Dot(t)
	// tup.Z = self.GetRow(2).Dot(t)
	// tup.W = self.GetRow(3).Dot(t)
	tup.X = self[0][0]*t.X + self[0][1]*t.Y + self[0][2]*t.Z + self[0][3]*t.W
	tup.Y = self[1][0]*t.X + self[1][1]*t.Y + self[1][2]*t.Z + self[1][3]*t.W
	tup.Z = self[2][0]*t.X + self[2][1]*t.Y + self[2][2]*t.Z + self[2][3]*t.W
	tup.W = self[3][0]*t.X + self[3][1]*t.Y + self[3][2]*t.Z + self[3][3]*t.W

	return tup
}

func IdentityMatrix() *Matrix {
	m := new(Matrix)

	m[0][0] = 1
	m[1][1] = 1
	m[2][2] = 1
	m[3][3] = 1

	return m
}

func (self *Matrix) Transpose() {
	*self = *self.GetTranspose()
}

func (self *Matrix) GetTranspose() *Matrix {
	m := new(Matrix)

	for row := range *self {
		for col := range *self {
			m[row][col] = self[col][row]
		}
	}
	return m
}

func (self *Matrix) SubMatrix(row uint32, col uint32) *Matrix3 {
	m := new(Matrix3)
	var i, j, r, c uint32
	for ; i < 3; r++ {
		c = 0
		j = 0
		if r == row {
			r++
		}
		for ; j < 3; c++ {
			if c == col {
				c++
			}
			m[i][j] = self[r][c]
			j++
		}
		i++
	}
	return m
}

func (self *Matrix3) SubMatrix(row uint32, col uint32) *Matrix2 {
	m := new(Matrix2)
	var i, j, r, c uint32
	for ; i < 2; r++ {
		c = 0
		j = 0
		if r == row {
			r++
		}
		for ; j < 2; c++ {
			if c == col {
				c++
			}
			m[i][j] = self[r][c]
			j++
		}
		i++
	}
	return m
}

func (self *Matrix2) Determinant() float64 {
	return self[0][0]*self[1][1] - self[0][1]*self[1][0]
}

func (self *Matrix3) Minor(row uint32, col uint32) float64 {
	return self.SubMatrix(row, col).Determinant()
}

func (self *Matrix3) Cofactor(row uint32, col uint32) float64 {
	if (row+col)%2 != 0 {
		return -self.Minor(row, col)
	}
	return self.Minor(row, col)
}

func (self *Matrix3) Determinant() float64 {
	return self.Cofactor(0, 0)*self[0][0] + self.Cofactor(0, 1)*self[0][1] + self.Cofactor(0, 2)*self[0][2]
}

func (self *Matrix) Cofactor(row uint32, col uint32) float64 {
	cof := self.SubMatrix(row, col).Determinant()
	if (row+col)%2 != 0 {
		return -cof
	}
	return cof
}

func (self *Matrix) Determinant() float64 {
	return self.Cofactor(0, 0)*self[0][0] +
		self.Cofactor(0, 1)*self[0][1] +
		self.Cofactor(0, 2)*self[0][2] +
		self.Cofactor(0, 3)*self[0][3]
}

func IsInvertable(m HasDeterminante) bool {
	if utils.Equalf(m.Determinant(), 0) {
		return false
	}
	return true
}

func (self *Matrix) Inverse() *Matrix {
	inv := new(Matrix)

	if !IsInvertable(self) {
		panic("Can not invert a matrix")
	}
	invDet := 1 / self.Determinant()
	for i := uint32(0); i < 4; i++ {
		for j := uint32(0); j < 4; j++ {
			inv[j][i] = self.Cofactor(i, j) * invDet
		}
	}
	return inv
}

func GetTransformMatrix(x float64, y float64, z float64) *Matrix {
	identity := IdentityMatrix()
	identity[0][3] = x
	identity[1][3] = y
	identity[2][3] = z
	return identity
}

func GetScaleMatrix(x float64, y float64, z float64) *Matrix {
	identity := IdentityMatrix()
	identity[0][0] = x
	identity[1][1] = y
	identity[2][2] = z
	return identity
}

func GetRotateXMatrix(rad float64) *Matrix {
	identity := IdentityMatrix()
	identity[1][1] = math.Cos(rad)
	identity[1][2] = -math.Sin(rad)
	identity[2][1] = math.Sin(rad)
	identity[2][2] = math.Cos(rad)
	return identity
}

func GetRotateYMatrix(rad float64) *Matrix {
	identity := IdentityMatrix()
	identity[0][0] = math.Cos(rad)
	identity[0][2] = math.Sin(rad)
	identity[2][0] = -math.Sin(rad)
	identity[2][2] = math.Cos(rad)
	return identity
}

func GetRotateZMatrix(rad float64) *Matrix {
	identity := IdentityMatrix()
	identity[0][0] = math.Cos(rad)
	identity[0][1] = -math.Sin(rad)
	identity[1][0] = math.Sin(rad)
	identity[1][1] = math.Cos(rad)
	return identity
}

func GetShearingMatrix(Xy float64, Xz float64, Yx float64, Yz float64, Zx float64, Zy float64) *Matrix {
	identity := IdentityMatrix()
	identity[0][1] = Xy
	identity[0][2] = Xz
	identity[1][0] = Yx
	identity[1][2] = Yz
	identity[2][0] = Zx
	identity[2][1] = Zy
	return identity
}
