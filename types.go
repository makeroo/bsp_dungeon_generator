package bsp_dungeon_generator

type Rect struct {
	X, Y, Width, Height int
}

type direction int

const (
	horizontal direction = iota
	vertical
)

type Segment struct {
	X1, Y1, X2, Y2 int
}

type PathDefinition struct {
	Segments []Segment
}

type BspDungeonGenerator struct {
	rnd         RandomNumberGenerator
	parent      *BspDungeonGenerator
	Rect        Rect
	Sub1, Sub2  *BspDungeonGenerator
	Room        Rect
	Path        *PathDefinition
	minStepSize int // = 4
	minRoomSize int // = 3
}
