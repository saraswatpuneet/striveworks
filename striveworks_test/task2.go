package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"container/heap"
	"fmt"
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// lets define a Point type
type Point struct {
	X        int
	Y        int
	T        string
	Distance int
}

func (p Point) distanceFromCenter() int {
	return p.X*p.X + p.Y*p.Y
}

// PointSlice is collection of points type that implemeents simple sorting
// this extends go sort interface
type PointsCollection []Point

func (p PointsCollection) Len() int {
	return len(p)
}

func (p PointsCollection) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p PointsCollection) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// this can also be solved using priority type queue as we only care about points closer to center
type PointsQueue []Point

func (pq PointsQueue) Len() int {
	return len(pq)
}

func (pq PointsQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PointsQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PointsQueue) Push(p interface{}) {
	*pq = append(*pq, p.(Point))
}

func (pq *PointsQueue) Pop() interface{} {
	pqOld := *pq
	length := len(pqOld)
	item := pqOld[length-1]
	*pq = pqOld[0 : length-1]
	return item
}

func (pq *PointsQueue) Poll() interface{} {
	pqOld := *pq
	item := pqOld[0]
	*pq = pqOld[1:]
	return item
}

func Solution(S string, X []int, Y []int) int {
	if len(S) == 0 || len(X) == 0 || len(Y) == 0 {
		return 0
	}
	if len(S) == 1 {
		return 1
	}
	cPoints := make([]Point, 0)

	for i := 0; i < len(X); i++ {
		currentPoint := Point{X[i], Y[i], string(S[i]), 0}
		currentPoint.Distance = currentPoint.distanceFromCenter()
		cPoints = append(cPoints, currentPoint)
	}
	// simply calculate distance from center
	// if a points have same distance and tag do not include them
	return maxNonDuplicate(cPoints)

}

func maxNonDuplicate(points []Point) int {

	// sort points with distance form center

	sort.Sort(PointsCollection(points))
	pMap := make(map[string]Point)
	for _, p := range points {
		if _, ok := pMap[p.T]; ok {
			oldPoint := pMap[p.T]
			if oldPoint.distanceFromCenter() == p.distanceFromCenter() {
				// we found point with same distance and tag map it out
				return len(pMap) - 1
			} else {
				return len(pMap)
			}
		} else {
			pMap[p.T] = p
		}
	}
	return len(pMap)
}

// use priority queue
func Solution2(S string, X []int, Y []int) int {
	if len(S) == 0 || len(X) == 0 || len(Y) == 0 {
		return 0
	}
	if len(S) == 1 {
		return 1
	}
	// use priority queue
	pq := &PointsQueue{}

	for i := 0; i < len(X); i++ {
		currentPoint := Point{X[i], Y[i], string(S[i]), 0}
		currentPoint.Distance = currentPoint.distanceFromCenter()
		// push to heap
		heap.Push(pq, currentPoint)
		
	}
	// simply calculate distance from center
	// if a points have same distance and tag do not include them
	return maxNonDuplicate2(pq)

}

func maxNonDuplicate2(pq *PointsQueue) int {
	pMap := make(map[string]Point)
	for pq.Len() > 0 {
		p := pq.Poll().(Point)
		if _, ok := pMap[p.T]; ok {
			oldPoint := pMap[p.T]
			if oldPoint.distanceFromCenter() == p.distanceFromCenter() {
				// we found point with same distance and tag map it out
				return len(pMap) - 1
			} else {
				return len(pMap)
			}
		} else {
			pMap[p.T] = p
		}
	}
	return len(pMap)
}

func main() {
	fmt.Println(Solution2("ABB", []int{1, -2, -2}, []int{1, -2, 2}))

}
