package bsp_dungeon_generator

type Dungeon struct {
	Width  int
	Height int
	tiles  []TileType
}

func (g *Dungeon) TileAt(x, y int) TileType {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return Rock
	}

	return g.getTile(x, y)
}

func (g *Dungeon) getTile(x, y int) TileType {
	return g.tiles[y*g.Width+x]
}

/*func (g *Dungeon) setTile(x, y int, tileType TileType) {
	g.tiles[y*g.Width+x] = tileType
}*/

/*
func New(width int, height int) *BspDungeonGenerator {
	tiles := make([]TileType, height*width)

	return &BspDungeonGenerator{width, height, tiles}
}
*/
