package bsp_dungeon_generator

func (r Rect) center() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

func New(width int, height int, rnd RandomNumberGenerator, minStepSize, minRoomSize int) *BspDungeonGenerator {
	return &BspDungeonGenerator{rnd, nil, Rect{0, 0, width, height}, nil, nil, Rect{}, nil, minStepSize, minRoomSize}
}

func (g *BspDungeonGenerator) Generate() {
	g.splitSpace()

	g.generateRooms()
}

func (s *BspDungeonGenerator) generatePath() {
	c1x, c1y := s.Sub1.Rect.center()
	c2x, c2y := s.Sub2.Rect.center()

	s.Path = &PathDefinition{[]Segment{{c1x, c1y, c2x, c2y}}}
}

func (s *BspDungeonGenerator) generateRooms() {
	if !s.isLeaf() {
		s.Sub1.generateRooms()
		s.Sub2.generateRooms()

		s.generatePath()

		return
	}

	s.Room = s.rnd.Rect(s.Rect, s.minRoomSize)
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
		dir = s.rnd.Direction()

		if dir == horizontal {
			size = s.Rect.Height

		} else {
			size = s.Rect.Width
		}
	}

	sub_size := s.rnd.Split(size, s.minStepSize)

	s.split(dir, sub_size)

	s.Sub1.splitSpace()
	s.Sub2.splitSpace()
}

func (s *BspDungeonGenerator) split(dir direction, sub_size int) {
	if dir == horizontal {
		r1 := Rect{s.Rect.X, s.Rect.Y, s.Rect.Width, sub_size}
		r2 := Rect{s.Rect.X, s.Rect.Y + sub_size, s.Rect.Width, s.Rect.Height - sub_size}

		s.Sub1 = &BspDungeonGenerator{s.rnd, s, r1, nil, nil, Rect{}, nil, s.minStepSize, s.minRoomSize}
		s.Sub2 = &BspDungeonGenerator{s.rnd, s, r2, nil, nil, Rect{}, nil, s.minStepSize, s.minRoomSize}
	} else {
		r1 := Rect{s.Rect.X, s.Rect.Y, sub_size, s.Rect.Height}
		r2 := Rect{s.Rect.X + sub_size, s.Rect.Y, s.Rect.Width - sub_size, s.Rect.Height}

		s.Sub1 = &BspDungeonGenerator{s.rnd, s, r1, nil, nil, Rect{}, nil, s.minStepSize, s.minRoomSize}
		s.Sub2 = &BspDungeonGenerator{s.rnd, s, r2, nil, nil, Rect{}, nil, s.minStepSize, s.minRoomSize}
	}
}

func (s *BspDungeonGenerator) isLeaf() bool {
	return s.Sub1 == nil
}

/*func (r rect) center() (int, int) {
	return (r.Width + r.x) / 2, (r.height + r.y) / 2
}*/
