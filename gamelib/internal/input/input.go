package input

// Dir represents a direction.
type InputDir int

const (
	DirUp InputDir = iota
	DirRight
	DirDown
	DirLeft
)

// Vector returns a [-1, 1] value for each axis.
func (d InputDir) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}
