package main

import (
	"github.com/dickykmrlh/themaze/src"
)

func main(){
	//simpleMazeFactory := &src.SimpleMaze{}
	enchantedMazeFactory := &src.SimpleMaze{}
	src.CreateMaze(enchantedMazeFactory)
}
