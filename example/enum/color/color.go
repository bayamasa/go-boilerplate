package color

//go:generate stringer -type=Color
type Color int

const (
	Red Color = iota + 1
	Yellow
	Blue
)
