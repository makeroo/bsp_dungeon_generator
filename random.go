package bsp_dungeon_generator

import "math/rand"

type RandomNumberGenerator interface {
	Intn(int) int
}

type StandardRandomNumberGenerator struct {
	rnd *rand.Rand
}

func NewStandardRandomNumberGenerator(rnd *rand.Rand) StandardRandomNumberGenerator {
	return StandardRandomNumberGenerator{rnd}
}

func (s StandardRandomNumberGenerator) Intn(rng int) int {
	v := s.rnd.Intn(rng)

	return v
}
