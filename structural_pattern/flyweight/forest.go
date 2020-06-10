package flyweight

import (
	"fmt"
	"hash/fnv"
)

type treeType struct {
	name, color string
}

type TreeTypeFactory struct {
	treeTypePools map[uint32]*treeType
}


func getHash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (f *TreeTypeFactory)GetTreeType(name, color string) *treeType {
	if f.treeTypePools == nil {
		f.treeTypePools = make(map[uint32]*treeType)
	}

	key := getHash(name + color)
	val, ok := f.treeTypePools[key]
	if ok && val != nil{
		return val
	}

	newTreeType := &treeType{
		name: name, color: color,
	}

	f.treeTypePools[key] = newTreeType
	return newTreeType
}


type Tree struct {
	x, y int
	aType *treeType
}

func (t Tree) Draw() {
	if t.aType == nil {
		panic("tree with no type not allowed")
	}
	fmt.Printf("(%d, %d) %s : %s\n", t.x, t.y, t.aType.name, t.aType.color)
}

var forest []Tree

func PopulateForesWithSomeTree() {
	factory := TreeTypeFactory{}
	// red maple
	forest = append(forest, Tree{
		x: 1, y: 2,
		aType: factory.GetTreeType("maple", "red"),
	})

	// oak tree
	forest = append(forest, Tree{
		x: 3, y: 4,
		aType: factory.GetTreeType("oak", "black"),
	})

	// cedar tree
	forest = append(forest, Tree{
		x: 5, y: 6,
		aType: factory.GetTreeType("cedar", "brown"),
	})

	// red maple tree again
	forest = append(forest, Tree{
		x: 7, y: 8,
		aType: factory.GetTreeType("maple", "red"),
	})

	// cedar tree again
	forest = append(forest, Tree{
		x: 9, y: 10,
		aType: factory.GetTreeType("cedar", "brown"),
	})
}

func DrawForest() {
	for _, f := range forest {
		f.Draw()
	}
}
