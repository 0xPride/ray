package objects

import (
	"math"
	"testing"

	"github.com/0xpride/ray/utils"
)

func TestTupleAdd(t *testing.T) {
	a1 := Tuple{X: 3, Y: -2, Z: 5, W: 1}
	a2 := Tuple{X: -2, Y: 3, Z: 1, W: 0}
	res := Tuple{X: 1, Y: 1, Z: 6, W: 1}
	got := a1.Add(&a2)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleSub1(t *testing.T) {
	a1 := Tuple{X: 3, Y: 2, Z: 1, W: 1}
	a2 := Tuple{X: 5, Y: 6, Z: 7, W: 1}
	res := Tuple{X: -2, Y: -4, Z: -6, W: 0}
	got := a1.Sub(&a2)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleSub2(t *testing.T) {
	a1 := Tuple{X: 3, Y: 2, Z: 1, W: 0}
	a2 := Tuple{X: 5, Y: 6, Z: 7, W: 0}
	res := Tuple{X: -2, Y: -4, Z: -6, W: 0}
	got := a1.Sub(&a2)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleGetNegate(t *testing.T) {
	a1 := Tuple{X: 1, Y: -2, Z: 3, W: 4}
	res := Tuple{X: -1, Y: 2, Z: -3, W: -4}
	got := a1.GetNegate()

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleNegate(t *testing.T) {
	a1 := Tuple{X: 1, Y: -2, Z: 3, W: 4}
	res := Tuple{X: -1, Y: 2, Z: -3, W: -4}
	a1.Negate()

	if res != a1 {
		t.Errorf("got %v insted of %v", a1, res)
	}
}

func TestTupleScale1(t *testing.T) {
	a1 := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	res := Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14}
	got := a1.Scale(3.5)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleScale2(t *testing.T) {
	a1 := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	res := Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	got := a1.Scale(0.5)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleDiv(t *testing.T) {
	a1 := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	res := Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	got := a1.Div(2)

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleLen1(t *testing.T) {
	a1 := Tuple{X: 1, Y: 0, Z: 0, W: 0}
	res := float64(1)
	got := a1.Len()

	if res != got {
		t.Errorf("got %v insted of %v", got, res)
	}
}

func TestTupleLen2(t *testing.T) {
	a1 := Tuple{X: 0, Y: 1, Z: 0, W: 0}
	res := float64(1)
	got := a1.Len()

	if res != got {
		t.Errorf("got %v insted of %v", got, res)
	}
}

func TestTupleLen3(t *testing.T) {
	a1 := Tuple{X: 0, Y: 0, Z: 1, W: 0}
	res := float64(1)
	got := a1.Len()

	if res != got {
		t.Errorf("got %v insted of %v", got, res)
	}
}

func TestTupleLen4(t *testing.T) {
	a1 := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	res := math.Sqrt(14)
	got := a1.Len()

	if !utils.Equalf(res, got) {
		t.Errorf("got %v insted of %v", got, res)
	}
}

func TestTupleLen5(t *testing.T) {
	a1 := Tuple{X: -1, Y: -2, Z: -3, W: 0}
	res := math.Sqrt(14)
	got := a1.Len()

	if !utils.Equalf(res, got) {
		t.Errorf("got %v insted of %v", got, res)
	}
}

func TestTupleNorm(t *testing.T) {
	a1 := Tuple{X: 4, Y: 0, Z: 0, W: 0}
	res := Tuple{X: 1, Y: 0, Z: 0, W: 0}
	a1.Norm()

	if res != a1 {
		t.Errorf("got %v insted of %v", a1, res)
	}
}

func TestTupleGetNorm1(t *testing.T) {
	a1 := Tuple{X: 4, Y: 0, Z: 0, W: 0}
	res := Tuple{X: 1, Y: 0, Z: 0, W: 0}
	got := a1.GetNorm()

	if res != *got {
		t.Errorf("got %v insted of %v", *got, res)
	}
}

func TestTupleGetNorm2(t *testing.T) {
	a1 := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	got := a1.GetNorm()

	if !utils.Equalf(got.X, 0.26726) || !utils.Equalf(got.Y, 0.53452) || !utils.Equalf(got.Z, 0.80178) {
		t.Errorf("error")
	}
}

func TestNormLen(t *testing.T) {
	a1 := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	// got := a1.GetNorm()

	if got := a1.GetNorm(); !utils.Equalf(1, got.Len()) {
		t.Errorf("got %v insted of 1", got.Len())
	}
}

func TestTupleDot(t *testing.T) {
	a1 := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	a2 := Tuple{X: 2, Y: 3, Z: 4, W: 0}

	if got := a1.Dot(&a2); got != 20 {
		t.Errorf("got %v insted of 20", got)
	}
}

func TestCross(t *testing.T) {
	a1 := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	a2 := Tuple{X: 2, Y: 3, Z: 4, W: 0}
	res1 := Tuple{X: -1, Y: 2, Z: -1, W: 0}
	res2 := Tuple{X: 1, Y: -2, Z: 1, W: 0}
	got1 := a1.Cross(&a2)
	got2 := a2.Cross(&a1)
	if *got1 != res1 || *got2 != res2 {
		t.Errorf("got %v %v insted of %v %v", got1, got2, res1, res2)
	}
}
