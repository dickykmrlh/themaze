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

func CreateMaze(builder MazeBuilder) *Maze{
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}