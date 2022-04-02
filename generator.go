package bsp_dungeon_generator

type TileType int

const (
	Rock TileType = iota
	Room
	Path
)

type Rect struct {
	X, Y, Width, Height int
}

type direction int

const (
	horizontal direction = iota
	vertical
)

type BspDungeonGenerator struct {
	rnd         RandomNumberGenerator
	parent      *BspDungeonGenerator
	Rect        Rect
	Sub1, Sub2  *BspDungeonGenerator
	Room        Rect
	minStepSize int // = 4
	minRoomSize int // = 3
}

func New(width int, height int, rnd RandomNumberGenerator, minStepSize, minRoomSize int) *BspDungeonGenerator {
	return &BspDungeonGenerator{rnd, nil, Rect{0, 0, width, height}, nil, nil, Rect{}, minStepSize, minRoomSize}
}

func (g *BspDungeonGenerator) Generate() {
	g.splitSpace()

	g.generateRooms()
}

/*func (s *step) generatePath() {
	c1x, c1y := s.sub1.rect.center()
	c2x, c2y := s.sub2.rect.center()

	s.drawPath(c1x, c1y, c2x, c2y)
}*/

func (s *BspDungeonGenerator) generateRooms() {
	if !s.isLeaf() {
		s.Sub1.generateRooms()
		s.Sub2.generateRooms()

		// TODO: restore s.generatePath()

		return
	}

	room_width := s.Rect.Width - s.minRoomSize

	if room_width > 0 {
		room_width = s.rnd.Intn(room_width+1) + s.minRoomSize
	}

	room_height := s.Rect.Height - s.minRoomSize

	if room_height > 0 {
		room_height = s.rnd.Intn(room_height+1) + s.minRoomSize
	}

	room_x := s.Rect.Width - room_width

	if room_x > 0 {
		room_x = s.rnd.Intn(room_x + 1)
	}

	room_y := s.Rect.Height - room_height

	if room_y > 0 {
		room_y = s.rnd.Intn(room_y + 1)
	}

	s.Room = Rect{room_x + s.Rect.X, room_y + s.Rect.Y, room_width, room_height}
}

func (s *BspDungeonGenerator) splitSpace() {
	var size int
	var dir direction

	if s.Rect.Width <= 2*s.minStepSize {
		if s.Rect.Height <= 2*s.minStepSize {
			return
		}

		dir = horizontal
		size = s.Rect.Height

	} else if s.Rect.Height <= 2*s.minStepSize {
		dir = vertical
		size = s.Rect.Width

	} else {
		dir = direction(s.rnd.Intn(2))

		if dir == horizontal {
			size = s.Rect.Height

		} else {
			size = s.Rect.Width
		}
	}

	sub_size := s.rnd.Intn(size-2*s.minStepSize+1) + s.minStepSize

	s.split(dir, sub_size)

	s.Sub1.splitSpace()
	s.Sub2.splitSpace()
}

func (s *BspDungeonGenerator) split(dir direction, sub_size int) {
	if dir == horizontal {
		r1 := Rect{s.Rect.X, s.Rect.Y, s.Rect.Width, sub_size}
		r2 := Rect{s.Rect.X, s.Rect.Y + sub_size, s.Rect.Width, s.Rect.Height - sub_size}

		s.Sub1 = &BspDungeonGenerator{s.rnd, s, r1, nil, nil, Rect{}, s.minStepSize, s.minRoomSize}
		s.Sub2 = &BspDungeonGenerator{s.rnd, s, r2, nil, nil, Rect{}, s.minStepSize, s.minRoomSize}
	} else {
		r1 := Rect{s.Rect.X, s.Rect.Y, sub_size, s.Rect.Height}
		r2 := Rect{s.Rect.X + sub_size, s.Rect.Y, s.Rect.Width - sub_size, s.Rect.Height}

		s.Sub1 = &BspDungeonGenerator{s.rnd, s, r1, nil, nil, Rect{}, s.minStepSize, s.minRoomSize}
		s.Sub2 = &BspDungeonGenerator{s.rnd, s, r2, nil, nil, Rect{}, s.minStepSize, s.minRoomSize}
	}
}

func (s *BspDungeonGenerator) isLeaf() bool {
	return s.Sub1 == nil
}

/*func (r rect) center() (int, int) {
	return (r.Width + r.x) / 2, (r.height + r.y) / 2
}*/
