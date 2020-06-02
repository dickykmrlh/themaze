package src

// Maze
type Maze struct {
	rooms map[int]*Room
}

func CreateMaze(builder MazeBuilder) *Maze{
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}