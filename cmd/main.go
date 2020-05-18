package main

import (
	"github.com/dickykmrlh/themaze/src"
)

func main(){
	simpleMazeFactory := &src.SimpleMaze{}
	src.CreateMaze(simpleMazeFactory)
}
