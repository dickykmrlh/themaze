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