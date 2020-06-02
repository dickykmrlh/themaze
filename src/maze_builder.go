package src

import (
	"github.com/dickykmrlh/themaze/common"
	"math/rand"
)

func (m *Maze) AddRoom(room *Room) {
	m.rooms[room.RoomNumber] = room
}

func (m *Maze) SearchForRoom(roomNo int) (*Room, bool) {
	r, ok := m.rooms[roomNo]
	return r, ok
}

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(n int)
	BuildDoor(roomFrom, roomTo int)
	GetMaze() *Maze
}

type SimpleMaze struct {
	currentMaze *Maze
}


func (s *SimpleMaze) BuildMaze()  {
	s.currentMaze = &Maze{}
}

func (s *SimpleMaze) BuildRoom(n int) {
	if _, ok := s.currentMaze.rooms[n]; !ok {
		r1 := &Room{RoomNumber: 1}
		r1.SetSide(common.North, &Wall{})
		r1.SetSide(common.East, &Wall{})
		r1.SetSide(common.South, &Wall{})
		r1.SetSide(common.West, &Wall{})

		s.currentMaze.AddRoom(r1)
	}
}

func (s *SimpleMaze) BuildDoor(roomFrom, roomTo int) {
	r1 := s.currentMaze.rooms[roomFrom]
	r2  := s.currentMaze.rooms[roomTo]
	d := &Door{r1, r2, false, false}

	wallDirrection := rand.Intn(4)
	oppositeWallDirrection := (wallDirrection + 2) % 4

	r1.SetSide(common.Direction(wallDirrection), d)
	r1.SetSide(common.Direction(oppositeWallDirrection), d)
}

func (s *SimpleMaze) GetMaze() *Maze {
	return s.currentMaze
}

type EnchantedMaze struct {
	currentMaze *Maze
}


func (e *EnchantedMaze) BuildMaze()  {
	e.currentMaze = &Maze{}
}

func (e *EnchantedMaze) BuildRoom(n int) {
	if _, ok := e.currentMaze.rooms[n]; !ok {
		r1 := &Room{RoomNumber: 1, isEnchanted: true}
		r1.SetSide(common.North, &Wall{})
		r1.SetSide(common.East, &Wall{})
		r1.SetSide(common.South, &Wall{})
		r1.SetSide(common.West, &Wall{})

		e.currentMaze.AddRoom(r1)
	}
}

func (e *EnchantedMaze) BuildDoor(roomFrom, roomTo int) {
	r1 := e.currentMaze.rooms[roomFrom]
	r2  := e.currentMaze.rooms[roomTo]
	d := &Door{Room1: r1, Room2: r2, isOpen: false, isEnchanted: true}

	wallDirrection := rand.Intn(4)
	oppositeWallDirrection := (wallDirrection + 2) % 4

	r1.SetSide(common.Direction(wallDirrection), d)
	r1.SetSide(common.Direction(oppositeWallDirrection), d)
}

func (e *EnchantedMaze) GetMaze() *Maze {
	return e.currentMaze
}