package src

import (
	"github.com/dickykmrlh/themaze/common"
)

func CreateMaze(factory MazeFactory) *Maze{
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	theDoor := factory.MakeDoor(r1, r2)

	r1.SetSide(common.North, factory.MakeWall())
	r1.SetSide(common.East, theDoor)
	r1.SetSide(common.South, factory.MakeWall())
	r1.SetSide(common.West, factory.MakeWall())

	r2.SetSide(common.North, factory.MakeWall())
	r2.SetSide(common.East, factory.MakeWall())
	r2.SetSide(common.South, factory.MakeWall())
	r2.SetSide(common.West, factory.MakeWall())

	aMaze := factory.MakeMaze()
	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	return aMaze
}