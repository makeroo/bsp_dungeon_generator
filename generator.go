package bsp_dungeon_generator

import "fmt"

type TileType int

const (
	Rock TileType = iota
	Room
	Path
)

const min_step_size int = 4
const min_room_size int = 3

type BspDungeonGenerator struct {
	Width  int
	Height int
	tiles  []TileType
}

type rect struct {
	x, y, width, height int
}

type direction int

const (
	horizontal direction = iota
	vertical
)

type step struct {
	g          *BspDungeonGenerator
	rnd        RandomNumberGenerator
	parent     *step
	rect       rect
	dir        direction
	sub1, sub2 *step
	room       rect
}

func New(width int, height int) *BspDungeonGenerator {
	tiles := make([]TileType, height*width)

	return &BspDungeonGenerator{width, height, tiles}
}

func (g *BspDungeonGenerator) Generate(rnd RandomNumberGenerator) {
	s := &step{g, rnd, nil, rect{0, 0, g.Width, g.Height}, horizontal, nil, nil, rect{}}

	s.splitSpace()

	s.generateRooms()

	g.dump(s)
}

func (g *BspDungeonGenerator) dump(s *step) {
	//var d int = 0
	var f func(*step)
	//tiles := make([]int, g.Width*g.Height)

	/*drawRect := func(r rect) {
		for y := 0; y < r.height; y++ {
			for x := 0; x < r.width; x++ {
				tiles[y*g.Width+x] = d
			}
		}
	}*/

	f = func(s *step) {
		fmt.Printf("{ \"rect\": {\"x\": %d, \"y\": %d, \"width\": %d, \"height\": %d }",
			s.rect.x, s.rect.y, s.rect.width, s.rect.height,
		)

		if s.sub1 == nil && s.sub2 == nil {
			//drawRect(s.rect)
			fmt.Printf(", \"room\": {\"x\": %d, \"y\": %d, \"width\": %d, \"height\": %d }",
				s.room.x, s.room.y, s.room.width, s.room.height,
			)
		}

		//d++

		if s.sub1 != nil {
			fmt.Print(", \"sub1\": ")
			f(s.sub1)
		}
		if s.sub2 != nil {
			fmt.Print(", \"sub2\": ")
			f(s.sub2)
		}

		fmt.Println(" }")
	}

	f(s)

	/*for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			t := tiles[y*g.Width+x]

			fmt.Printf("%d", tiles[t])
		}

		fmt.Println()
	}*/
}

func (g *BspDungeonGenerator) TileAt(x, y int) TileType {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return Rock
	}

	return g.getTile(x, y)
}

func (g *BspDungeonGenerator) getTile(x, y int) TileType {
	return g.tiles[y*g.Width+x]
}

func (g *BspDungeonGenerator) setTile(x, y int, tileType TileType) {
	g.tiles[y*g.Width+x] = tileType
}

func (s *step) drawRect(r rect) {
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			s.g.setTile(x+r.x, y+r.y, Room)
		}
	}
}

func (s *step) drawPath(x1, y1, x2, y2 int) {
	if x1 == x2 {
		for y := y1; y <= y2; y++ {
			if s.g.getTile(x1, y) != Rock {
				continue
			}

			s.g.setTile(x1, y, Path)
		}

		return
	}

	if y1 == y2 {
		for x := x1; x <= x2; x++ {
			if s.g.getTile(x, y1) != Rock {
				continue
			}

			s.g.setTile(x, y1, Path)
		}

		return
	}

	panic("cannot happen") // TODO: handle error properly
}

func (s *step) generatePath() {
	c1x, c1y := s.sub1.rect.center()
	c2x, c2y := s.sub2.rect.center()

	s.drawPath(c1x, c1y, c2x, c2y)
}

func (s *step) generateRooms() {
	if !s.isLeaf() {
		s.sub1.generateRooms()
		s.sub2.generateRooms()

		s.generatePath()

		return
	}

	room_width := s.rect.width - min_room_size

	if room_width > 0 {
		room_width = s.rnd.Intn(room_width+1) + min_room_size
	}

	room_height := s.rect.height - min_room_size

	if room_height > 0 {
		room_height = s.rnd.Intn(room_height+1) + min_room_size
	}

	room_x := s.rect.width - room_width

	if room_x > 0 {
		room_x = s.rnd.Intn(room_x + 1)
	}

	room_y := s.rect.height - room_height

	if room_y > 0 {
		room_y = s.rnd.Intn(room_y + 1)
	}

	s.room = rect{room_x + s.rect.x, room_y + s.rect.y, room_width, room_height}

	s.drawRect(s.room)
}

func (s *step) splitSpace() {
	if s.rect.width <= 2*min_step_size || s.rect.height <= 2*min_step_size {
		return
	}

	var size int

	if s.dir == horizontal {
		size = s.rect.height

	} else {
		size = s.rect.width
	}

	sub_size := s.rnd.Intn(size-2*min_step_size+1) + min_step_size

	s.split(sub_size)

	s.sub1.splitSpace()
	s.sub2.splitSpace()
}

func (s *step) split(sub_size int) {
	if s.dir == horizontal {
		r1 := rect{s.rect.x, s.rect.y, s.rect.width, sub_size}
		r2 := rect{s.rect.x, s.rect.y + sub_size, s.rect.width, s.rect.height - sub_size}

		s.sub1 = &step{s.g, s.rnd, s, r1, vertical, nil, nil, rect{}}
		s.sub2 = &step{s.g, s.rnd, s, r2, vertical, nil, nil, rect{}}
	} else {
		r1 := rect{s.rect.x, s.rect.y, sub_size, s.rect.height}
		r2 := rect{s.rect.x + sub_size, s.rect.y, s.rect.width - sub_size, s.rect.height}

		s.sub1 = &step{s.g, s.rnd, s, r1, horizontal, nil, nil, rect{}}
		s.sub2 = &step{s.g, s.rnd, s, r2, horizontal, nil, nil, rect{}}
	}
}

func (s *step) isLeaf() bool {
	return s.sub1 == nil
}

func (r rect) center() (int, int) {
	return (r.width + r.x) / 2, (r.height + r.y) / 2
}
