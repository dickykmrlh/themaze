package main

import (
	"github.com/dickykmrlh/themaze/creational_pattern"
)

func main(){
	//simpleMazeBuilder := &creational_pattern.SimpleMaze{}
	enchantedMazeBuilder := &creational_pattern.SimpleMaze{}
	creational_pattern.CreateMaze(enchantedMazeBuilder)
}
