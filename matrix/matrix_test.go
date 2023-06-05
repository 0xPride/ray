package matrix

import (
	"testing"

	"github.com/0xpride/ray/core"
)

func TestMatrixMult(t *testing.T) {
	m1 := Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	m2 := Matrix{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}}
	res := Matrix{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}

	got := m1.Mult(&m2)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixGetColum(t *testing.T) {
	m1 := Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	res := core.Tuple{X: 2, Y: 6, Z: 8, W: 4}
	got := m1.GetColum(1)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixGetRow(t *testing.T) {
	m1 := Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	res := core.Tuple{X: 5, Y: 4, Z: 3, W: 2}
	got := m1.GetRow(3)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixMultTuple(t *testing.T) {
	m1 := Matrix{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}}
	var tup, res core.Tuple
	tup.X = 1
	tup.Y = 2
	tup.Z = 3
	tup.W = 1
	res.X = 18
	res.Y = 24
	res.Z = 33
	res.W = 1
	got := m1.MultTuple(&tup)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixIdentity(t *testing.T) {
	res := Matrix{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	got := IdentityMatrix()
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m1 := Matrix{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}}
	res := Matrix{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}
	m1.Transpose()
	got := &m1
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrixGetTranspose(t *testing.T) {
	m1 := Matrix{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}}
	res := Matrix{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}
	got := m1.GetTranspose()
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestSubMatrix(t *testing.T) {
	m1 := Matrix{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}}
	res := Matrix3{{-6, 1, 6}, {-8, 8, 6}, {-7, -1, 1}}
	got := m1.SubMatrix(2, 1)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestSubMatrix3(t *testing.T) {
	m1 := Matrix3{{1, 5, 0}, {-3, 2, 7}, {0, 6, -3}}
	res := Matrix2{{-3, 2}, {0, 6}}
	got := m1.SubMatrix(0, 2)
	if *got != res {
		t.Errorf("got %v\n insted %v", *got, res)
	}
}

func TestMatrix3Minor(t *testing.T) {
	m1 := Matrix3{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}
	res := float64(25)
	got := m1.Minor(1, 0)
	if got != res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestMatrix3Cofactor(t *testing.T) {
	m1 := Matrix3{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}
	res1 := float64(-12)
	res2 := float64(-25)

	got1 := m1.Cofactor(0, 0)
	got2 := m1.Cofactor(1, 0)

	if got1 != res1 {
		t.Errorf("got %v\n insted %v", got1, res1)
	}

	if got2 != res2 {
		t.Errorf("got %v\n insted %v", got2, res2)
	}
}

func TestMatrix3Determinant(t *testing.T) {
	m1 := Matrix3{{1, 2, 6}, {-5, 8, -4}, {2, 6, 4}}
	res := float64(-196)
	got := m1.Determinant()
	if got != res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestMatrixDeterminan(t *testing.T) {
	m1 := Matrix{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}}
	res := float64(-4071)
	got := m1.Determinant()
	if got != res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestIsInvertibale(t *testing.T) {
	m1 := Matrix{{-4, 2, -2, -3}, {9, 6, 2, 6}, {0, -5, 1, -5}, {0, 0, 0, 0}}
	res := false
	got := IsInvertable(&m1)
	if got != res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestMatrixInverse(t *testing.T) {
	m1 := Matrix{{-5, 2, 6, -8}, {1, -5, 1, 8}, {7, 7, -6, -7}, {1, -3, 7, 4}}
	// res := Matrix{
	// 	{0.21805, 0.45113, 0.24060, -0.04511},
	// 	{-0.80827, -1.45677, -0.44361, 0.52068},
	// 	{-0.07895, -0.22368, -0.05263, 0.19737},
	// 	{-0.52256, -0.81391, -0.30075, 0.30639},
	// }
	determ := m1.Determinant()
	if determ != 532 {
		t.Errorf("got %v\n insted %v", determ, 532)
	}
	cof := m1.Cofactor(2, 3)
	if cof != -160 {
		t.Errorf("got %v\n insted %v", cof, -160)
	}
	// got := m1.Inverse()
	// if *got != res {
	// 	t.Errorf("got %v\n insted %v", got, res)
	// }
}

func TestTransformMatrix1(t *testing.T) {
	p := core.PointInit(-3, 4, 5)
	res := core.PointInit(2, 1, 7)
	trans := GetTransformMatrix(5, -3, 2)
	got := trans.MultTuple(p)
	if *got != *res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestTransformMatrix2(t *testing.T) {
	p := core.VectorInit(-3, 4, 5)
	res := core.VectorInit(-3, 4, 5)
	trans := GetTransformMatrix(5, -3, 2)
	got := trans.MultTuple(p)
	if *got != *res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}

func TestScaleMatrix2(t *testing.T) {
	p := core.VectorInit(-4, 6, 8)
	res := core.VectorInit(-8, 18, 32)
	trans := GetScaleMatrix(2, 3, 4)
	got := trans.MultTuple(p)
	if *got != *res {
		t.Errorf("got %v\n insted %v", got, res)
	}
}
