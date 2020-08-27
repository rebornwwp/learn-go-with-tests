package structs

type Rectangle struct {
	height float64
	width  float64
}

func (self *Rectangle) Area() float64 {
	return self.height * self.width
}
