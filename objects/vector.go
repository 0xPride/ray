package objects

type Vec3d struct {
	tuple
}

func (self *Vec3d) Init() {
	self.W = 0
}
