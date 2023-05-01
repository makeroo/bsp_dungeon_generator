package bsp_dungeon_generator

import "math/rand"

type RandomNumberGenerator interface {
	Direction() direction

	Split(size, minSize int) int

	Rect(container Rect, minSize int) Rect
}

type StandardRandomNumberGenerator struct {
	rnd *rand.Rand
}

func NewStandardRandomNumberGenerator(rnd *rand.Rand) StandardRandomNumberGenerator {
	return StandardRandomNumberGenerator{rnd}
}

func (s StandardRandomNumberGenerator) Direction() direction {
	return direction(s.rnd.Intn(2))
}

func (s StandardRandomNumberGenerator) Split(size, minSize int) int {
	return s.rnd.Intn(size-2*minSize+1) + minSize
}

func (s StandardRandomNumberGenerator) Rect(container Rect, minSize int) Rect {
	if container.Width < minSize || container.Height < minSize {
		return Rect{0, 0, 0, 0}
	}

	room_width := s.rnd.Intn(container.Width-minSize+1) + minSize

	room_height := s.rnd.Intn(container.Height-minSize+1) + minSize

	room_x := s.rnd.Intn(container.Width - room_width + 1)

	room_y := s.rnd.Intn(container.Height - room_height + 1)

	return Rect{room_x + container.X, room_y + container.Y, room_width, room_height}
}
