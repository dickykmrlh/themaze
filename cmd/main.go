package main

import (
	"github.com/dickykmrlh/themaze/src"
)

func main(){
	//simpleMazeBuilder := &src.SimpleMaze{}
	enchantedMazeBuilder := &src.SimpleMaze{}
	src.CreateMaze(enchantedMazeBuilder)
}
