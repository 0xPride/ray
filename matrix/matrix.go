package matrix

import (
	"github.com/0xpride/ray/objects"
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

func (self *Matrix) GetColum(col uint32) *objects.Tuple {
	tup := new(objects.Tuple)

	tup.X = self[0][col]
	tup.Y = self[1][col]
	tup.Z = self[2][col]
	tup.W = self[3][col]

	return tup
}

func (self *Matrix) GetRow(row uint32) *objects.Tuple {
	tup := new(objects.Tuple)

	tup.X = self[row][0]
	tup.Y = self[row][1]
	tup.Z = self[row][2]
	tup.W = self[row][3]

	return tup
}

func (self *Matrix) MultTuple(t *objects.Tuple) *objects.Tuple {
	tup := new(objects.Tuple)

	tup.X = self.GetRow(0).Dot(t)
	tup.Y = self.GetRow(1).Dot(t)
	tup.Z = self.GetRow(2).Dot(t)
	tup.W = self.GetRow(3).Dot(t)

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
	identity[3][0] = x
	identity[3][1] = y
	identity[3][2] = z
	return identity
}
