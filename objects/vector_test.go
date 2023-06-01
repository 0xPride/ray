package objects

import (
	"math"
	"testing"
)

func TestAddVector(t *testing.T) {
	v := Vector3d{X: 1, Y: 2, Z: 3}
	v2 := Vector3d{X: 1, Y: 2, Z: 3}
	got := v.Add(&v2)

	if got.X != 2 || got.Y != 4 || got.Z != 6 {
		t.Errorf("Got %v %v %v insted of 2 4 6", got.X, got.Y, got.Z)
	}
}

func TestSubVector(t *testing.T) {
	v := Vector3d{X: 1, Y: 2, Z: 3}
	v2 := Vector3d{X: 1, Y: 2, Z: 3}
	got := v.Sub(&v2)

	if got.X != 0 || got.Y != 0 || got.Z != 0 {
		t.Errorf("Got %v %v %v insted of 0 0 0", got.X, got.Y, got.Z)
	}
}

func TestScaleVector(t *testing.T) {
	v := Vector3d{X: 1, Y: 2, Z: 3}
	got := v.Scale(2)

	if got.X != 2 || got.Y != 4 || got.Z != 6 {
		t.Errorf("Got %v %v %v insted of 2 4 6", got.X, got.Y, got.Z)
	}
}

func TestDivVector(t *testing.T) {
	v := Vector3d{X: 2, Y: 4, Z: 6}
	got := v.Div(2)

	if got.X != 1 || got.Y != 2 || got.Z != 3 {
		t.Errorf("Got %v %v %v insted of 1 2 3", got.X, got.Y, got.Z)
	}
}

func TestLenVector(t *testing.T) {
	v := Vector3d{X: 4, Y: 5, Z: 6}

	got := v.Len()
	if got != math.Sqrt(77) {
		t.Errorf("Got %v insted of %v", got, math.Sqrt(77))
	}

}

func TestGetNegateVector(t *testing.T) {
	v := Vector3d{X: 4, Y: 5, Z: 6}

	got := v.GetNegate()
	if got.X != -4 || got.Y != -5 || got.Z != -6 {
		t.Errorf("Got %v %v %v insted of -4 -5 -6", got.X, got.Y, got.Z)
	}

}

func TestNegateVector(t *testing.T) {
	got := Vector3d{X: 4, Y: 5, Z: 6}

	got.Negate()
	if got.X != -4 || got.Y != -5 || got.Z != -6 {
		t.Errorf("Got %v %v %v insted of -4 -5 -6", got.X, got.Y, got.Z)
	}

}

func TestDotVector(t *testing.T) {
	v1 := Vector3d{X: 4, Y: 5, Z: 6}
	v2 := Vector3d{X: 1, Y: 33, Z: 7}

	got := v1.Dot(&v2)
	if got != 211 {
		t.Errorf("got %v insted of 211", got)
	}
}

func TestCrossVector(t *testing.T) {
	v1 := Vector3d{X: 4, Y: 5, Z: 6}
	v2 := Vector3d{X: 1, Y: 33, Z: 7}

	got := v1.Cross(&v2)
	if got.X != -163 || got.Y != -22 || got.Z != 127 {
		t.Errorf("got %v %v %v insted of ", got.X, got.Y, got.Z)
	}
}

func TestNormVector(t *testing.T) {
	v := Vector3d{X: 1, Y: 33, Z: 7}

	got := v.GetNorm()
	if got.X != -163 || got.Y != -22 || got.Z != 127 {
		t.Errorf("got %v %v %v insted of ", got.X, got.Y, got.Z)
	}
}
