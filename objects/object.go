package objects

type Hitter interface {
	RgbColor
	Hit(r Ray, tmin float64, sr *ShadeRec) bool
}

type ShadeRec struct {
	HitAnObject   bool
	LocalHitpoint Point3d
	Normal        Vector3d
	// w *World
}
