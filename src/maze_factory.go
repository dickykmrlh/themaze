package src

// Maze
type Maze struct {
	rooms map[int]*Room
}

func (m *Maze) AddRoom(room *Room) {
	m.rooms[room.RoomNumber] = room
}

func (m *Maze) SearchForRoom(roomNo int) (*Room, bool) {
	r, ok := m.rooms[roomNo]
	return r, ok
}

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
	return &Door{Room1: r1, Room2: r2}
}