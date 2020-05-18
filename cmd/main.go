package main

import (
	"github.com/dickykmrlh/themaze/src"
)

func main(){
	//simpleMazeFactory := &src.SimpleMaze{}
	enchantedMazeFactory := &src.EnchantedMaze{}
	src.CreateMaze(enchantedMazeFactory)
}
