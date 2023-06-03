package objects

type Point3d struct {
	tuple
}

func (self *Point3d) Init() {
	self.W = 1
}
