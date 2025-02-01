package shape

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}
