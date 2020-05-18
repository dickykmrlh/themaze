package src

import (
	"fmt"
	"github.com/dickykmrlh/themaze/common"
)

type MapSite interface {
	Enter()
}

// Room
type Room struct {
	sides [4]MapSite
	RoomNumber int
}

func (r *Room) SetSide(direction common.Direction, site MapSite) {
	r.sides[int(direction)] = site
}
func (r *Room) GetSide(direction common.Direction) MapSite {
	return r.sides[int(direction)]
}

func (r *Room) Enter() {
	fmt.Printf("you enter room: %d\n", r.RoomNumber)
}

func (r *Room) GetRoomNumber() int{
	return r.RoomNumber
}
// Wall
type Wall struct {
}

func (w *Wall) Enter() {
	fmt.Println("you hit a wall")
}

// Door
type Door struct {
	Room1, Room2 *Room
	isOpen bool
}

func (d *Door) Enter() {
	if d.isOpen {
		fmt.Printf("You enter room: %d\n", d.Room2.RoomNumber)
	}
	fmt.Println("the door is closed, you broke ur nose")
}

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

func CreateMaze() *Maze{
	r1 := Room{RoomNumber: 1}
	r2 := Room{RoomNumber: 2}
	theDoor := &Door{Room1: &r1, Room2: &r2}

	r1.SetSide(common.North, &Wall{})
	r1.SetSide(common.East, theDoor)
	r1.SetSide(common.South, &Wall{})
	r1.SetSide(common.West, &Wall{})

	r2.SetSide(common.North, &Wall{})
	r2.SetSide(common.East, &Wall{})
	r2.SetSide(common.South, &Wall{})
	r2.SetSide(common.West, theDoor)

	aMaze := &Maze{}
	aMaze.AddRoom(&r1)
	aMaze.AddRoom(&r2)

	return aMaze
}