package objects

type Plane struct {
	p       Point3d
	normal  Vector3d
	c       RgbColor
	epsilon float64
}

func (p Plane) Hit(r Ray, tmin float64, sr *ShadeRec) bool {
	return true
}
