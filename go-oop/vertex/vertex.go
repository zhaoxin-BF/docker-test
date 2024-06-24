package vertex

type Vertex[T int | int64 | float64] struct {
	X T
	Y T
}

func New[T int | int64 | float64](x, y T) *Vertex[T] {
	return &Vertex[T]{
		X: x,
		Y: y,
	}
}

func (v *Vertex[T]) Abs() T {
	ret := v.X*v.X + v.Y*v.Y
	v.X = 10
	return ret
}
