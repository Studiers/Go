package shape

type Rect struct { width, height float64 }

func (r *Rect) area() float64 {
	return (*r).width * (*r).height
}