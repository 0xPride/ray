package objects

import "testing"

func TestPoint3d(t *testing.T) {
	got := Point3d{}

	got.Init()
	if got.X != 0 || got.Y != 0 || got.Z != 0 || got.W != 1 {
		t.Errorf("got %v insted of 0, 0, 0, 1", got)
	}
}
