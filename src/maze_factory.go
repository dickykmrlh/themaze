package src

type MazeFactory interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(n int) *Room
	MakeDoor(r1, r2 *Room) *Door
}

type SimpleMaze struct {
}

func (s *SimpleMaze) MakeMaze() *Maze {
	return &Maze{}
}

func (s *SimpleMaze) MakeWall() *Wall {
	return &Wall{}
}

func (s *SimpleMaze) MakeRoom(n int) *Room {
	return &Room{RoomNumber: 1}
}

func (s *SimpleMaze) MakeDoor(r1, r2 *Room) *Door {
	return &Door{Room1: r1, Room2: r2, isOpen: false}
}

type EnchantedMaze struct {
}

func (e *EnchantedMaze) MakeMaze() *Maze {
	return &Maze{}
}

func (e *EnchantedMaze) MakeWall() *Wall {
	return &Wall{}
}

func (e *EnchantedMaze) MakeRoom(n int) *Room {
	return &Room{RoomNumber: 1, isEnchanted: true}
}

func (e *EnchantedMaze) MakeDoor(r1, r2 *Room) *Door {
	return &Door{Room1: r1, Room2: r2, isOpen: false, isEnchanted: true}
}