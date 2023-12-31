package main

import (
	"adventofcode/utils/models"
	"adventofcode/utils/set"
	"container/heap"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main(){
	day()
	night()
}

func night() {
	  grid := models.ParseAsMatrixFromInput(input)
		grid= buildNewMatrix(grid)
	from := models.Position{X:0, Y:0}
	to := models.Position{X: grid.MaxX(), Y: grid.MaxY()}
	fmt.Println("day",dijkstra(from, to, grid))
}

func day() {
  grid := models.ParseAsMatrixFromInput(input)
	from := models.Position{X:0, Y:0}
	to := models.Position{X: grid.MaxX(), Y: grid.MaxY()}
	fmt.Println("day",dijkstra(from, to, grid))
}



func buildNewMatrix(m models.Matrix) models.Matrix {
	mm := make([][]rune, 5*len(m))
	for j, l := range m {
		for kj := 0; kj < 5; kj++ {
			new_j := kj*len(m) + j
			mm[new_j] = make([]rune, 5*len(l))
			for i, risk := range l {
				for ki := 0; ki < 5; ki++ {
					new_i := ki*len(l) + i
					new_risk := risk + rune(ki) + rune(kj)
					if new_risk > 9 {
						new_risk = new_risk % 9
					}
					mm[new_j][new_i] = rune(new_risk)
				}
			}
		}
	}
	return mm
}


func dijkstra(from, to models.Position, grid models.Matrix) int{
	pQueue := &PriorityQueue{}
	heap.Init(pQueue)
	heap.Push(pQueue, &Item{Value: from, priority: 0})
	
	visited := set.NewSet[models.Position]()
	length := map[models.Position]int{from: 0}
	for pQueue.Len() > 0 {
		currentItem := heap.Pop(pQueue).(*Item)
		current, distance := currentItem.Value, currentItem.priority
		if visited.Contains(current){
			continue
		}
		if current == to {
			return distance
		}

		visited.Add(currentItem.Value)
		for _,n:=range neighbors(currentItem.Value, grid){
			if visited.Contains(n){
				continue
			}
			potential_distance := distance + grid.IntValue(n)
			if _, ok :=length[n]; !ok || potential_distance < length[n] {
				heap.Push(pQueue,&Item{priority: potential_distance, Value: n})
				length[n]=potential_distance
			}
		}
	}
	return -1
}



func neighbors(c models.Position, grid models.Matrix) (res []models.Position){
	x,y:=c.X,c.Y
	candiPos :=[]models.Position{{X: x-1,Y: y}, {X: x+1, Y: y}, {X: x, Y: y-1}, {X: x, Y: y+1}}
	for _,ca:=range candiPos{
		if grid.IsValidPosition(ca){
			res = append(res, ca)
		}
	}
	return 
}



// using PriorityQueue implement dijkstra
type Item struct {
	Value    models.Position // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue. In this case it is weight
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value models.Position, priority int) {
	item.Value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
