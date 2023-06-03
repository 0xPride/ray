package objects

import "testing"

func TestVec3d(t *testing.T) {
	got := Vec3d{}

	got.Init()
	if got.X != 0 || got.Y != 0 || got.Z != 0 || got.W != 0 {
		t.Errorf("got %v insted of 0, 0, 0, 0", got)
	}
}
