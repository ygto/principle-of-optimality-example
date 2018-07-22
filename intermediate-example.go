package main

import (
	"github.com/ygto/goui"
	"github.com/ygto/goui/entity"
	"github.com/ygto/goui/color"
	"github.com/ygto/goui-map-creator"
	"github.com/ygto/goui/entity/actions"
	"fmt"
	"math"
	"time"
)

//movement cost
var lambda = 0.9
//drawed tiles size
var tileSize = 20

func CalculateNodesValue(goalNode *mapCreator.Node) {
	queue := make([]*mapCreator.Node, 0, 0)

	queue = append(queue, goalNode)
	for {
		//if queue is empty calculation completed
		if len(queue) == 0 {
			break
		}

		var tempNode *mapCreator.Node
		tempNode, queue = queue[0], queue[1:]
		Calculate(tempNode)
		for _, n := range tempNode.GetNeighbours() {
			if !Calculated(n) {
				queue = append(queue, n)
			}
		}
	}
}

func Calculate(node *mapCreator.Node) float64 {
	if Calculated(node) {
		return node.GetValue()
	}
	var max float64
	//calculate the worst value from neighbours
	for _, n := range node.GetNeighbours() {
		max = math.Max(max, node.GetValue()+(n.GetValue()*lambda))
	}
	node.SetValue(max)

	return max
}
func Calculated(node *mapCreator.Node) bool {
	return node.GetValue() != 0
}

func SetAsFinish(node *mapCreator.Node) {
	//max value is 1
	node.SetValue(1)
}
func Travel(node *mapCreator.Node) []*mapCreator.Node {

	var travelPath []*mapCreator.Node
	//goal node
	if node.GetValue() == 1 {
		lastNode := make([]*mapCreator.Node, 0, 0)
		return append(lastNode, node)
	} else {
		travelPath = append(travelPath, node)
	}
	var max float64
	var bestNeighbour *mapCreator.Node

	//choose best neighbour
	for _, n := range node.GetNeighbours() {
		nodeVal := n.GetValue()
		if nodeVal > max {
			max = nodeVal
			bestNeighbour = n
		}
	}
	//continue to travel with best neighbour
	travelPath = append(Travel(bestNeighbour), travelPath...)

	return travelPath
}

func main() {

	//create map schema
	schema := mapCreator.NewSchema(11, 11, mapCreator.TILE_ROAD)
	schema.SetSchema(10, 10, mapCreator.TILE_GOAL)
	schema.SetSchema(0, 1, mapCreator.TILE_WALL)
	schema.SetSchema(0, 3, mapCreator.TILE_WALL)
	schema.SetSchema(0, 5, mapCreator.TILE_WALL)
	schema.SetSchema(0, 8, mapCreator.TILE_WALL)
	schema.SetSchema(0, 10, mapCreator.TILE_WALL)
	schema.SetSchema(1, 1, mapCreator.TILE_WALL)
	schema.SetSchema(1, 6, mapCreator.TILE_WALL)
	schema.SetSchema(1, 8, mapCreator.TILE_WALL)
	schema.SetSchema(2, 1, mapCreator.TILE_WALL)
	schema.SetSchema(2, 2, mapCreator.TILE_WALL)
	schema.SetSchema(2, 4, mapCreator.TILE_WALL)
	schema.SetSchema(2, 10, mapCreator.TILE_WALL)
	schema.SetSchema(3, 4, mapCreator.TILE_WALL)
	schema.SetSchema(3, 5, mapCreator.TILE_WALL)
	schema.SetSchema(3, 6, mapCreator.TILE_WALL)
	schema.SetSchema(3, 8, mapCreator.TILE_WALL)
	schema.SetSchema(4, 0, mapCreator.TILE_WALL)
	schema.SetSchema(4, 2, mapCreator.TILE_WALL)
	schema.SetSchema(4, 3, mapCreator.TILE_WALL)
	schema.SetSchema(4, 6, mapCreator.TILE_WALL)
	schema.SetSchema(4, 7, mapCreator.TILE_WALL)
	schema.SetSchema(4, 10, mapCreator.TILE_WALL)
	schema.SetSchema(5, 3, mapCreator.TILE_WALL)
	schema.SetSchema(5, 5, mapCreator.TILE_WALL)
	schema.SetSchema(5, 9, mapCreator.TILE_WALL)
	schema.SetSchema(6, 1, mapCreator.TILE_WALL)
	schema.SetSchema(6, 2, mapCreator.TILE_WALL)
	schema.SetSchema(6, 7, mapCreator.TILE_WALL)
	schema.SetSchema(6, 8, mapCreator.TILE_WALL)
	schema.SetSchema(7, 1, mapCreator.TILE_WALL)
	schema.SetSchema(7, 3, mapCreator.TILE_WALL)
	schema.SetSchema(7, 4, mapCreator.TILE_WALL)
	schema.SetSchema(7, 5, mapCreator.TILE_WALL)
	schema.SetSchema(8, 1, mapCreator.TILE_WALL)
	schema.SetSchema(8, 5, mapCreator.TILE_WALL)
	schema.SetSchema(8, 6, mapCreator.TILE_WALL)
	schema.SetSchema(8, 7, mapCreator.TILE_WALL)
	schema.SetSchema(8, 9, mapCreator.TILE_WALL)
	schema.SetSchema(9, 1, mapCreator.TILE_WALL)
	schema.SetSchema(9, 2, mapCreator.TILE_WALL)
	schema.SetSchema(9, 3, mapCreator.TILE_WALL)
	schema.SetSchema(9, 7, mapCreator.TILE_WALL)
	schema.SetSchema(9, 8, mapCreator.TILE_WALL)
	schema.SetSchema(9, 9, mapCreator.TILE_WALL)
	schema.SetSchema(9, 10, mapCreator.TILE_WALL)
	schema.SetSchema(10, 5, mapCreator.TILE_WALL)


	//change maze
	schema.SetSchema(10, 3, mapCreator.TILE_WALL)
	schema.SetSchema(9, 10, mapCreator.TILE_ROAD)



	//init goui
	g, err := goui.CreateGoui(tileSize*11, tileSize*11, "principle of optimality example")
	g.SetFps(60)
	mainScene, err := goui.CreateScene("main", g)

	if err != nil {
		panic(err)
	}
	g.AddScene(mainScene)
	g.SetActiveScene("main")

	maze := mapCreator.NewMap2D(schema, mainScene, tileSize)

	//create player
	player := entity.CreateRectangle(0, 0, float32(tileSize), float32(tileSize))
	player.SetColor(color.CreateColor(150, 150, 50))
	mainScene.AddEntity(player)

	//set goal
	goalNode := maze.GetNode(10, 10)
	SetAsFinish(goalNode)
	startNode := maze.GetNode(0, 0)
	//before to start travel nodes value must be calculated
	CalculateNodesValue(goalNode)
	//start travel and get the best travel path
	var travelPath []*mapCreator.Node
	if startNode.GetValue() != 0 {
		travelPath = Travel(startNode)
	}
	//nodes name and value
	for i, node := range travelPath {
		fmt.Println(i, node.GetName(), node.GetValue())
	}

	//movements
	var node *mapCreator.Node

	//set first action
	node, travelPath = mapCreator.ArrayPop(travelPath)
	mainAction := actions.MoveTo(startNode.GetEntity().GetPosition(), 0.05, nil)
	if node != nil {
		mainAction = actions.MoveTo(node.GetEntity().GetPosition(), 0.05, nil)
	}
	player.AddAction("move", mainAction)

	//draw
	for g.Update() {
		//if there is a path to travel and action finished, add next action
		if len(travelPath) > 0 && !mainAction.HasTarget() {
			node, travelPath = mapCreator.ArrayPop(travelPath)
			mainAction = actions.MoveTo(node.GetEntity().GetPosition(), 0.05, nil)
			player.AddAction("move", mainAction)
		}
		g.Draw()
	}
}
