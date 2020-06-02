package main

import (
	"fmt"
	"github.com/dickykmrlh/themaze/src"
)

func main(){
	//simpleMazeBuilder := &src.SimpleMaze{}
	enchantedMazeBuilder := &src.SimpleMaze{}
	mazeA := src.CreateMaze(enchantedMazeBuilder)
	mazeB := mazeA.Clone()
	StartGame(mazeB)
}

func StartGame(maze src.MazeProtoType) {
	fmt.Println("Game Start")
}
